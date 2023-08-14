package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

type RedisClient struct {
	Client *redis.Client
}

func (redisClient *RedisClient) StoreToRedistWithExpired(key string, val interface{}, duration string) error {
	dur, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}

	b, err := json.Marshal(val)
	if err != nil {
		return err
	}

	err = redisClient.Client.Set(key, string(b), dur).Err()

	return err
}


func (redisClient *RedisClient) StoreToRedis(key string, val interface{}) error {
	b, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = redisClient.Client.Set(key, string(b), 0).Err()

	return err
}

func (redisClient *RedisClient) GetFromRedis(key string, cb interface{}) error {
	res, err := redisClient.Client.Get(key).Result()
	if err != nil {
		return err
	}

	if res == "" {
		return errors.New("[Redis] Value of " + key + " is empty.")
	}

	err = json.Unmarshal([]byte(res), &cb)
	if err != nil {
		return err
	}

	return err
}

func (redisClient *RedisClient) RemoveFromRedis(key string) error {
	return redisClient.Client.Del(key).Err()
}

func (redisClient *RedisClient) Get(key string) ([]byte, error) {
    res, err := redisClient.Client.Get(key).Result()
    if err != nil {
        return nil, err
    }
    return []byte(res), nil
}

func (redisClient *RedisClient) Set(key string, value []byte, expiration time.Duration) error {
    return redisClient.Client.Set(key, string(value), expiration).Err()
}


func (redisClient *RedisClient) Reset() error {
    return redisClient.Client.FlushDB().Err()
}

func (redisClient *RedisClient) Close() error {
    return redisClient.Client.Close()
}

func (redisClient *RedisClient) Delete(key string) error {
    return redisClient.Client.Del(key).Err()
}
