package service

import (
	"time"

	"github.com/hossein1376/gotp/pkg/domain"
	"github.com/hossein1376/gotp/pkg/service/loginsrvc"
	"github.com/hossein1376/gotp/pkg/service/usersrvc"
)

type Services struct {
	LoginService *loginsrvc.LoginService
	UserService  *usersrvc.UserService
}

func NewServices(
	repo *domain.Repository, // cfg config.RateLimit
) *Services {
	return &Services{
		// Ideally, these vales should come from a config file
		LoginService: loginsrvc.NewLoginService(repo, 3, 1, 10*time.Second),
		UserService:  usersrvc.NewUserService(repo),
	}
}
