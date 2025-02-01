package logger

import (
	"log/slog"
	"os"
)

const (
	dev   = "dev"
	local = "local"
	prod  = "prod"
)

func SetupLogger(env string) *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == "file" {
				return slog.Attr{}
			}
			return a
		},
	}

	switch env {
	case dev:
		opts.Level = slog.LevelDebug
		handler := slog.NewJSONHandler(os.Stdout, opts)
		return slog.New(handler)
	case local:
		opts.Level = slog.LevelDebug
		handler := slog.NewTextHandler(os.Stdout, opts)
		return slog.New(handler)
	case prod:
		opts.Level = slog.LevelError
		handler := slog.NewJSONHandler(os.Stdout, opts)
		return slog.New(handler)
	default:
		handler := slog.NewJSONHandler(os.Stdout, opts)
		return slog.New(handler)
	}
}
