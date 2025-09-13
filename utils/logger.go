package utils

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger() {
	logFile, err := os.OpenFile("inventory.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{}))
		slog.SetDefault(Logger)
		Logger.Error("Failed to create log file, logging to stderr", "error", err)
		return
	}

	Logger = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(Logger)

	Logger.Info("Inventory CLI started", "logFile", "inventory.log")
}
