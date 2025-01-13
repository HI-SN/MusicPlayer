package database

import (
	"context"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// init redis
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64) //可不填
	client := redis.NewClient(&redis.Options{
		Addr:       "127.0.0.1:6379", // Redis默认地址
		Password:   "",               // Redis密码，一般建立时不设置
		DB:         int(db),
		MaxRetries: 1,
	})

	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		panic("can't connect redis")
	}

	RedisClient = client
}
