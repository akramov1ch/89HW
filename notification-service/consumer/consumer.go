package consumer

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type OrderConsumer struct {
	Reader *kafka.Reader
}

func NewOrderConsumer(brokers []string, topic string) *OrderConsumer {
	return &OrderConsumer{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: "notification-service",
		}),
	}
}

func (c *OrderConsumer) Consume() (kafka.Message, error) {
	return c.Reader.ReadMessage(context.Background())
}

func (c *OrderConsumer) Close() {
	c.Reader.Close()
}
