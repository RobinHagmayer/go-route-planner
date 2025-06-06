package main

import (
	"fmt"
	"log/slog"

	"github.com/RobinHagmayer/go-route-planner/internal/logger"
)

func main() {
	logger.Init()
	defer logger.Close()
	slog.Info("App started")
	fmt.Println("Hello go-route-planner")
}
