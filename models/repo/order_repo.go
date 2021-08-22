package repo

import (
	"errors"

	"github.com/Cheep2Workshop/proj-web/models"
	"gorm.io/gorm/clause"
)

type OrderReq struct {
	UserId    int
	ProductId []int
	Amount    []int
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
	var err error
	orders := []models.Order{}

	// join method (not useful)
	// err = client.Debug().
	// 	Select("*").
	// 	//Where("id=?", oids).
	// 	Joins("JOIN order_details ON orders.id=order_details.order_id").
	// 	Find(&orders).
	// 	Error
	// return orders, err

	// preload method
	err = client.Debug().
		Where("id=?", oids).
		Preload(clause.Associations).
		// Preload("OrderDetails").
		Preload("OrderDetails.Product").
		Find(&orders).Error
	return orders, err
}
