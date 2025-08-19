package jwthndlr

import (
	"crypto/rand"
	"errors"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
)

var secret = []byte(rand.Text())

type Claims struct {
	jwt.RegisteredClaims
}

func NewJWT() (string, error) {
	now := time.Now()
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func VerifyJWT(claim string) error {
	token, err := jwt.ParseWithClaims(
		claim, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				slog.Error(
					"Unexpected signing method",
					slog.Any("signing_method", token.Header["alg"]),
				)
				return nil, ErrUnauthorized
			}
			return secret, nil
		})

	if err == nil {
		if _, ok := token.Claims.(*Claims); ok && token.Valid {
			return nil
		}
	}

	return ErrUnauthorized
}
