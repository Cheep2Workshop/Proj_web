package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Cheep2Workshop/proj-web/models"
	utils "github.com/Cheep2Workshop/proj-web/utils/encryption"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	Account  string
	Password string
	Ip       string
	Port     string
	DbName   string
}

type DbClient struct {
	*gorm.DB
	Config DbConfig
}

const (
	PwdMin = 6
	PwdMax = 16
	Salt   = "WiZilZnYZ4nNuA5SBKLSyw"
)

var Client *DbClient

// get the db url of config
func (config *DbConfig) GetUrl() string {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Account, config.Password, config.Ip, config.Port, config.DbName)
	return url
}

// connect db with config (enable cache the client)
func (config *DbConfig) ConnectDb(cacheClient bool) (*DbClient, error) {
	url := config.GetUrl()
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	client := &DbClient{db, *config}
	if cacheClient {
		Client = client
	}
	return client, err
}

func InitRemoteDB(origin *DbConfig, dest *DbConfig) {
	client, err := origin.ConnectDb(false)
	if err != nil {
		log.Panic(err)
	}

	sqldb, err := client.DB.DB()
	if err != nil {
		log.Panic(err)
	}
	sqldb.SetConnMaxLifetime(1 * time.Second)
	sqldb.SetMaxIdleConns(20)
	sqldb.SetMaxOpenConns(2000)

	if result := client.Exec("CREATE DATABASE IF NOT EXISTS dashboard;"); result.Error != nil {
		log.Panic(result.Error)
	}
	client, err = dest.ConnectDb(true)
	if err != nil {
		log.Panic(err)
	}

	client.AutoMigrate(&models.User{}, &models.DashboardLoginLog{})
}

// init db with create needed tables
func (client *DbClient) Init() {
	sqldb, err := client.DB.DB()
	if err != nil {
		log.Panic(err)
	}
	sqldb.SetConnMaxLifetime(1 * time.Second)
	sqldb.SetMaxIdleConns(20)
	sqldb.SetMaxOpenConns(2000)

	if result := client.Exec("CREATE DATABASE IF NOT EXISTS dashboard;"); result.Error != nil {
		log.Panic(result.Error)
	}
	if err != nil {
		log.Panic(err)
	}

	client.AutoMigrate(&models.User{}, &models.DashboardLoginLog{})
}

// drop all talbe of db
func (client *DbClient) DropAllTables() {
	client.Migrator().DropTable(&models.User{}, &models.DashboardLoginLog{})
}

func (client *DbClient) Signup(user models.User) (bool, error) {
	exist := client.CheckUserExist(user.Email)
	if exist {
		return false, errors.New("The email of account has been existed.")
	}
	_, keyHash, err := utils.GenerateFromPassword(user.Password, Salt, utils.DefaultParams)
	if err != nil {
		return false, err
	}
	user.Password = keyHash
	client.Create(&user)
	log.Println(user)
	return true, nil
}

func (client *DbClient) CheckUserExist(email string) bool {
	var users []models.User
	client.Model(&models.User{}).Where("email = ?", email).Scan(&users)
	return len(users) > 0
}

func (client *DbClient) CheckAuth(name string, email string) (bool, error) {
	var user *models.User
	result := client.Model(&models.User{}).First(&user, "name = ? AND email = ?", name, email)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

type LoginReq struct {
	Email    string `json:",omitempty"`
	Password string `json:",omitempty"`
}

func (client *DbClient) Login(req LoginReq) (*models.User, error) {
	var user *models.User
	hash := utils.DefaultParams.ToHash(Salt)
	client.Model(&models.User{}).First(&user, "email = ?", req.Email)
	// check matched password
	match, err := utils.ComparePasswordAndHash(req.Password, user.Password, hash)
	if match {
		return user, nil
	}
	return user, err
}

func (client *DbClient) SaveLoginLog(email string) error {
	var user models.User
	result := client.First(&user, "email = ?", email)
	if result.Error != nil {
		return result.Error
	}
	loginLog := models.DashboardLoginLog{UserId: user.ID}
	client.Create(&loginLog)
	log.Printf("Save login log: %v (%v) at %s", user.Name, user.Email, loginLog.CreatedAt.Format("2006-01-02 15:04:05"))
	return nil
}

func (client *DbClient) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := client.Omit("ID", "Password").First(&user, "email = ?", email)
	return user, result.Error
}

type LoginLogRes struct {
	ID        int       `json:",omitempty"`
	Name      string    `json:",omitempty"`
	Email     string    `json:",omitempty"`
	Admin     bool      `json:",omitempty"`
	CreatedAt time.Time `json:",omitempty"`
}

type LoginLogList []LoginLogRes

func (logRes LoginLogRes) MarshalBinary() ([]byte, error) {
	return json.Marshal(logRes)
}

func (list LoginLogList) MarshalBinary() ([]byte, error) {
	return json.Marshal(list)
}

func (client *DbClient) GetLoginLogs(email string) ([]LoginLogRes, error) {
	var logs []LoginLogRes
	result := client.Model(&models.User{}).Select("dashboard_login_logs.id, users.name, users.email, users.admin, dashboard_login_logs.created_at").Where("users.email = ?", email).Joins("JOIN dashboard_login_logs ON users.id=dashboard_login_logs.user_id").Scan(&logs)
	return logs, result.Error
}

type SetUserReq struct {
	Email           string `json:"Email,omitempty"`
	TargetEmail     string `json:"TargetEmail,omitempty"`
	ChangedName     string `json:"ChangedName,omitempty"`
	ChangedPassword string `json:"ChangedPassword,omitempty"`
}

func (client *DbClient) SetUser(req SetUserReq) error {
	// check the user who requested has permission
	user, err := client.GetUserByEmail(req.Email)
	if err != nil {
		return errors.New("Unknown email.")
	}
	if user.Email != req.TargetEmail && !user.Admin {
		return errors.New("Permission denied.")
	}

	nameLen := len(req.ChangedName)
	newPwdLen := len(req.ChangedPassword)
	// check if name of password not empty
	if nameLen <= 0 && newPwdLen <= 0 {
		return errors.New("Either name must be not empty or new password must be between 6 and 16.")
	}
	// check if password is between min and max
	if newPwdLen > 0 && newPwdLen < PwdMin && newPwdLen > PwdMax {
		return errors.New("New password must be between 6 and 16.")
	}
	// try get user matched email
	target := models.User{Email: req.TargetEmail}
	result := client.Model(&models.User{}).First(&target, &target)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Before set user data %v", target)
	// write name if not empty
	if len(req.ChangedName) > 0 {
		target.Name = req.ChangedName
	}
	// write password if not empty
	if newPwdLen >= PwdMin && newPwdLen <= PwdMax {
		_, pwdHash, err := utils.GenerateFromPassword(req.ChangedPassword, Salt, utils.DefaultParams)
		if err != nil {
			return err
		}
		target.Password = pwdHash
	}
	result = client.Save(&target)
	log.Printf("After set user data %v", target)
	return result.Error
}

type DeleteUserReq struct {
	Email       string `json:",omitempty"`
	DeleteEmail string `json:",omitempty"`
}

func (client *DbClient) DeleteUser(req DeleteUserReq) error {
	// get the author
	author, err := client.GetUserByEmail(req.Email)
	if err != nil {
		return err
	}
	// check the access of delete
	access := req.Email == req.DeleteEmail || author.Admin
	if access {
		user := &models.User{Email: req.DeleteEmail}
		result := client.Delete(&user, &user)
		return result.Error
	}
	return errors.New("Permission denied to delete user.")
}
