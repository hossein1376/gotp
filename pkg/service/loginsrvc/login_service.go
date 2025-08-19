package loginsrvc

import (
	"github.com/hossein1376/gotp/pkg/domain"
)

type LoginService struct {
	repo *domain.Repository
}

func NewLoginService(repo *domain.Repository) *LoginService {
	return &LoginService{
		repo: repo,
	}
}
