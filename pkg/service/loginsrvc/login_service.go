package loginsrvc

import (
	"github.com/hossein1376/gotp/pkg/domain"
	"github.com/hossein1376/gotp/pkg/domain/model"
)

type LoginService struct {
	setKey string
	db     domain.Database
}

func NewLoginService(db domain.Database) *LoginService {
	return &LoginService{
		setKey: model.UsersSetKey,
		db:     db,
	}
}
