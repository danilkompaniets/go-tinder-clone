package producer

import (
	"github.com/IBM/sarama"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka"
	"log"
)

type KafkaProducer struct {
	AsyncProducer sarama.AsyncProducer
	Topic         string
}

func NewKafkaProducer(brokers []string, topic kafka.Topic) (*KafkaProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Retry.Max = 5

	producer, err := sarama.NewAsyncProducer(brokers, cfg)
	if err != nil {
		return nil, err
	}

	kp := &KafkaProducer{
		AsyncProducer: producer,
		Topic:         string(topic),
	}

	go func() {
		for err := range producer.Errors() {
			log.Printf("[KafkaProducer] Error: %v", err)
		}
	}()

	return kp, nil
}

func (p *KafkaProducer) SendMessage(key string, value []byte, headers []sarama.RecordHeader) {
	msg := &sarama.ProducerMessage{
		Topic:   p.Topic,
		Key:     sarama.StringEncoder(key),
		Value:   sarama.ByteEncoder(value),
		Headers: headers,
	}

	p.AsyncProducer.Input() <- msg
}
