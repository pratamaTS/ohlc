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

func ResponseData(c *gin.Context, statusCode int, message string, err error, data []map[string]any) {
	resp := map[string]any{
		"error":   false,
		"message": message,
	}

	if err != nil {
		resp["error"] = true
		resp["message"] = err.Error()
	}

	if len(data) > 0 {
		resp["data"] = data
	}

	c.IndentedJSON(statusCode, resp)
	return
}
