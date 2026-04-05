package sloggokit

import (
	"log/slog"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func goKitLevelFunc(logger log.Logger, lvl slog.Level) log.Logger {
	switch {
	case lvl >= slog.LevelError:
		logger = level.Error(logger)
	case lvl >= slog.LevelWarn:
		logger = level.Warn(logger)
	case lvl >= slog.LevelInfo:
		logger = level.Info(logger)
	default:
		logger = level.Debug(logger)
	}

	return logger
}
