package main

import (
	"context"
	"net/http"
	"time"

	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/unfadingDawn/payment-service/internal/config"
	"github.com/unfadingDawn/payment-service/internal/db"
	"github.com/unfadingDawn/payment-service/internal/handler"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	config, err := config.Load()
	// TODO: make normal err processing
	if err != nil {
		return
	}
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	err2 := handler.RegisterRoutes(r)
	if err2 != nil {
		return
	}
	ctx2, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// dsn := "postgres://paygo:paygo@localhost:5432/paygo?sslmode=disable"

	dsn := config.DatabaseURL

	pool, err := db.NewPool(ctx2, dsn)
	if err != nil {
		return
	}
	defer pool.Close()
	err3 := db.Ping(ctx2, pool)
	if err3 != nil {
		return
	}

	err1 := http.ListenAndServe(config.HTTPPort, r)
	if err1 != nil {
		return
	}
}
