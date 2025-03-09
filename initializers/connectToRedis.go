package initializers

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func ConnectToRedis() {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})
}
