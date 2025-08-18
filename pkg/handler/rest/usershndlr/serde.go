package usershndlr

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/hossein1376/gotp/pkg/domain/model"
)

type GetUserByPhoneResponse struct {
	User *model.User `json:"user"`
}

type ListUsersResponse struct {
	Users []model.User `json:"users"`
}

type ListUsersRequest struct {
	count  int
	page   int
	offset int
	desc   bool
}

func bindListUsersRequest(r *http.Request) (*ListUsersRequest, error) {
	query := r.URL.Query()
	count, err := strconv.Atoi(query.Get("count"))
	if err != nil {
		return nil, fmt.Errorf("invalid count query: %w", err)
	}
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		return nil, fmt.Errorf("invalid page query: %w", err)
	}
	offset, err := strconv.Atoi(query.Get("offset"))
	if err != nil {
		return nil, fmt.Errorf("invalid offset query: %w", err)
	}
	desc, err := strconv.ParseBool(query.Get("desc"))
	if err != nil {
		return nil, fmt.Errorf("invalid desc query: %w", err)
	}

	return &ListUsersRequest{count, page, offset, desc}, nil
}
