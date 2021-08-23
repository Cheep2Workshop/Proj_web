package repo

import (
	"errors"

	"github.com/Cheep2Workshop/proj-web/models"
)

func (client *DbClient) AddProduct(p *models.Product) error {
	var err error
	if err = checkProduct(p); err != nil {
		return err
	}
	err = client.Create(&p).Error
	return err
}

func (client *DbClient) AddProducts(ps ...models.Product) error {
	var err error
	for _, p := range ps {
		if err = checkProduct(&p); err != nil {
			return err
		}
	}
	err = client.Debug().Create(&ps).Error
	// err = client.Debug().CreateInBatches(&ps, len(ps)).Error
	return err
}

func (client *DbClient) UpdateProduct(p *models.Product) error {
	var err error
	if err = checkProduct(p); err != nil {
		return err
	}
	err = client.Save(&p).Error
	return err
}

func (client *DbClient) DeleteProductById(pids ...int) error {
	err := client.Model(&models.Product{}).Delete(&models.Product{}, pids).Error
	return err
}

func (client *DbClient) DeleteProduct(p models.Product) error {
	err := client.Delete(&p).Error
	return err
}

func (client *DbClient) GetProduct(pid int) (models.Product, error) {
	var p models.Product
	err := client.Model(&models.Product{}).First(&p, pid).Error
	return p, err
}

func (client *DbClient) GetProducts(pids ...int) ([]models.Product, error) {
	ps := make([]models.Product, len(pids))
	err := client.Find(&ps, pids).Error
	return ps, err
}

func (client *DbClient) GetAllProducts() ([]models.Product, error) {
	var err error
	var products []models.Product
	err = client.Model(&models.Product{}).Scan(&products).Error
	return products, err
}

func (client *DbClient) SetProduct(p *models.Product) error {
	err := client.Save(p).Error
	return err
}

func checkProduct(p *models.Product) error {
	if len(p.ProductName) == 0 {
		return errors.New("Product name is empty.")
	}
	if p.Price < 0 {
		return errors.New("Product price is negative.")
	}
	return nil
}
