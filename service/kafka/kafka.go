package kafka

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

var initializeKafka = false

var ProducerInit = &KafkaProducer{}

func InitKafka() {
	kafkaConfig := GetKafkaConfig("", "")
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

	ProducerInit = &KafkaProducer{
		Producer: producers,
	}
	initializeKafka = true
	return
}

func GetKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}

func GetKafkaClient() *KafkaProducer {
	if initializeKafka == false || ProducerInit == nil {
		InitKafka()
	}
	return ProducerInit
}
