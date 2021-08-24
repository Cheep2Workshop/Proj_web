package service

import (
	"errors"

	"github.com/Cheep2Workshop/proj-web/models"
	"github.com/Cheep2Workshop/proj-web/models/repo"
)

type OrderReq struct {
	UserId    int
	ProductId []int
	Amount    []int
}

func AddOrder(req OrderReq) (models.Order, error) {
	if len(req.ProductId) != len(req.Amount) {
		return models.Order{}, errors.New("length of product and amount not matched.")
	}
	items := map[int]int{}
	for i, pid := range req.ProductId {
		items[pid] = req.Amount[i]
	}

	order, err := repo.Client.AddOrder(req.UserId, items)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}
