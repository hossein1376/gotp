package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type DB struct {
	client *redis.Client
}

func New(ctx context.Context, addr string) (*DB, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Protocol: 2,
		// We might want to get these values from config as well
		Password: "",
		DB:       0,
	})

	resp := rdb.Ping(ctx)
	if err := resp.Err(); err != nil {
		return nil, err
	}
	return &DB{client: rdb}, nil
}

func (db DB) Client() *redis.Client {
	return db.client
}

func (db DB) Close() error {
	return db.client.Close()
}
