package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	db *redis.Client
}

func New(ctx context.Context, addr string) (*Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		// We might want to get these values from config as well
		Password: "",
		DB:       0,
	})

	resp := rdb.Ping(ctx)
	if err := resp.Err(); err != nil {
		return nil, err
	}
	return &Client{db: rdb}, nil
}

func (c *Client) Close() error {
	return c.db.Close()
}
