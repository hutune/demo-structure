package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Config for Redis connection
type Config struct {
	Addr     string
	Password string
	DB       int
}

// Client wraps go-redis Client
type Client struct {
	*redis.Client
}

// New creates a Redis client
func New(cfg Config) (*Client, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &Client{Client: cli}, nil
}

// Close closes the Redis connection
func (c *Client) Close() error {
	return c.Client.Close()
}
