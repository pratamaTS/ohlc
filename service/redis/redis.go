package redis

import (
	"errors"
	"log"
	"os"

	"github.com/go-redis/redis"
)

var initializeRedis = false

var (
	redisClient *redis.Client
)

func InitRedis() error {
	if !initializeRedis {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})
		_, err := redisClient.Ping().Result()
		if err != nil {
			log.Println("Failed connect to Redis")
			return err
		}

		initializeRedis = true
	}
	return nil
}

func GetRedisClient() (*redis.Client, error) {
	if initializeRedis == false || redisClient == nil {
		err := InitRedis()
		if err != nil {
			return nil, errors.New("Failed to initialize Redis")
		}
	}
	return redisClient, nil
}
