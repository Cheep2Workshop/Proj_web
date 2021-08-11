package models

import "time"

type User struct {
	ID        int       `gorm:"primary_key; AUTO_INCREMENT"`
	Name      string    `gorm:"type:varchar(50); NOT_NULL"`
	Email     string    `gorm:"type:varchar(50); UNIQUE; NOT_NULL"`
	Password  string    `gorm:"type:varchar(50); NOT_NULL"`
	Admin     bool      `gorm:"type:bool; default:false"`
	Banned    bool      `gorm:"type:bool; default:false"`
	CreatedAt time.Time `gorm:"type:timestamp; NOT_NULL"`
	// Role []string
}

type DashboardLoginLog struct {
	ID        int       `gorm:"primary_key; AUTO_INCREMENT"`
	UserId    int       `gorm:"foreign_key"`
	CreatedAt time.Time `gorm:"type:timestamp; NOT_NULL"`
}
