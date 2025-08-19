package database

import (
	"context"
	"fmt"

	"github.com/hossein1376/gotp/pkg/domain"
	"github.com/hossein1376/gotp/pkg/infrastructure/database/cache"
	"github.com/hossein1376/gotp/pkg/infrastructure/database/loginrp"
	"github.com/hossein1376/gotp/pkg/infrastructure/database/usersrp"
)

func NewRepo(ctx context.Context, db *cache.DB) (*domain.Repository, error) {
	lginrp, err := loginrp.NewLoginRepo(db)
	if err != nil {
		return nil, fmt.Errorf("new login repo: %w", err)
	}
	usrrp, err := usersrp.NewUserRepo(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("new users repo: %w", err)
	}
	return &domain.Repository{LoginRepo: lginrp, UserRepo: usrrp}, nil
}
