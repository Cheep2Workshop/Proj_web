package test

import (
	"log"
	"testing"

	"github.com/Cheep2Workshop/proj-web/models/repo"
	"github.com/stretchr/testify/suite"
)

type OrderSuite struct {
	suite.Suite
	client *repo.DbClient
}

func (t *OrderSuite) SetupSuite() {
	client, err := testconfig.ConnectDb(false)
	if err != nil {
		log.Fatal(err)
	}
	// init client with tables and users
	client.Init()
	t.client = client
}

func (t *OrderSuite) TearDownSuite() {

}

func TestOrder(t *testing.T) {
	suite.Run(t, new(OrderSuite))
}
