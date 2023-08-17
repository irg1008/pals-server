package log

import (
	"os"
	"time"

	"log/slog"

	"github.com/lmittmann/tint"
)

func newLogger(debug bool) *slog.Logger {
	w := os.Stderr
	var handler slog.Handler

	if debug {
		handler = tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		})
	} else {
		handler = slog.NewJSONHandler(w, nil)
	}

	return slog.New(handler)
}

func SetDefaultLogger(debug bool) *slog.Logger {
	logger := newLogger(debug)
	slog.SetDefault(logger)
	return logger
}
