package domain

import (
	"context"
	"time"

	"github.com/hossein1376/gotp/pkg/domain/model"
)

type Repository struct {
	LoginRepo LoginRepository
	UserRepo  UserRepository
}

type UserRepository interface {
	InsertIfNotExists(ctx context.Context, key string, user model.User) error
	FindByPhone(ctx context.Context, phone string) (*model.User, error)
	ListUsers(
		ctx context.Context, opts model.ListOptions[model.UserField],
	) ([]*model.User, error)
}

type LoginRepository interface {
	SetOTP(
		ctx context.Context, key string, data []byte, expiration time.Duration,
	) error
	GetOTP(ctx context.Context, key string) ([]byte, error)
	IsRateLimited(
		ctx context.Context, phone string, credit, cost int, window time.Duration,
	) (bool, error)
}
