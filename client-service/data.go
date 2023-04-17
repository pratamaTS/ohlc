package main

import (
	"encoding/json"
	"log"
	"my-project/ohlc-service/client-service/kafka"
	"my-project/ohlc-service/service/redis"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

func GetDataSummary() (resp map[string]any, err error) {
	kafkaConfig := kafka.GetKafkaConfig("", "")
	log.Print("kafka host ", os.Getenv("KAFKA_HOST"))
	consumers, err := sarama.NewConsumer([]string{os.Getenv("KAFKA_HOST")}, kafkaConfig)
	if err != nil {
		return
	}

	defer func() {
		if err := consumers.Close(); err != nil {
			log.Print(err)
		}
	}()

	kafka := &kafka.KafkaConsumer{
		Consumer: consumers,
	}

	signals := make(chan os.Signal, 1)
	go kafka.Consume([]string{"ohlc-notif"}, signals)
	timeout := time.After(2 * time.Second)
	for {
		select {
		case <-timeout:
			signals <- os.Interrupt
			return
		}
	}
	return
}

func GetDataSummeryWithRedis() (resp []map[string]any, err error) {
	redisClient, err := redis.GetRedisClient()
	if err != nil {
		return
	}

	resp = make([]map[string]any, 0)
	mapRes := map[string]any{}

	respA, err := redisClient.LRange("ohlc-result", 0, -1).Result()
	if err != nil {
		return
	}

	for _, v := range respA {
		err = json.Unmarshal([]byte(v), &mapRes)
		if err != nil {
			return
		}

		resp = append(resp, mapRes)
	}

	jsonData, _ := json.MarshalIndent(resp, "", " ")
	log.Print("resp ", string(jsonData))
	return
}
