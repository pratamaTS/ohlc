package response

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, statusCode int, message string, err error) {
	resp := map[string]any{
		"error":   false,
		"message": message,
	}

	if err != nil {
		resp["error"] = true
		resp["message"] = err.Error()
	}

	c.IndentedJSON(statusCode, resp)
	return
}
