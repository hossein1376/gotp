package model

type OTP struct {
	Code      string `json:"code"`
	CreatedAt int64  `json:"created_at"`
}
