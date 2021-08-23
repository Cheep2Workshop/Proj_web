package models

import "gorm.io/plugin/soft_delete"

type Product struct {
	Id          int                   `gorm:"primary_key; AUTO_INCREMENT" json:",omitempty"`
	ProductName string                `gorm:"type:varchar(50); NOT_NULL" json:",omitempty"`
	ProductDesc string                `gorm:"type:varchar(50)" json:",omitempty"`
	Price       int                   `gorm:"default:0;NOT_NULL" json:",omitempty"`
	DeleteAt    soft_delete.DeletedAt `gorm:"default:0" json:",omitempty"`
}
