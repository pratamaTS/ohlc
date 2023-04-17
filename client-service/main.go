package main

import (
	"log"
	"my-project/ohlc-service/service/redis"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	redis.InitRedis()

	router := gin.Default()

	router.GET("/health", HealthHandler)
	router.GET("/get-summary", SummaryHandler)

	log.Print("Starting service")
	router.Run(":" + os.Getenv("CLIENT_PORT"))
}
