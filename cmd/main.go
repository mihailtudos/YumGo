package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mihailtudos/yumgo/backend"
	"github.com/mihailtudos/yumgo/common/log"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	log.Init(slog.LevelInfo)

	dsn := os.Getenv("POSTGRES_URL")
	if dsn == "" {
		panic("POSTGRES_URL is required")
	}

	dbPgx, err := pgxpool.New(ctx, dsn)
	if dsn == "" {
		panic(err)
	}

	backend.New(ctx, dbPgx)
}
