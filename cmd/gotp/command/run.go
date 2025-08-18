package command

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/hossein1376/gotp/pkg/handler/rest"
	"github.com/hossein1376/gotp/pkg/infrastructure/database/redis"
	"github.com/hossein1376/gotp/pkg/service"
)

func Run() error {
	ctx := context.Background()

	db, err := redis.New(ctx, "127.0.0.1:6379")
	if err != nil {
		return fmt.Errorf("redis.New: %w", err)
	}
	defer db.Close()

	services := service.NewServices(db)

	srv := rest.NewServer("127.0.0.1:9009", services)

	errCh := make(chan error)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	go func() {
		errCh <- srv.ListenAndServe()
	}()

	select {
	case err = <-errCh:
		return fmt.Errorf("srv.ListenAndServe: %w", err)
	case <-signalCh:
		return srv.Shutdown(ctx)
	}
}
