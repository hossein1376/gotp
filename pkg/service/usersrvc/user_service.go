package usersrvc

import (
	"github.com/hossein1376/gotp/pkg/domain"
	"github.com/hossein1376/gotp/pkg/domain/model"
)

type UserService struct {
	setKey string
	db     domain.Database
}

func NewUserService(db domain.Database) *UserService {
	return &UserService{
		setKey: model.UsersSetKey,
		db:     db,
	}
}
