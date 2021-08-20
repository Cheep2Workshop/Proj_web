package repo

import (
	"errors"

	"github.com/Cheep2Workshop/proj-web/models"
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

	return nil
}
