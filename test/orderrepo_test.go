package test

import (
	"log"
	"testing"

	"github.com/Cheep2Workshop/proj-web/models"
	"github.com/Cheep2Workshop/proj-web/models/repo"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type OrderSuite struct {
	suite.Suite
	client *repo.DbClient
}

var (
	user = models.User{
		Name:     "Customer",
		Email:    "Customer@gmail.com",
		Password: "123456",
	}
	prods = []models.Product{
		{ProductName: "A", ProductDesc: "AAA", Price: 1},
		{ProductName: "B", ProductDesc: "BBB", Price: 3},
		{ProductName: "C", ProductDesc: "CCC", Price: 5},
	}
)

func (t *OrderSuite) SetupSuite() {
	client, err := testconfig.ConnectDb(false)
	if err != nil {
		log.Fatal(err)
	}
	// init client with tables and users
	client.Init()
	client.AutoMigrate(&models.Product{}, &models.Order{}, &models.OrderDetail{})
	t.client = client
	// signup customer
	client.Signup(user)
	customer, err := client.GetUserByEmail(user.Email)
	if err != nil {
		log.Fatal(err)
	}
	// add products
	client.AddProducts(prods...)

	// create order
	pids := make([]int, len(prods))
	for i, p := range prods {
		pids[i] = p.Id
	}
	amount := []int{2, 4, 6}
	req := repo.OrderReq{
		UserId:    customer.ID,
		ProductId: pids,
		Amount:    amount,
	}
	client.AddOrder(req)
}

func (t *OrderSuite) TearDownSuite() {
	t.client.Migrator().DropTable(&models.Product{}, &models.Order{}, &models.OrderDetail{})
}

func (t *OrderSuite) TestAddOrder() {
	customer, err := t.client.GetUserByEmail(user.Email)
	require.NoError(t.T(), err)
	pids := make([]int, len(prods))
	for i, p := range prods {
		pids[i] = p.Id
	}
	amount := []int{1, 2, 3}
	req := repo.OrderReq{
		UserId:    customer.ID,
		ProductId: pids,
		Amount:    amount,
	}
	err = t.client.AddOrder(req)
	require.NoError(t.T(), err)
}

func (t *OrderSuite) TestGetOrder() {
	orders, err := t.client.GetOrders(1)
	require.NoError(t.T(), err)
	for _, o := range orders {
		log.Printf("%v, total:%v\n", o, o.GetSumPrice())
		for _, d := range o.OrderDetails {
			log.Printf("%+v\n", d)
		}
	}
}

func (t *OrderSuite) TestDeleteProduct() {
	var err error
	p := models.Product{
		ProductName: "F",
		ProductDesc: "FFF",
		Price:       100,
	}
	err = t.client.AddProduct(&p)
	require.NoError(t.T(), err)
	err = t.client.DeleteProduct(p)
	require.NoError(t.T(), err)
	result, err := t.client.GetProduct(p.Id)
	require.Error(t.T(), err)
	log.Println(result)
}

func (t *OrderSuite) TestSetProduct() {
	var err error
	p := models.Product{
		ProductName: "G",
		ProductDesc: "GGG",
		Price:       10,
	}
	err = t.client.AddProduct(&p)
	require.NoError(t.T(), err)
	p.Price = 100
	err = t.client.SetProduct(&p)
	require.NoError(t.T(), err)
	result, err := t.client.GetProduct(p.Id)
	require.NoError(t.T(), err)
	require.Equal(t.T(), 100, result.Price)
}

func TestOrder(t *testing.T) {
	suite.Run(t, new(OrderSuite))
}
