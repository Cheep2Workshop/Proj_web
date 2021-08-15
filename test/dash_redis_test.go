package test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/Cheep2Workshop/proj-web/dashredis"
	"github.com/Cheep2Workshop/proj-web/models/repo"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RedisSuite struct {
	suite.Suite
	Redis *dashredis.RedisClient
	// Mock   redismock.ClientMock
}

func (t *RedisSuite) SetupSuite() {
	var err error
	t.Redis = dashredis.NewRedisClient()
	size := 3
	testVals = make([]repo.LoginLogRes, size)
	for i := 0; i < size; i++ {
		subtime := time.Duration(i-size) * time.Minute
		testVals[i] = repo.LoginLogRes{
			ID:        i,
			Name:      testName,
			Email:     testEmail,
			Admin:     false,
			CreatedAt: time.Now().Add(subtime),
		}
	}
	err = t.Redis.SetLoginLogs(testEmail, testVals...)
	if err != nil {
		log.Fatalf("Failed to setup suite : %s\n", err.Error())
	}
}

func (t *RedisSuite) TearDownSuite() {
	// flush all redis data
	status := t.Redis.Client.FlushAll(context.Background())

	if status.Err() != nil {
		log.Printf("Failed to flush all : %s\n", status.Err().Error())
	}
}

var (
	testName  = "Test"
	testEmail = "Test@email.com"
	testVals  []repo.LoginLogRes
)

func (t *RedisSuite) TestSet() {
	err := t.Redis.SetLoginLogs(testEmail, testVals...)
	require.NoError(t.T(), err)
}

func (t *RedisSuite) TestGet() {
	result, err := t.Redis.GetLoginLogs(testEmail)
	require.NoError(t.T(), err)
	log.Println(result)
}

func (t *RedisSuite) TestDelete() {
	err := t.Redis.DeleteWithEmail(testEmail)
	require.NoError(t.T(), err)
	result, err := t.Redis.GetLoginLogs(testEmail)
	require.Nil(t.T(), result)
	require.Error(t.T(), err)
	require.Equal(t.T(), redis.Nil, err)
}

func (t *RedisSuite) TestSend() {
	channel := "test"
	msg := "HelloWorld!"
	w := make(chan int, 1)

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		sub := t.Redis.Client.Subscribe(ctx, channel)
		ch := sub.Channel()
		log.Println((<-ch).Payload)
		w <- 1
	}()
	time.Sleep(time.Second)
	cmd := t.Redis.Client.Publish(context.Background(), channel, msg)
	require.NoError(t.T(), cmd.Err())
	<-w
}

func TestRedis(t *testing.T) {
	suite.Run(t, new(RedisSuite))
}
