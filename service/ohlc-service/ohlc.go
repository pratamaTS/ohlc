package ohlcservice

import (
	"encoding/json"
	"fmt"
	"log"
	"my-project/ohlc-service/service/kafka"
	"my-project/ohlc-service/service/redis"
	"os"
	"strconv"

	"github.com/Shopify/sarama"
)

type SubSetData struct {
	Type      string `json:"type"`
	OrderBook string `json:"order_book"`
	Price     string `json:"price"`
	StockCode string `json:"stock_code"`
	Quantity  string `json: quantity`
}

type ChangeRecord struct {
	StockCode string
	Price     int64
	Quantity  int64
}

type IndexMember struct {
	StockCode string
	IndexCode string
}

type Summary struct {
	StockCode string   `json:"stock_code"`
	IndexCode []string `json:"index_code"`
	Open      int64    `json:"open"`
	High      int64    `json:"high"`
	Low       int64    `json:"low"`
	Close     int64    `json:"close"`
	Prev      int64    `json:"prev"`
	Volume    int64    `json:"volume"`
	Value     int64    `json:"value"`
	AvgPrice  int64    `json:"avg_price"`
}

var temp = Summary{}

func Ohlc(code []string, data []SubSetData) (result map[string]Summary, err error) {
	result = map[string]Summary{}

	for _, c := range code {
		if _, found := result[c]; !found {
			temp = Summary{}
		}

		temp.StockCode = c
		for _, v := range data {
			if v.StockCode != c {
				continue
			}

			price, err1 := strconv.ParseInt(v.Price, 10, 64)
			if err1 != nil {
				err = err1
				return
			}

			log.Print("qty before ", v.Quantity)

			qty, err1 := strconv.ParseInt(v.Quantity, 10, 64)
			if err1 != nil {
				err = err1
				return
			}

			log.Print("qty ", qty)
			switch {
			case qty == 0:
				temp.Prev = price
				result[c] = temp
				log.Print("prev price updated")
				continue
			case qty > 0 && result[c].Open == 0:
				temp.Open = price
				temp.Value += qty * temp.Open
				result[c] = temp
				log.Print("open price updated")
				continue
			default:
				temp.Close = price
				if temp.High < price {
					temp.High = price
					temp.Value += qty * temp.High
					log.Print("high price updated")
				}
				if temp.Low > price {
					temp.Low = price
					temp.Value += qty * temp.Low
					log.Print("low price updated")
				}

				log.Print("close price updated")
				result[c] = temp
			}
			temp.Volume += qty
			if result[c].Close != 0 {
				temp.AvgPrice = temp.Value / temp.Volume
			}
			result[c] = temp
			log.Print("volume ", temp.Volume)
			log.Print("value ", temp.Value)
			log.Print("avg price ", temp.AvgPrice)
		}
	}

	redisClient, err := redis.GetRedisClient()
	if err != nil {
		return
	}

	kafkaConfig := kafka.GetKafkaConfig("", "")
	log.Print("kafka host ", os.Getenv("KAFKA_HOST"))
	producers, err := sarama.NewSyncProducer([]string{os.Getenv("KAFKA_HOST")}, kafkaConfig)
	if err != nil {
		fmt.Errorf("Unable to create kafka producer got error %v", err)
		return
	}
	defer func() {
		if err := producers.Close(); err != nil {
			fmt.Errorf("Unable to stop kafka producer: %v", err)
			return
		}
	}()

	log.Print("Success create kafka sync-producer")

	produce := &kafka.KafkaProducer{
		Producer: producers,
	}

	for _, v := range result {
		jsonDataResp, _ := json.Marshal(v)
		err = redisClient.LPush("ohlc-result", jsonDataResp).Err()
		if err != nil {
			return
		}

		err = produce.SendMessage("ohlc-notif", string(jsonDataResp))
		if err != nil {
			err = fmt.Errorf("Send message should not be error but have: %v", err)
			return
		}
	}

	return
}
