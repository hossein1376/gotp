package service

import (
	"github.com/hossein1376/gotp/pkg/domain"
	"github.com/hossein1376/gotp/pkg/service/loginsrvc"
	"github.com/hossein1376/gotp/pkg/service/usersrvc"
)

type Services struct {
	LoginService *loginsrvc.LoginService
	UserService  *usersrvc.UserService
}

func NewServices(
	repo *domain.Repository,
) *Services {
	return &Services{
		LoginService: loginsrvc.NewLoginService(repo),
		UserService:  usersrvc.NewUserService(repo),
	}
}
