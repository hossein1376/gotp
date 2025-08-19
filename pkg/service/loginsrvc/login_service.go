package loginsrvc

import (
	"errors"
	"time"

	"github.com/hossein1376/gotp/pkg/domain"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrTooManyRequests = errors.New("too many requests")
)

type LoginService struct {
	repo      *domain.Repository
	rateLimit rateLimit
}

type rateLimit struct {
	credit, cost int
	window       time.Duration
}

func NewLoginService(
	repo *domain.Repository, credit, cost int, window time.Duration,
) *LoginService {
	return &LoginService{
		repo: repo, rateLimit: rateLimit{credit, cost, window},
	}
}
