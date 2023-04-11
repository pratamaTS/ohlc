package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	router.GET("/health", HealthHandler)

	log.Print("Starting service")
	router.Run(":" + os.Getenv("PORT"))
}
