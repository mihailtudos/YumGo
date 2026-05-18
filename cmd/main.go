package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mihailtudos/yumgo/backend"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

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
