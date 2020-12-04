package main

import (
	"github.com/go-redis/redis"
	"time"
)

var redisCli *redisClient

type redisClient struct {
	client *redis.Client
}

func GetRedisClient() *redisClient {
	return redisCli
}

func NewRedisOptions(redisAddr string, redisDB int) *redis.Options {
	return &redis.Options{
		Addr:        redisAddr,
		Password:    "",
		DB:          redisDB,
		DialTimeout: 10 * time.Second,
		ReadTimeout: 3 * time.Second,
		PoolSize:    5,
		PoolTimeout: 10 * time.Second,
	}
}

func InitRedisClient(options *redis.Options) (string, error) {
	client := redis.NewClient(options)
	pong, err := client.Ping().Result()
	if err != nil {
		return "", err
	}
	redisCli.client = client
	return pong, nil
}

