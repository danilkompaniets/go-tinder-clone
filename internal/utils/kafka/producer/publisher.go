package producer

import (
	"github.com/IBM/sarama"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka"
)

type Publisher struct {
	producer *KafkaProducer
}

func NewPublisher(producer *KafkaProducer) *Publisher {
	return &Publisher{producer: producer}
}

func (p *Publisher) PublishEvent(eventType kafka.EventType, key string, value []byte) {
	headers := []sarama.RecordHeader{
		{
			Key:   []byte("type"),
			Value: []byte(eventType),
		},
	}

	p.producer.SendMessage(key, value, headers)
}
