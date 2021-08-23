package test

import (
	"log"
	"testing"

	"github.com/Cheep2Workshop/proj-web/models"
	"github.com/Cheep2Workshop/proj-web/models/repo"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ProductSuite struct {
	suite.Suite
	client *repo.DbClient
}

var (
	products = []models.Product{
		{ProductName: "A", ProductDesc: "AAA", Price: 1},
		{ProductName: "B", ProductDesc: "BBB", Price: 2},
		{ProductName: "C", ProductDesc: "CCC", Price: 3},
	}
)

func (t *ProductSuite) SetupSuite() {
	client, err := testconfig.ConnectDb(false)
	if err != nil {
		log.Fatal(err)
	}
	// init client with tables and users
	client.Init()
	client.AutoMigrate(&models.Product{})
	t.client = client
	t.client.AddProducts(products...)
}

func (t *ProductSuite) TearDownSuite() {
	t.client.Migrator().DropTable(&models.Product{})
}

func (t *ProductSuite) TestAddProduct() {
	p := models.Product{
		ProductName: "D",
		ProductDesc: "DDD",
		Price:       100,
	}
	err := t.client.AddProduct(&p)
	log.Println(p)
	require.NoError(t.T(), err)
	get, err := t.client.GetProduct(p.Id)
	require.NoError(t.T(), err)
	require.Equal(t.T(), p, get)
}

func (t *ProductSuite) TestAddProducts() {
	ps := []models.Product{
		{ProductName: "E", ProductDesc: "EEE", Price: 100},
		{ProductName: "F", ProductDesc: "FFF", Price: 101},
	}
	err := t.client.AddProducts(ps...)
	for _, p := range ps {
		log.Println(p)
	}
	require.NoError(t.T(), err)
	gets, err := t.client.GetProducts([]int{ps[0].Id, ps[1].Id}...)
	require.NoError(t.T(), err)
	for i, get := range gets {
		require.Equal(t.T(), ps[i], get)
	}
}

func (t *ProductSuite) TestUpdateProduct() {
	p, err := t.client.GetProduct(1)
	require.NoError(t.T(), err)
	log.Println(p)
	p.Price += 10
	t.client.UpdateProduct(&p)
	log.Println(p)
	pp, err := t.client.GetProduct(1)
	require.NoError(t.T(), err)
	require.Equal(t.T(), p, pp)
}

func (t *ProductSuite) TestGetProduct() {
	products, err := t.client.GetAllProducts()
	require.NoError(t.T(), err)
	for _, p := range products {
		log.Println(p)
	}
}

func (t *ProductSuite) TestDeleteProductById() {
	err := t.client.DeleteProductById(2)
	require.NoError(t.T(), err)
	_, err = t.client.GetProduct(2)
	require.Error(t.T(), err)
}

func (t *ProductSuite) TestDeleteProduct() {
	p := models.Product{
		ProductName: "G",
		ProductDesc: "GGG",
		Price:       0,
	}
	err := t.client.AddProduct(&p)
	require.NoError(t.T(), err)
	err = t.client.DeleteProduct(p)
	require.NoError(t.T(), err)
	_, err = t.client.GetProduct(p.Id)
	require.Error(t.T(), err)
}

func TestProductRepo(t *testing.T) {
	suite.Run(t, new(ProductSuite))
}
