package kafka

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	Consumer sarama.Consumer
}

func (c *KafkaConsumer) Consume(topics []string, signals chan os.Signal) {
	chanMessage := make(chan *sarama.ConsumerMessage, 256)

	for _, topic := range topics {
		partitionList, err := c.Consumer.Partitions(topic)
		if err != nil {
			logrus.Errorf("Unable to get partition got error %v", err)
			continue
		}
		log.Print("before partition ")
		for _, partition := range partitionList {
			log.Print("partition ", partition)
			go consumeMessage(c.Consumer, topic, partition, chanMessage)
		}
	}
	log.Print("Kafka is consuming....")

ConsumerLoop:
	for {
		select {
		case msg := <-chanMessage:
			log.Print("New Message from kafka, message: %v", string(msg.Value))
		case sig := <-signals:
			if sig == os.Interrupt {
				break ConsumerLoop
			}
		}
	}
}

func consumeMessage(consumer sarama.Consumer, topic string, partition int32, c chan *sarama.ConsumerMessage) {
	msg, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		logrus.Errorf("Unable to consume partition %v got error %v", partition, err)
		return
	}

	defer func() {
		if err := msg.Close(); err != nil {
			logrus.Errorf("Unable to close partition %v: %v", partition, err)
		}
	}()

	for {
		msg := <-msg.Messages()
		log.Print("message ", msg)
		c <- msg
	}

}
