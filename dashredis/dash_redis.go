package dashredis

import (
	"context"
	"encoding/json"
	"time"

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

func (client *RedisClient) SetLoginLogs(email string, logs ...orm.LoginLogRes) error {
	list := orm.LoginLogList(logs)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status := client.Client.Set(ctx, email, list, 0)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

// GetLoginLogs will return nil error while key not found
func (client *RedisClient) GetLoginLogs(email string) (orm.LoginLogList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := client.Client.Get(ctx, email).Result()
	if err != nil {
		return nil, err
	}
	var logs orm.LoginLogList
	err = json.Unmarshal([]byte(result), &logs)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (client *RedisClient) DeleteWithEmail(email string) error {
	cmd := client.Client.Del(context.Background(), email)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}
