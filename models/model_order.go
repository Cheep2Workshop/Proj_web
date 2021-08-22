package models

import "time"

type Order struct {
	Id           int           `gorm:"primary_key;AUTO_INCREMENT"`
	UserId       int           `gorm:"foreign_key; NOT_NULL"`
	CreatedAt    time.Time     `gorm:"timestamp"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderId;references:Id"`
}

type SaleItem interface {
	GetSumPrice() int
}

func (o *Order) GetSumPrice() int {
	sum := 0
	for _, detail := range o.OrderDetails {
		sum += detail.GetSumPrice()
	}
	return sum
}
