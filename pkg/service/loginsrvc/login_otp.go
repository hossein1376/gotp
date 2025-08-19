package loginsrvc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hossein1376/gotp/pkg/domain/model"
	"github.com/hossein1376/gotp/pkg/infrastructure/database/loginrp"
	"github.com/hossein1376/gotp/pkg/tools/errs"
)

var (
	ErrNotFound = errors.New("not found")
)

func (s *LoginService) LoginOTP(ctx context.Context, phone, code string) error {
	data, err := s.repo.LoginRepo.GetOTP(ctx, phone)
	if err != nil {
		switch {
		case errors.Is(err, loginrp.ErrNotFound):
			return errs.NotFound(ErrNotFound)
		default:
			return fmt.Errorf("get data: %w", err)
		}
	}
	otp := model.LoginOTP{}
	err = json.Unmarshal(data, &otp)
	if err != nil {
		return fmt.Errorf("unmarshal otp object: %w", err)
	}

	if otp.Code != code {
		return errs.NotFound(ErrNotFound)
	}

	if err = s.repo.UserRepo.InsertIfNotExists(
		ctx, model.UserKeyPrefix+phone, model.User{
			Phone:     phone,
			CreatedAt: time.Unix(otp.CreatedAt, 0),
			LastLogin: time.Now(),
		},
	); err != nil {
		return fmt.Errorf("add user: %w", err)
	}

	return nil
}
