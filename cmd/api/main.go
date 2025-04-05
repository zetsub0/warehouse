package main

import (
	"context"
	"log/slog"

	"warehouse/internal/app/http"
	"warehouse/internal/config"
	"warehouse/internal/logger"
)

func main() {
	ctx := context.Background()

	cfg := config.ParseConfig()

	log := logger.SetupLogger(cfg.Env)
	slog.SetDefault(log)

	srv := http.New(cfg.HTTPServer, nil)

	srv.Run(ctx)
}
