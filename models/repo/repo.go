package repo

import (
	"fmt"
	"log"
	"time"

	"github.com/Cheep2Workshop/proj-web/models"
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
	tx     *gorm.DB
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
	client := &DbClient{db, nil, *config}
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
