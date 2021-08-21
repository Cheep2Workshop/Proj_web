package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type User struct {
	ID        int                   `gorm:"primary_key; AUTO_INCREMENT"`
	Name      string                `gorm:"type:varchar(50); NOT_NULL"`
	Email     string                `gorm:"type:varchar(50); UNIQUE; NOT_NULL"`
	Password  string                `gorm:"type:varchar(50); NOT_NULL"`
	Admin     bool                  `gorm:"type:bool; default:false"`
	CreatedAt time.Time             `gorm:"type:timestamp; NOT_NULL"`
	DeleteAt  soft_delete.DeletedAt `gorm:"default:0"`
	// Role []string
	// Orders []Order				`gorm:"foreignKey:user_id;references:id"`
}
