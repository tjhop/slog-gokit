package sloggokit_test

import (
	"log/slog"
	"os"

	"github.com/go-kit/log"
	slgk "github.com/tjhop/slog-gokit"
)

func Example_dynamicLevel() {
	// Take an existing go-kit/log Logger:
	gklogger := log.NewLogfmtLogger(os.Stderr)

	// Create an slog Logger that chains log calls to the go-kit/log Logger.
	//
	// To dynamically change the logger's level, create an `slog.LevelVar`
	// and use it to set the logger's level as needed:
	lvl := &slog.LevelVar{} // Info level by default
	slogger := slog.New(slgk.NewGoKitHandler(gklogger, lvl))
	slogger.WithGroup("example_group").With("foo", "bar").Info("hello world")

	// Change level to debug, etc:
	lvl.Set(slog.LevelDebug)
	slogger.WithGroup("example_group").With("foo", "bar").Debug("helpful debug info")
}
