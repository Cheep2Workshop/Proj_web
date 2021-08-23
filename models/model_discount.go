package models

import "time"

type Discount struct {
	Id         int       `gorm:"primary_key"`
	ProductId  int       `gorm:"foreign_key; NOT_NULL"`
	Percentage float32   `gorm:"default:0;NOT_NULL"`
	StartAt    time.Time `gorm:"type:timestamp;NOT_NULL"`
	EndAt      time.Time `gorm:"type:timestamp"`
}
