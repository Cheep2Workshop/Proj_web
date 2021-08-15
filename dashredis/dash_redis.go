package dashredis

import (
	"context"
	"encoding/json"

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
	status := client.Client.Set(context.Background(), email, list, 0)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

// GetLoginLogs will return nil error while key not found
func (client *RedisClient) GetLoginLogs(email string) (orm.LoginLogList, error) {
	result, err := client.Client.Get(context.Background(), email).Result()
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
