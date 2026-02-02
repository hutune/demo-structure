package kafka

import (
	"context"
)

// Handler processes a Kafka message
type Handler func(ctx context.Context, key, value []byte) error

// Subscriber subscribes to Kafka topics
type Subscriber interface {
	Subscribe(ctx context.Context, topic string, h Handler) error
	Close() error
}

type subscriber struct{}

// NewSubscriber creates a Kafka subscriber
func NewSubscriber(cfg SubscriberConfig) (Subscriber, error) {
	return &subscriber{}, nil
}

// SubscriberConfig config for Kafka consumer
type SubscriberConfig struct {
	Brokers []string
	GroupID string
}

// Subscribe starts consuming from topic
func (s *subscriber) Subscribe(ctx context.Context, topic string, h Handler) error {
	return nil
}

// Close closes the consumer
func (s *subscriber) Close() error {
	return nil
}
