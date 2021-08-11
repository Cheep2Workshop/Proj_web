package dashredis

import (
	"context"

	orm "github.com/Cheep2Workshop/proj-web/models/repo"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &RedisClient{rdb}
}

func (client *RedisClient) SetLoginLogs(email string, logs []orm.LoginLogRes) error {
	status := client.Client.Set(context.Background(), email, logs, 0)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}
