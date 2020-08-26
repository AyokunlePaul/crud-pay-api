package redis_client

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
)

const redisName = "REDIS_CONTAINER_NAME"

var redisClient *redis.Client

func init() {
	redisName := os.Getenv(redisName)

	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisName, "6379"),
		DB:   0,
		MaxRetries: 5,
	})
	_, pingResult := redisClient.Ping(context.Background()).Result()
	if pingResult != nil {
		panic(pingResult)
	}
}

func Get() *redis.Client {
	return redisClient
}
