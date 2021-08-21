package repo

import (
	"errors"

	"github.com/Cheep2Workshop/proj-web/models"
)

type TradeDetail struct {
	Product models.Product
	Amount  int `json:",omitempty"`
	// TotalPrice int `json:",omitempty"`
}

type OrderReq struct {
	UserId    int
	ProductId []int
	Amount    []int
}

type OrderRes struct {
	UserId   int    `json:",omitempty"`
	UserName string `json:",omitempty"`
	// TotalPrice int    `json:",omitempty"`
	Details []TradeDetail
}

func (client *DbClient) AddOrder(req OrderReq) error {
	var err error
	if len(req.ProductId) != len(req.Amount) {
		return errors.New("Product and amount not matched.")
	}
	tx := client.Begin()
	order := models.Order{
		UserId: req.UserId,
	}
	if err = tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	details := make([]models.OrderDetail, len(req.ProductId))
	for i, p := range req.ProductId {
		details[i] = models.OrderDetail{
			OrderId:       order.Id,
			ProductId:     p,
			ProductAmount: req.Amount[i],
		}
	}
	// debug : diff of batch vs create
	if err = tx.Create(&details).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (client *DbClient) GetOrders(oids ...int) ([]models.Order, error) {
	// SELECT `users`.`id`, `users`.`name`, `products`.*, `order_details`.`product_amount`
	// FROM `orders`
	// JOIN `order_details`
	// ON `order_details`.`order_id` = `orders`.`id`
	// JOIN `users`
	// ON `users`.`id` = `orders`.`user_id`
	// JOIN `products`
	// ON `products`.`id` = `order_details`.`product_id`

	// also can user preload
	// orders := []models.Order{}
	// err := client.Debug().
	//Table("orders").
	// Select("*").
	// Where("id=?", oids).
	// Joins("LEFT JOIN order_details ON order_details.order_id = orders.id").
	// Joins("users ON users.id = orders.user_id").
	// Joins("products On products.id = order_details.product_id").
	// Find(&orders).Error
	// return orders, err
	orders := []models.Order{}
	err := client.Debug().
		Where("id=?", oids).
		Preload("OrderDetails").
		Find(&orders).Error
	return orders, err
}
