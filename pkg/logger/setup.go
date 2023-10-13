package logger

import (
	"go-boilerplate/pkg/config"
	"log/slog"
	"os"
)

func SetupLogger() {
	var levelVar slog.LevelVar
	if config.App.Debug {
		levelVar.Set(slog.LevelDebug)
	} else {
		levelVar.Set(slog.LevelError)
	}

	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     &levelVar,
	}

	if config.App.Debug {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &opts)))
	} else {
		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &opts)))
	}
}
