package rest

import (
	"context"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/hossein1376/gotp/pkg/handler/rest/jwthndlr"
	"github.com/hossein1376/gotp/pkg/handler/rest/serde"
	"github.com/hossein1376/gotp/pkg/tools/reqid"
	"github.com/hossein1376/gotp/pkg/tools/slogger"
)

func jwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			unauthorizedResp(w)
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			unauthorizedResp(w)
			return
		}
		if err := jwthndlr.VerifyJWT(parts[1]); err != nil {
			unauthorizedResp(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := reqid.NewRequestID()
		ctx := context.WithValue(r.Context(), reqid.RequestIDKey, id)
		ctx = slogger.WithAttrs(ctx, slog.Any("request_id", id))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if msg := recover(); msg != nil {
				slogger.Error(r.Context(), "recovered panic", slog.Any("msg", msg))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		path := r.URL.Path
		raw := r.URL.RawQuery
		defer func() {
			if raw != "" {
				path = path + "?" + raw
			}
			slogger.Info(r.Context(), "http server",
				slog.Group(
					"request",
					slog.String("method", r.Method),
					slog.String("request_path", path),
				),
				slog.Group(
					"response",
					slog.String("time_took", time.Since(start).String()),
				),
			)
		}()

		next.ServeHTTP(w, r)
	})
}

func unauthorizedResp(w http.ResponseWriter) {
	resp := serde.Response{Message: "unauthorized"}
	serde.WriteJson(w, http.StatusUnauthorized, resp, nil)
}

func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return ""
	}
	return parts[1]
}
