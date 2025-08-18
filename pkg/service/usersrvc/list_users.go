package usersrvc

import (
	"context"

	"github.com/hossein1376/gotp/pkg/domain/model"
)

func (s *UserService) ListUsers(
	ctx context.Context, count int, page int, offset int, desc bool,
) ([]model.User, error) {

	return nil, nil
}
