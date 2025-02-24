package sloggokit_test

import (
	"log/slog"
	"os"

	"github.com/go-kit/log"
	slgk "github.com/tjhop/slog-gokit"
)

func Example_constantLevel() {
	// Take an existing go-kit/log Logger:
	gklogger := log.NewLogfmtLogger(os.Stderr)

	// Create an slog Logger that chains log calls to the go-kit/log Logger.
	//
	// The level for the slog Logger can be set explicitly with anything
	// that satisfies the `slog.Leveler` interface, such as slog's level
	// constants:
	slogger := slog.New(slgk.NewGoKitHandler(gklogger, slog.LevelDebug))
	slogger.WithGroup("example_group").With("foo", "bar").Info("hello world")
}
