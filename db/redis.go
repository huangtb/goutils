package db

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"time"
)

var RedisCli *redis.Client

func NewRedisOptions(redisAddr string, redisDB int) *redis.Options {
	return &redis.Options{
		Addr:        redisAddr,
		Password:    "",
		DB:          redisDB,
		DialTimeout: 3 * time.Second,
		ReadTimeout: 3 * time.Second,
		PoolSize:    5,
		PoolTimeout: 3 * time.Second,
	}
}

func InitRedisClient(options *redis.Options) (string, error) {
	client := redis.NewClient(options)
	pong, err := client.Ping().Result()
	if err != nil {
		return "", errors.Errorf("Init redis client error:%s", err.Error())
	}
	RedisCli = client
	return pong, nil
}
