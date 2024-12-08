package log

import (
	"log/slog"
	"os"
)

func LogInit() *slog.Logger {

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return log
}
