package command

import (
	"context"
	"flag"
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

	var cfg config
	flag.StringVar(&cfg.redisAddr, "cache", "127.0.0.1:6379", "Redis address")
	flag.StringVar(&cfg.listenAddr, "listen", "127.0.0.1:9009", "Listen address")
	flag.Parse()

	logger := slogger.NewJSONLogger(slog.LevelDebug, os.Stdout)
	slog.SetDefault(logger)

	db, err := cache.New(ctx, cfg.redisAddr)
	if err != nil {
		return fmt.Errorf("cache.New: %w", err)
	}
	defer db.Close()
	slog.Debug("cache started")

	repo, err := database.NewRepo(ctx, db)
	services := service.NewServices(repo)
	srv := rest.NewServer(cfg.listenAddr, services)
	slog.Debug("initialized repositories, services and handlers")

	errCh := make(chan error)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	go func() {
		slog.Info("starting server", slog.String("addr", cfg.listenAddr))
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
