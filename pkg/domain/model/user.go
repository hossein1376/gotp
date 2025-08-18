package model

import (
	"regexp"
	"time"
)

const UsersSetKey = "users"

var PhoneRegEx = regexp.MustCompile(`^09\d{9}$`)

type User struct {
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

type SortedSet struct {
	Key  string
	Data any
}
