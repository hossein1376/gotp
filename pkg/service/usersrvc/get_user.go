package usersrvc

import (
	"context"
	"errors"
	"time"

	"github.com/hossein1376/gotp/pkg/domain/model"
	"github.com/hossein1376/gotp/pkg/tools/errs"
)

var (
	ErrMissingUser = errors.New("user not found")
)

func (s *UserService) GetByPhone(
	ctx context.Context, phone string,
) (*model.User, error) {
	createdAt, err := s.db.GetByValueSorted(ctx, s.setKey, phone)
	if err != nil {
		return nil, errs.NotFound(ErrMissingUser)
	}

	return &model.User{
		Phone:     phone,
		CreatedAt: time.Unix(int64(createdAt), 0),
	}, nil
}
