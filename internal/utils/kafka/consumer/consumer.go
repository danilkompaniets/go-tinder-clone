package consumer

import (
	"github.com/IBM/sarama"
	"log"
)

type GroupHandler struct {
	HandlerFunc func(message *sarama.ConsumerMessage)
}

func (h *GroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *GroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h *GroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("[KafkaConsumer] Message: %s", string(msg.Value))
		h.HandlerFunc(msg)
		session.MarkMessage(msg, "")
	}

	return nil
}
