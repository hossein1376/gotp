package loginsrvc

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/hossein1376/gotp/pkg/domain/model"
)

func (s *LoginService) SendLoginOTP(
	ctx context.Context, phone string,
) error {
	code, err := generateRandomCode()
	if err != nil {
		return fmt.Errorf("generate random code: %w", err)
	}

	otp := &model.OTP{Code: code, CreatedAt: time.Now().Unix()}
	data, err := json.Marshal(otp)
	if err != nil {
		return fmt.Errorf("marshal otp object: %w", err)
	}
	err = s.db.Set(ctx, phone, data, 2*time.Minute)
	if err != nil {
		return fmt.Errorf("insert otp object: %w", err)
	}

	// The code should be sent to the user here
	fmt.Println("OTP code:", code)

	return nil
}

func generateRandomCode() (string, error) {
	// Generate random number within the range 0 to (rangeValue - 1).
	r, err := rand.Int(rand.Reader, big.NewInt(int64(999_999-100_000)+1))
	if err != nil {
		return "", err
	}

	// AddSorted the minimum value to get a random number between 100,000 and 999,999
	return strconv.FormatInt(r.Int64()+int64(100_000), 10), nil
}
