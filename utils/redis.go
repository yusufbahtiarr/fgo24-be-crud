package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Redis() error {
	godotenv.Load()

	redisAddr := os.Getenv("RDADDRESS")
	redisPassword := os.Getenv("RDPASSWORD")
	redisDB, _ := strconv.Atoi(os.Getenv("RDDB"))

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	return nil
}
