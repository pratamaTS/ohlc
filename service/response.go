package main

import "github.com/gin-gonic/gin"

func SuccessResponse(c *gin.Context, statusCode int, message string) {
	resp := map[string]any{
		"error":   false,
		"message": message,
	}
	c.IndentedJSON(statusCode, resp)
	return
}
