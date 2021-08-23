package models

import "time"

type DashboardLoginLog struct {
	ID        int       `gorm:"primary_key; AUTO_INCREMENT"`
	UserId    int       `gorm:"foreign_key"`
	CreatedAt time.Time `gorm:"type:timestamp; NOT_NULL"`
}
