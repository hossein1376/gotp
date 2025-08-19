package usersrp

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"

	"github.com/hossein1376/gotp/pkg/domain"
	"github.com/hossein1376/gotp/pkg/domain/model"
	"github.com/hossein1376/gotp/pkg/infrastructure/database/cache"
	"github.com/hossein1376/gotp/pkg/infrastructure/database/invconv"
)

var (
	ErrUserNotFound = errors.New("user was not found")
)

type UserRepo struct {
	client     *redis.Client
	insertUser *redis.Script
}

var _ domain.UserRepository = (*UserRepo)(nil)

func NewUserRepo(ctx context.Context, db *cache.DB) (*UserRepo, error) {
	client := db.Client()
	if err := createUsersSchema(ctx, client); err != nil {
		return nil, fmt.Errorf("creating users schema: %w", err)
	}
	insertUserScript, err := os.ReadFile("assets/scripts/insert_user.lua")
	if err != nil {
		return nil, fmt.Errorf("reading insert_user.lua: %w", err)
	}
	insertUser := redis.NewScript(string(insertUserScript))

	return &UserRepo{client: client, insertUser: insertUser}, nil
}

func (r *UserRepo) InsertIfNotExists(
	ctx context.Context, key string, user model.User,
) error {
	err := r.insertUser.Run(
		ctx,
		r.client,
		[]string{key},
		user.Phone,
		user.CreatedAt.Unix(),
		user.LastLogin.Unix(),
	).Err()
	if errors.Is(err, redis.Nil) {
		return nil
	}
	return err
}

func (r *UserRepo) FindByPhone(
	ctx context.Context, phone string,
) (*model.User, error) {
	query := fmt.Sprintf("@phone:%s", phone)
	results, err := r.client.FTSearch(ctx, model.UserIndexKey, query).Result()
	switch {
	case err != nil:
		return nil, fmt.Errorf("searching users: %w", err)
	case results.Total == 0:
		return nil, ErrUserNotFound
	case results.Total > 1:
		return nil, fmt.Errorf("expected 1 user, found %d", results.Total)
	}

	result := results.Docs[0]
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("getting user: %w", err)
	}
	user, err := invconv.UserInverter(result.Fields)
	if err != nil {
		return nil, fmt.Errorf("mapping user fields: %w", err)
	}

	return user, nil
}

func (r *UserRepo) ListUsers(
	ctx context.Context, opts model.ListOptions[model.UserField],
) ([]*model.User, error) {
	query := &redis.FTSearchOptions{
		SortBy: []redis.FTSearchSortBy{{
			FieldName: opts.SortBy.String(), Asc: !opts.Desc, Desc: opts.Desc,
		}},
		SortByWithCount: false,
		LimitOffset:     opts.Page * opts.Count,
		Limit:           opts.Count,
	}
	results, err := r.client.FTSearchWithArgs(
		ctx, model.UserIndexKey, "*", query,
	).Result()
	if err != nil {
		return nil, fmt.Errorf("searching users: %w", err)
	}
	users := make([]*model.User, results.Total)
	for i, result := range results.Docs {
		user, err := invconv.UserInverter(result.Fields)
		if err != nil {
			return nil, fmt.Errorf("mapping user fields: %w", err)
		}
		users[i] = user
	}

	return users, nil
}

func createUsersSchema(ctx context.Context, rdb *redis.Client) error {
	err := rdb.FTInfo(ctx, model.UserIndexKey).Err()
	switch {
	case err == nil:
		return nil
	case err.Error() == "Unknown index name":
		// continue
	default:
		return fmt.Errorf("getting index info: %w", err)
	}

	err = rdb.FTCreate(
		ctx,
		model.UserIndexKey,
		&redis.FTCreateOptions{
			OnHash: true, Prefix: []any{1, model.UserKeyPrefix},
		},
		&redis.FieldSchema{
			FieldName: "phone",
			FieldType: redis.SearchFieldTypeText,
			Sortable:  true,
		},
		&redis.FieldSchema{
			FieldName: "created_at",
			FieldType: redis.SearchFieldTypeNumeric,
			Sortable:  true,
		},
		&redis.FieldSchema{
			FieldName: "Last_login",
			FieldType: redis.SearchFieldTypeNumeric,
			Sortable:  true,
		},
	).Err()
	if err != nil {
		return fmt.Errorf("creating users index: %w", err)
	}

	return nil
}
