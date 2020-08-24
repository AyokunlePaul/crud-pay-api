package redis_client

import (
	"github.com/go-redis/redis/v8"
	"os"
)

const redisDnsConstant = "REDIS_DNS"
const redisPasswordConstant = "REDIS_PASSWORD"

var redisClient *redis.Client

func init() {
	redisDNS := os.Getenv(redisDnsConstant)

	redisClient = redis.NewClient(&redis.Options{
		Addr: redisDNS,
		DB:   0,
	})
}

func Get() *redis.Client {
	return redisClient
}
