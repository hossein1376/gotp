package model

const (
	LoginOTPKeyPrefix = "login:otp:"
)

type LoginOTP struct {
	Phone     string `json:"phone"`
	Code      string `json:"code"`
	CreatedAt int64  `json:"created_at"`
}
