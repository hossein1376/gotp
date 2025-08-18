package serde

import (
	"errors"
	"net/http"

	"github.com/hossein1376/gotp/pkg/tools/errs"
)

type Response struct {
	Message string `json:"message"`
}

func ExtractFromErr(err error) (int, Response) {
	if err == nil {
		panic("ExtractFromErr was called with nil error")
	}

	var e errs.Error
	if errors.As(err, &e) {
		return e.HTTPStatusCode, Response{e.Message}
	}
	return http.StatusInternalServerError, Response{
		http.StatusText(http.StatusInternalServerError),
	}
}
