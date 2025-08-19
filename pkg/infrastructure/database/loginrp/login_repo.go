package loginrp

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/hossein1376/gotp/pkg/domain"
	"github.com/hossein1376/gotp/pkg/infrastructure/database/cache"
)

var (
	ErrNotFound = errors.New("not found")
)

type LoginRepo struct {
	client *redis.Client
}

var _ domain.LoginRepository = (*LoginRepo)(nil)

func NewLoginRepo(db *cache.DB) *LoginRepo {
	return &LoginRepo{client: db.Client()}
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
