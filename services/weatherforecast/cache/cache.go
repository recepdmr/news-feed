package cache

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func init() {
	redisConnectionString := os.Getenv("REDIS_CONNECTION_STRING")

	opt, err := redis.ParseURL(redisConnectionString)
	if err != nil {
		panic(err)
	}
	RedisClient = redis.NewClient(opt)
}
