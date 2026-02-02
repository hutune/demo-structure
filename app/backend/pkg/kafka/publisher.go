package kafka

import (
	"context"
)

// Publisher publishes messages to Kafka
type Publisher interface {
	Publish(ctx context.Context, topic string, key, value []byte) error
	Close() error
}

// PublisherConfig config for Kafka producer
type PublisherConfig struct {
	Brokers []string
}

type publisher struct {
	// producer *kafka.Producer
}

// NewPublisher creates a Kafka publisher
func NewPublisher(cfg PublisherConfig) (Publisher, error) {
	return &publisher{}, nil
}

// Publish sends a message to the topic
func (p *publisher) Publish(ctx context.Context, topic string, key, value []byte) error {
	return nil
}

// Close closes the producer
func (p *publisher) Close() error {
	return nil
}
