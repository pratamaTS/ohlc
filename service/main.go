package main

import (
	"log"
	"my-project/ohlc-service/service/kafka"
	ohlcservice "my-project/ohlc-service/service/ohlc-service"
	"my-project/ohlc-service/service/redis"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	redis.InitRedis()
	kafka.InitKafka()

	router := gin.Default()

	router.GET("/health", HealthHandler)
	router.GET("/ohlc", ohlcservice.OhlcHandler)

	log.Print("Starting service")
	router.Run(":" + os.Getenv("PORT"))
}
