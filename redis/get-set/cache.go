package main

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	rdb *redis.Client
}

func NewCache(cfg *Redis) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Server,
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &Cache{rdb: rdb}
}

func (c *Cache) SetValue(key string, value interface{}) error {
	return c.rdb.Set(ctx, key, value, time.Minute).Err()
}

func (c *Cache) GetValue(key string) (string, error) {
	return c.rdb.GetDel(ctx, key).Result()
}
