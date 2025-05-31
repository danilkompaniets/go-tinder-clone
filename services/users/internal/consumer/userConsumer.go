package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka"
	kafkautil "github.com/danilkompanites/tinder-clone/internal/utils/kafka/consumer"
	"github.com/danilkompanites/tinder-clone/services/users/internal/repository/sql"
	"github.com/danilkompanites/tinder-clone/services/users/pkg/model"
)

type UserConsumer struct {
	kafkaConsumer *kafkautil.KafkaConsumer
	repo          *sql.Repository
}

func NewUserConsumer(brokers []string, repo *sql.Repository) (*UserConsumer, error) {
	handlerFunc := func(msg *sarama.ConsumerMessage) {
		var eventType kafka.EventType
		for _, h := range msg.Headers {
			if string(h.Key) == "type" {
				eventType = kafka.EventType(h.Value)
				break
			}
		}

		switch eventType {
		case kafka.Events.UserCreated:
			var user model.User
			if err := json.Unmarshal(msg.Value, &user); err != nil {
				log.Printf("Failed to unmarshal: %v", err)
				return
			}
			if err := repo.InsertUser(context.Background(), user); err != nil {
				log.Printf("Failed to insert user: %v", err)
			}
		default:
			log.Printf("Unknown event type: %s", eventType)
		}
	}

	kafkaC, err := kafkautil.NewKafkaConsumer(kafkautil.KafkaConsumerConfig{
		Brokers:       brokers,
		GroupID:       string(kafka.ConsumerGroups.UserService),
		Topics:        []string{string(kafka.Topics.User)},
		HandlerFunc:   handlerFunc,
		InitialOffset: sarama.OffsetOldest,
	})
	if err != nil {
		return nil, err
	}

	return &UserConsumer{
		kafkaConsumer: kafkaC,
		repo:          repo,
	}, nil
}

func (c *UserConsumer) Shutdown() error {
	return c.kafkaConsumer.Shutdown()
}
