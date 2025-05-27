package logger

import (
	"fmt"
	"log/slog"
	"os"
)

var Log *slog.Logger

func Init() {
	file, err := os.OpenFile("go-route-planner.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error while opening the log file: %w", err)
		os.Exit(1)
	}

	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	Log = slog.New(handler)
	slog.SetDefault(Log)
}
