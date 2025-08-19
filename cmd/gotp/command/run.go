package command

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/hossein1376/gotp/pkg/handler/rest"
	"github.com/hossein1376/gotp/pkg/infrastructure/database"
	"github.com/hossein1376/gotp/pkg/infrastructure/database/cache"
	"github.com/hossein1376/gotp/pkg/service"
	"github.com/hossein1376/gotp/pkg/tools/slogger"
)

type config struct {
	redisAddr  string
	listenAddr string
}

func Run() error {
	ctx := context.Background()

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "0.0.0.0:6379"
	}
	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = "0.0.0.0:9009"
	}

	logger := slogger.NewJSONLogger(slog.LevelDebug, os.Stdout)
	slog.SetDefault(logger)

	db, err := cache.New(ctx, redisAddr)
	if err != nil {
		return fmt.Errorf("cache.New: %w", err)
	}
	defer db.Close()
	slog.Debug("database started")

	repo, err := database.NewRepo(ctx, db)
	if err != nil {
		return fmt.Errorf("database.New: %w", err)
	}
	services := service.NewServices(repo)
	srv := rest.NewServer(listenAddr, services)
	slog.Debug("initialized repositories, services and handlers")

	errCh := make(chan error)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	go func() {
		slog.Info("starting server", slog.String("addr", listenAddr))
		errCh <- srv.ListenAndServe()
	}()

	select {
	case err = <-errCh:
		return fmt.Errorf("srv.ListenAndServe: %w", err)
	case <-signalCh:
		slog.Info("shutdown signal received")
		return srv.Shutdown(ctx)
	}
}
