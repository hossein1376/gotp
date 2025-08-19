package loginrp

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/hossein1376/gotp/pkg/domain"
	"github.com/hossein1376/gotp/pkg/infrastructure/database/cache"
)

var (
	ErrNotFound = errors.New("not found")
)

type LoginRepo struct {
	client    *redis.Client
	rateLimit *redis.Script
}

var _ domain.LoginRepository = (*LoginRepo)(nil)

func NewLoginRepo(db *cache.DB) (*LoginRepo, error) {
	insertUserScript, err := os.ReadFile("assets/scripts/rate_limit.lua")
	if err != nil {
		return nil, fmt.Errorf("reading insert_user.lua: %w", err)
	}
	rateLimit := redis.NewScript(string(insertUserScript))

	return &LoginRepo{client: db.Client(), rateLimit: rateLimit}, nil
}

func (r *LoginRepo) SetOTP(
	ctx context.Context, key string, data []byte, expiration time.Duration,
) error {
	return r.client.Set(ctx, key, data, expiration).Err()
}

func (r *LoginRepo) GetOTP(ctx context.Context, key string) ([]byte, error) {
	cmd := r.client.Get(ctx, key)
	if errors.Is(cmd.Err(), redis.Nil) {
		return nil, ErrNotFound
	}
	b, err := cmd.Bytes()
	if err != nil {
		return nil, fmt.Errorf("get otp: %w", err)
	}

	return b, nil
}

func (r *LoginRepo) IsRateLimited(
	ctx context.Context, phone string, credit, cost int, window time.Duration,
) (bool, error) {
	key := "ratelimit:" + phone
	isAllowed, err := r.rateLimit.Run(
		ctx, r.client, []string{key}, credit, cost, window.Milliseconds(),
	).Int()
	if err != nil {
		return false, err
	}

	return isAllowed == 1, nil
}
