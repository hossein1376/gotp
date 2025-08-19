package usersrvc

import (
	"github.com/hossein1376/gotp/pkg/domain"
)

type UserService struct {
	repo *domain.Repository
}

func NewUserService(repo *domain.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}
