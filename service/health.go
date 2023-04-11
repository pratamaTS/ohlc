package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	SuccessResponse(c, http.StatusOK, "OK")
}
