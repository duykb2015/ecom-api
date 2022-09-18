package redis

import (
	"github.com/duykb2015/ecom-api/config"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(cfg *config.RedisUC) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addrs,
		Password: "",
		DB:       0,
		PoolSize: cfg.PoolSize,
	})
	err := client.Ping(client.Context()).Err()
	if err != nil {
		return nil, err
	}
	return client, err
}
