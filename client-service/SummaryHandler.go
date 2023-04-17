package main

import (
	"my-project/ohlc-service/service/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SummaryHandler(c *gin.Context) {
	resp, err := GetDataSummeryWithRedis()
	if err != nil {
		response.Response(c, http.StatusInternalServerError, "", err)
		return
	}
	response.ResponseData(c, http.StatusOK, "Success get data summary", nil, resp)
}
