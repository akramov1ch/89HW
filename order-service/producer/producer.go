package producer

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type OrderProducer struct {
	Writer *kafka.Writer
}

func NewOrderProducer(brokers []string, topic string) *OrderProducer {
	return &OrderProducer{
		Writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *OrderProducer) Produce(key string, message []byte) error {
	err := p.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: message,
		},
	)
	return err
}

func (p *OrderProducer) Close() {
	p.Writer.Close()
}
