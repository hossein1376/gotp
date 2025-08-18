package redis

import (
	"github.com/hossein1376/gotp/pkg/domain"
)

var _ domain.Database = Client{}
