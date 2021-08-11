package test

import (
	"context"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/Cheep2Workshop/proj-web/dashredis"
	"github.com/Cheep2Workshop/proj-web/models/repo"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RedisSuite struct {
	suite.Suite
	Client *dashredis.RedisClient
	// Mock   redismock.ClientMock
}

func (t *RedisSuite) SetupSuite() {
	// mock
	// rdb, mock := redismock.NewClientMock()
	// t.Client = &dashredis.RedisClient{Client: rdb}
	// t.Mock = mock

	t.Client = dashredis.NewRedisClient()
}

var testVals = []interface{}{
	"Hello",
	1,
	1.2,
	repo.LoginLogRes{
		ID:        0,
		Name:      "Test",
		Email:     "Test@gmail.com",
		Admin:     false,
		CreatedAt: time.Now(),
	},
	// remove redis by user
	repo.LoginLogList{
		{
			ID:        1,
			Name:      "Test1",
			Email:     "Test1@gmail.com",
			Admin:     false,
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Test2",
			Email:     "Test2@gmail.com",
			Admin:     false,
			CreatedAt: time.Now(),
		},
	},
}

func (t *RedisSuite) TestSet() {
	for i, val := range testVals {
		status := t.Client.Client.Set(context.Background(), strconv.Itoa(i), val, 30*time.Minute)
		require.NoError(t.T(), status.Err())
	}
}

func (t *RedisSuite) TestGet() {
	for i := 0; i < len(testVals); i++ {
		result, err := t.Client.Client.Get(context.Background(), strconv.Itoa(i)).Result()
		require.NoError(t.T(), err)
		log.Println(result)
	}
}

func TestRedis(t *testing.T) {
	suite.Run(t, new(RedisSuite))
}
