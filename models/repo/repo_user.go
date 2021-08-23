package repo

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/Cheep2Workshop/proj-web/models"
	utils "github.com/Cheep2Workshop/proj-web/utils/encryption"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

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
	result := client.Create(&user)
	if result.Error != nil {
		return false, result.Error
	}
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
	result := client.Scopes().Model(&models.User{}).First(&user, "name = ? AND email = ?", name, email)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

type LoginReq struct {
	Email    string `json:",omitempty"`
	Password string `json:",omitempty"`
}

// Begin transation of login
func (client *DbClient) BeginLogin(req LoginReq) (*models.User, error) {
	var user *models.User
	hash := utils.DefaultParams.ToHash(Salt)
	client.tx = client.Begin()
	result := client.tx.Model(&models.User{}).First(&user, "email = ?", req.Email)
	if result.Error != nil {
		return nil, result.Error
	}
	// check matched password
	match, err := utils.ComparePasswordAndHash(req.Password, user.Password, hash)
	if match {
		return user, nil
	}
	return nil, err
}

// Cancel and rollback transation of login
func (client *DbClient) CancelLogin() error {
	if client.tx == nil {
		return errors.New("tx is nil.")
	}
	client.tx.Rollback()
	client.tx = nil
	return nil
}

// End and commit of login
func (client *DbClient) EndLogin() error {
	if client.tx == nil {
		return errors.New("tx is nil.")
	}
	client.tx.Commit()
	client.tx = nil
	return nil
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
	var db *gorm.DB
	if client.tx != nil {
		db = client.tx
	} else {
		db = client.DB
	}
	result := db.First(&user, "email = ?", email)
	if result.Error != nil {
		return result.Error
	}
	loginLog := models.DashboardLoginLog{UserId: user.ID}
	db.Create(&loginLog)
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
		result := client.Where("email = ?", req.DeleteEmail).Delete(&models.User{})
		return result.Error
	}
	return errors.New("Permission denied to delete user.")
}
