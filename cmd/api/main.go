package main

import (
	"context"

	"warehouse/internal/app/http"
	"warehouse/internal/config"
)

func main() {
	ctx := context.Background()

	cfg := config.ParseConfig()
	srv := http.New(cfg.HTTPServer, nil)

	srv.Run(ctx)
}
