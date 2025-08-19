package usershndlr

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/hossein1376/gotp/pkg/domain/model"
	"github.com/hossein1376/gotp/pkg/handler/rest/serde"
)

type GetUserByPhoneResponse struct {
	User *model.User `json:"user"`
}

type ListUsersResponse struct {
	Users []*model.User `json:"users"`
}

type ListUsersRequest struct {
	count int
	page  int
	desc  bool
	sort  model.UserField
}

func bindListUsersRequest(r *http.Request) (*ListUsersRequest, error) {
	query := r.URL.Query()

	count, err := serde.ValueOrDefault(query.Get("count"), strconv.Atoi)
	if err != nil {
		return nil, fmt.Errorf("invalid count query: %w", err)
	}
	if count < 0 || count > 100 {
		return nil, fmt.Errorf("invalid count value: %d", count)
	}

	page, err := serde.ValueOrDefault(query.Get("page"), strconv.Atoi)
	if err != nil {
		return nil, fmt.Errorf("invalid page query: %w", err)
	}
	if page < 0 {
		return nil, fmt.Errorf("invalid page value: %d", page)
	}

	desc, err := serde.ValueOrDefault(query.Get("desc"), strconv.ParseBool)
	if err != nil {
		return nil, fmt.Errorf("invalid desc query: %w", err)
	}

	sort, err := serde.ValueOrDefault(query.Get("sort"), model.ParseUserField)
	if err != nil {
		return nil, fmt.Errorf("invalid sort query: %w", err)
	}

	return &ListUsersRequest{count, page, desc, sort}, nil
}
