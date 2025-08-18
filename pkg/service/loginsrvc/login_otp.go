package loginsrvc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hossein1376/gotp/pkg/domain/model"
	infraRedis "github.com/hossein1376/gotp/pkg/infrastructure/database/redis"
	"github.com/hossein1376/gotp/pkg/tools/errs"
)

var (
	ErrNotFound = errors.New("not found")
)

func (s *LoginService) LoginOTP(
	ctx context.Context, phone, code string,
) error {
	data, err := s.db.Get(ctx, phone)
	if err != nil {
		switch {
		case errors.Is(err, infraRedis.ErrMissing):
			return errs.NotFound(ErrNotFound)
		default:
			return fmt.Errorf("get data: %w", err)
		}
	}
	otp := model.OTP{}
	err = json.Unmarshal(data, &otp)
	if err != nil {
		return fmt.Errorf("unmarshal otp object: %w", err)
	}

	if otp.Code != code {
		return errs.NotFound(ErrNotFound)
	}

	if err = s.db.AddSorted(ctx, s.setKey, otp.CreatedAt, phone); err != nil {
		return fmt.Errorf("add user: %w", err)
	}

	return nil
}
