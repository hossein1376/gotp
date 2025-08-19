package model

import (
	"fmt"
	"regexp"
	"time"
)

const (
	UserIndexKey  = "idx_users"
	UserKeyPrefix = "user:"
)

var PhoneRegEx = regexp.MustCompile(`^09\d{9}$`)

type User struct {
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	LastLogin time.Time `json:"last_login"`
}

type UserField int

const (
	UsrFldCreatedAt UserField = iota
	UsrFldPhone
	UsrFldLastLogin
)

func (f UserField) String() string {
	switch f {
	case UsrFldPhone:
		return "phone"
	case UsrFldCreatedAt:
		return "created_at"
	case UsrFldLastLogin:
		return "last_login"
	default:
		panic(fmt.Errorf("invalid user field %d", f))
	}
}

func ParseUserField(s string) (UserField, error) {
	switch s {
	case "phone":
		return UsrFldPhone, nil
	case "created_at":
		return UsrFldCreatedAt, nil
	case "last_login":
		return UsrFldLastLogin, nil
	default:
		return 0, fmt.Errorf("invalid user field: %q", s)
	}
}

type ListOptions[T ~int] struct {
	SortBy T
	Desc   bool
	Page   int
	Count  int
}
