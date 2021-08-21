package models

import "time"

type Order struct {
	Id           int           `gorm:"primary_key;AUTO_INCREMENT"`
	UserId       int           `gorm:"foreign_key; NOT_NULL"`
	CreatedAt    time.Time     `gorm:"timestamp"`
	OrderDetails []OrderDetail `gorm:"foreignKey:order_id;references:id"`
}
