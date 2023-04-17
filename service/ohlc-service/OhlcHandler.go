package ohlcservice

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"my-project/ohlc-service/service/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func OhlcHandler(c *gin.Context) {
	code := []string{
		"BBCA", "BBRI", "ASII", "GOTO",
	}

	groupData := []SubSetData{}
	data := SubSetData{}
	file, err := ioutil.ReadFile("ohlc-service/testcase-data/2022-11-10-1668045608768580708.ndjson")
	if err != nil {
		response.Response(c, http.StatusInternalServerError, "", err)
		return
	}
	d := json.NewDecoder(strings.NewReader(string(file)))
	for {
		var v interface{}
		err := d.Decode(&v)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		jsonData, _ := json.Marshal(v)

		// Convert the JSON to a struct
		json.Unmarshal(jsonData, &data)
		groupData = append(groupData, data)
	}

	_, err = Ohlc(code, groupData)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, "", err)
		return
	}

	response.Response(c, http.StatusOK, "Success calculate ohlc", nil)
}
