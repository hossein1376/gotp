package invconv

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hossein1376/gotp/pkg/domain/model"
)

func UserInverter(fields map[string]string) (*model.User, error) {
	phone := fields[model.UsrFldPhone.String()]
	createdAt, err := strconv.ParseInt(
		fields[model.UsrFldCreatedAt.String()], 10, 64,
	)
	if err != nil {
		return nil, fmt.Errorf("pasring created at: %w", err)
	}
	lastLogin, err := strconv.ParseInt(
		fields[model.UsrFldLastLogin.String()], 10, 64,
	)
	if err != nil {
		return nil, fmt.Errorf("pasring last login: %w", err)
	}

	return &model.User{
		Phone:     phone,
		CreatedAt: time.Unix(createdAt, 0),
		LastLogin: time.Unix(lastLogin, 0),
	}, nil
}
