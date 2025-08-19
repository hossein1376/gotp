package usersrvc

import (
	"context"
	"errors"
	"fmt"

	"github.com/hossein1376/gotp/pkg/domain/model"
	"github.com/hossein1376/gotp/pkg/infrastructure/database/usersrp"
	"github.com/hossein1376/gotp/pkg/tools/errs"
)

var (
	ErrMissingUser = errors.New("user not found")
)

func (s *UserService) GetByPhone(
	ctx context.Context, phone string,
) (*model.User, error) {
	user, err := s.repo.UserRepo.FindByPhone(ctx, phone)
	if err != nil {
		switch {
		case errors.Is(err, usersrp.ErrUserNotFound):
			return nil, errs.NotFound(ErrMissingUser)
		default:
			return nil, fmt.Errorf("finding user by phone: %w", err)
		}
	}

	return user, nil
}
