package models

import "gorm.io/plugin/soft_delete"

type Product struct {
	Id          int                   `gorm:"primary_key; AUTO_INCREMENT"`
	ProductName string                `gorm:"type:varchar(50); NOT_NULL"`
	ProductDesc string                `gorm:"type:varchar(50)"`
	Price       int                   `gorm:"default:0;NOT_NULL"`
	DeleteAt    soft_delete.DeletedAt `gorm:"default:0"`
}
