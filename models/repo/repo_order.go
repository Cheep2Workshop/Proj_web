package repo

import (
	"github.com/Cheep2Workshop/proj-web/models"
	"gorm.io/gorm/clause"
)

func (client *DbClient) AddOrder(userId int, items map[int]int) (models.Order, error) {
	var err error
	tx := client.Begin()
	order := models.Order{
		UserId: userId,
	}
	if err = tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return models.Order{}, err
	}

	details := make([]models.OrderDetail, len(items))
	i := 0
	for pid, count := range items {
		details[i] = models.OrderDetail{
			OrderId:       order.Id,
			ProductId:     pid,
			ProductAmount: count,
		}
		i++
	}
	// debug : diff of batch vs create
	if err = tx.Create(&details).Error; err != nil {
		tx.Rollback()
		return models.Order{}, err
	}
	tx.Commit()

	orders, err := client.GetOrders(order.Id)
	if err != nil {
		return models.Order{}, err
	}
	return orders[0], nil
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

// func (client *DbClient) DeleteOrder(oid int) error {
// 	return errors.New("")
// }
