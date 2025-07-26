package event

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Writer *kafka.Writer
}

func NewProducer(broker, topic string) *Producer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &Producer{Writer: writer}
}

func (p *Producer) Send(msg string) error {
	return p.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(time.Now().Format(time.RFC3339Nano)),
			Value: []byte(msg),
		})
}

func (p *Producer) Close() {
	_ = p.Writer.Close()
}
