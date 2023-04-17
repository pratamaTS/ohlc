package kafka_test

import (
	"fmt"
	"my-project/ohlc-service/service/kafka"
	"testing"

	"github.com/Shopify/sarama/mocks"
)

func TestSendMessage(t *testing.T) {
	t.Run("Send message OK", func(t *testing.T) {
		// membuat producer mock
		mockedProducer := mocks.NewSyncProducer(t, nil)
		// membuat expect producer success atau berhasil mengirim pesan
		mockedProducer.ExpectSendMessageAndSucceed()
		kafka := &kafka.KafkaProducer{
			Producer: mockedProducer,
		}

		msg := "Message 1"

		err := kafka.SendMessage("test_topic", msg)
		if err != nil {
			t.Errorf("Send message should not be error but have: %v", err)
		}
	})

	t.Run("Send message NOK", func(t *testing.T) {
		mockedProducer := mocks.NewSyncProducer(t, nil)
		// membuat producer gagal mengirim pesan
		mockedProducer.ExpectSendMessageAndFail(fmt.Errorf("Error"))
		kafka := &kafka.KafkaProducer{
			Producer: mockedProducer,
		}

		msg := "Message 1"

		err := kafka.SendMessage("test_topic", msg)
		if err == nil {
			t.Error("this should be error")
		}
	})
}
