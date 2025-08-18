package domain

import (
	"context"
	"time"

	"github.com/hossein1376/gotp/pkg/domain/model"
)

type Database interface {
	Set(
		ctx context.Context, key string, data []byte, expiration time.Duration,
	) error
	Get(ctx context.Context, key string) ([]byte, error)
	AddSorted(ctx context.Context, key string, score int64, data any) error
	GetByValueSorted(
		ctx context.Context, key, value string,
	) (float64, error)
	ScoreRangeSorted(
		ctx context.Context, key string, start, stop string,
	) ([]model.SortedSet, error)
}
