package models

type OrderDetail struct {
	Id            int     `gorm:"primary_key; AUTO_INCREMENT"`
	OrderId       int     `gorm:"foreign_key; NOT_NULL"`
	ProductId     int     `gorm:"foreign_key"`
	ProductAmount int     `gorm:"NOT_NULL"`
	Product       Product `gorm:"foreignKey:ProductId"`
}

func (od *OrderDetail) GetSumPrice() int {
	return od.Product.Price * od.ProductAmount
}
