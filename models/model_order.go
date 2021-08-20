package models

import "time"

type Order struct {
	Id        int       `gorm:"primary_key;AUTO_INCREMENT"`
	UserId    int       `gorm:"foreign_key"`
	CreatedAt time.Time `gorm:"timestamp"`
}
