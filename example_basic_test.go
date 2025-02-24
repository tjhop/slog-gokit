package sloggokit_test

import (
	"log/slog"
	"os"

	"github.com/go-kit/log"
	slgk "github.com/tjhop/slog-gokit"
)

func Example_basic() {
	// Take an existing go-kit/log Logger:
	gklogger := log.NewLogfmtLogger(os.Stderr)

	// Create an slog Logger that chains log calls to the go-kit/log Logger:
	slogger := slog.New(slgk.NewGoKitHandler(gklogger, nil))
	slogger.WithGroup("example_group").With("foo", "bar").Info("hello world")
}
