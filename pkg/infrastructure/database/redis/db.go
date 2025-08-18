package redis

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/hossein1376/gotp/pkg/domain/model"
)

var (
	ErrMissing = errors.New("not found")
)

func (c Client) Set(
	ctx context.Context, key string, data []byte, expiration time.Duration,
) error {
	cmd := c.db.Set(ctx, key, data, expiration)
	return cmd.Err()
}

func (c Client) Get(ctx context.Context, key string) ([]byte, error) {
	cmd := c.db.Get(ctx, key)
	b, err := cmd.Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, ErrMissing
	}
	return b, nil
}

func (c Client) AddSorted(
	ctx context.Context, key string, score int64, data any,
) error {
	// precision can be lost here, but it's good enough for now
	cmd := c.db.ZAdd(ctx, key, redis.Z{Score: float64(score), Member: data})
	return cmd.Err()
}

func (c Client) GetByValueSorted(
	ctx context.Context, key, value string,
) (float64, error) {
	score, err := c.db.ZScore(ctx, key, value).Result()
	if err != nil {
		switch {
		case errors.Is(err, redis.Nil):
			return 0, ErrMissing
		default:
			return 0, err
		}
	}
	return score, nil
}

func (c Client) ScoreRangeSorted(
	ctx context.Context, key string, start, stop string,
) ([]model.SortedSet, error) {
	sets, err := c.db.ZRangeByScoreWithScores(
		ctx, key, &redis.ZRangeBy{Min: start, Max: stop},
	).Result()
	if err != nil {
		return nil, err
	}
	return setsFromRedisZ(sets), nil
}

func setsFromRedisZ(items []redis.Z) []model.SortedSet {
	sets := make([]model.SortedSet, len(items))
	for i, item := range items {
		sets[i] = model.SortedSet{
			Key:  strconv.FormatFloat(item.Score, 'f', -1, 64),
			Data: item.Member,
		}
	}
	return sets
}
