package logger

import (
	"fmt"
	"log/slog"
	"os"
)

var (
	Log  *slog.Logger
	file *os.File
)

func Init() {
	var err error
	file, err = os.OpenFile("go-route-planner.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Printf("Error while opening the log file: %v", err)
		os.Exit(1)
	}

	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	Log = slog.New(handler)
	slog.SetDefault(Log)
}

func Close() {
	if file == nil {
		fmt.Println("Error: can't close a not existing logging file")
		os.Exit(1)
	}
	file.Close()
}
