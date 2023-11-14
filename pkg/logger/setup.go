package logger

import (
	"log/slog"
	"os"
)

func SetupLogger(debug bool) {
	var levelVar slog.LevelVar
	if debug {
		levelVar.Set(slog.LevelDebug)
	} else {
		levelVar.Set(slog.LevelError)
	}

	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     &levelVar,
	}

	if debug {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &opts)))
	} else {
		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &opts)))
	}
}
