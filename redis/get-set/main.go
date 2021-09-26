package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-examples/common/config"
	"log"
)

const (
	key           = "my_test_key"
	value         = "my_test_value"
	nonExistedKey = "non_existed_key"
)

var ctx = context.Background()

func main() {
	opts, err := config.ParseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := config.NewConfig(opts.ConfigPath)
	if err != nil {
		log.Fatalf("Cannot access config: %v\n", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Server,
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := rdb.Set(ctx, key, value, 0).Err(); err != nil {
		log.Fatalln(err)
	}

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(key, val)

	val2, err := rdb.Get(ctx, nonExistedKey).Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(nonExistedKey, val2)
	}
}
