package consumer

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

type HandlerFunc func(msg *sarama.ConsumerMessage)

type ConsumerGroupHandler struct {
	HandlerFunc HandlerFunc
}

func (h *ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		h.HandlerFunc(msg)
		session.MarkMessage(msg, "")
	}
	return nil
}

type KafkaConsumer struct {
	client sarama.ConsumerGroup
	ctx    context.Context
	cancel context.CancelFunc
}

type KafkaConsumerConfig struct {
	Brokers       []string
	GroupID       string
	Topics        []string
	HandlerFunc   HandlerFunc
	InitialOffset int64 // sarama.OffsetOldest or sarama.OffsetNewest
}

func NewKafkaConsumer(cfg KafkaConsumerConfig) (*KafkaConsumer, error) {
	saramaCfg := sarama.NewConfig()
	saramaCfg.Version = sarama.V2_8_0_0
	saramaCfg.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	saramaCfg.Consumer.Offsets.Initial = cfg.InitialOffset

	client, err := sarama.NewConsumerGroup(cfg.Brokers, cfg.GroupID, saramaCfg)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	c := &KafkaConsumer{
		client: client,
		ctx:    ctx,
		cancel: cancel,
	}

	go func() {
		handler := &ConsumerGroupHandler{
			HandlerFunc: cfg.HandlerFunc,
		}

		for {
			if err := client.Consume(ctx, cfg.Topics, handler); err != nil {
				log.Printf("[KafkaConsumer] Error consuming: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	return c, nil
}

func (c *KafkaConsumer) Shutdown() error {
	c.cancel()
	return c.client.Close()
}
