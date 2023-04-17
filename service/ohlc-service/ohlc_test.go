package ohlcservice_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	ohlcservice "my-project/ohlc-service/service/ohlc-service"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOhlcReadRawData(t *testing.T) {
	log.Print("== RUN TestLogin")
	file, err := ioutil.ReadFile("testcase-data/2022-11-10-1668045608768580708.ndjson")

	assert.NoError(t, err)
	assert.NotEmpty(t, file)
}

func TestOhlc(t *testing.T) {
	log.Print("== RUN TestLogin")
	file, err := ioutil.ReadFile("testcase-data/2022-11-10-1668045608768580708.ndjson")
	assert.NoError(t, err)
	assert.NotEmpty(t, file)

	code := []string{
		"BBCA", "BBRI", "ASII", "GOTO",
	}

	groupData := []ohlcservice.SubSetData{}
	data := ohlcservice.SubSetData{}
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
		fmt.Println(data)
		groupData = append(groupData, data)
	}

	log.Print("len group data ", groupData)
	result, err := ohlcservice.Ohlc(code, groupData)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	jsonDataResp, _ := json.MarshalIndent(result, "", " ")
	log.Print("result data ", string(jsonDataResp))
}

// func TestOhlcOri(t *testing.T) {
// 	x := []string{
// 		"BBCA", "BBRI", "ASII", "GOTO",
// 	}
// 	w := []ohlcservice.ChangeRecord{
// 		{
// 			StockCode: "BBCA",
// 			Price:     8783,
// 			Quantity:  0,
// 		},
// 		{
// 			StockCode: "BBRI",
// 			Price:     3233,
// 			Quantity:  0,
// 		},
// 		{
// 			StockCode: "ASII",
// 			Price:     1223,
// 			Quantity:  0,
// 		},
// 		{
// 			StockCode: "GOTO",
// 			Price:     321,
// 			Quantity:  0,
// 		},

// 		{
// 			StockCode: "BBCA",
// 			Price:     8780,
// 			Quantity:  1,
// 		},
// 		{
// 			StockCode: "BBRI",
// 			Price:     3230,
// 			Quantity:  1,
// 		},
// 		{
// 			StockCode: "ASII",
// 			Price:     1220,
// 			Quantity:  1,
// 		},
// 		{
// 			StockCode: "GOTO",
// 			Price:     320,
// 			Quantity:  1,
// 		},

// 		{
// 			StockCode: "BBCA",
// 			Price:     8800,
// 			Quantity:  1,
// 		},
// 		{
// 			StockCode: "BBRI",
// 			Price:     3300,
// 			Quantity:  1,
// 		},
// 		{
// 			StockCode: "ASII",
// 			Price:     1300,
// 			Quantity:  1,
// 		},
// 		{
// 			StockCode: "GOTO",
// 			Price:     330,
// 			Quantity:  1,
// 		},

// 		{
// 			StockCode: "BBCA",
// 			Price:     8600,
// 			Quantity:  1,
// 		},
// 		{
// 			StockCode: "BBRI",
// 			Price:     3100,
// 			Quantity:  1,
// 		},
// 		{
// 			StockCode: "ASII",
// 			Price:     1100,
// 			Quantity:  1,
// 		},
// 		{
// 			StockCode: "GOTO",
// 			Price:     310,
// 			Quantity:  1,
// 		},

// 		{
// 			StockCode: "BBCA",
// 			Price:     8785,
// 			Quantity:  1,
// 		},
// 		{
// 			StockCode: "BBRI",
// 			Price:     3235,
// 			Quantity:  1,
// 		},
// 		{
// 			StockCode: "ASII",
// 			Price:     1225,
// 			Quantity:  1,
// 		},
// 		{
// 			StockCode: "GOTO",
// 			Price:     325,
// 			Quantity:  1,
// 		},
// 	}
// 	p := []ohlcservice.IndexMember{
// 		{
// 			StockCode: "BBCA",
// 			IndexCode: "IHSG",
// 		},
// 		{
// 			StockCode: "BBRI",
// 			IndexCode: "IHSG",
// 		},
// 		{
// 			StockCode: "ASII",
// 			IndexCode: "IHSG",
// 		},
// 		{
// 			StockCode: "GOTO",
// 			IndexCode: "IHSG",
// 		},
// 		{
// 			StockCode: "BBCA",
// 			IndexCode: "LQ45",
// 		},
// 		{
// 			StockCode: "BBRI",
// 			IndexCode: "LQ45",
// 		},
// 		{
// 			StockCode: "ASII",
// 			IndexCode: "LQ45",
// 		},
// 		{
// 			StockCode: "GOTO",
// 			IndexCode: "LQ45",
// 		},
// 		{
// 			StockCode: "BBCA",
// 			IndexCode: "KOMPAS100",
// 		},
// 		{
// 			StockCode: "BBRI",
// 			IndexCode: "KOMPAS100",
// 		},
// 	}
// 	r := ohlcservice.OhlcOri(x, w, p)
// 	assert.NotEmpty(t, r)
// 	for _, v := range r {
// 		jss, _ := json.MarshalIndent(v, "", " ")
// 		fmt.Println("summary: ", string(jss))
// 	}
// }
