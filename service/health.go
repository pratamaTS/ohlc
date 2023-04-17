package main

import (
	"my-project/ohlc-service/service/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	response.Response(c, http.StatusOK, "OK", nil)
}
