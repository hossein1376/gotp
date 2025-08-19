package usersrvc

import (
	"context"

	"github.com/hossein1376/gotp/pkg/domain/model"
)

func (s *UserService) ListUsers(
	ctx context.Context, opts model.ListOptions[model.UserField],
) ([]*model.User, error) {
	if opts.Count == 0 {
		opts.Count = 20
	}
	return s.repo.UserRepo.ListUsers(ctx, opts)
}
