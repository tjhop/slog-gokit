package sloggokit_test

import (
	"bytes"
	"fmt"
	"log/slog"
	"regexp"
	"strings"
	"testing"
	"testing/slogtest"

	"github.com/go-kit/log"
	"github.com/go-logfmt/logfmt"
	"github.com/stretchr/testify/require"

	slgk "github.com/tjhop/slog-gokit"
)

var (
	logRegexp = regexp.MustCompile(`level=(?P<LevelValue>warn|info|error|debug).*time=.+msg=.+`)
)

func TestNewGoKitHandler(t *testing.T) {
	t.Run("nil level", func(t *testing.T) {
		var buf bytes.Buffer
		h := slgk.NewGoKitHandler(log.NewLogfmtLogger(&buf), nil)

		results := func() []map[string]any {
			var ms []map[string]any

			// Print logs for humans.
			fmt.Println(buf.String())

			dec := logfmt.NewDecoder(&buf)
			for dec.ScanRecord() {
				m := make(map[string]any)

				for dec.ScanKeyval() {
					m[string(dec.Key())] = dec.Value()
				}
				ms = append(ms, m)
			}
			err := dec.Err()
			require.NoError(t, err, "failed to decode logfmt entry")

			return ms
		}

		err := slogtest.TestHandler(h, results)
		require.NoError(t, err, "failed slog handler test suite")
	})
	t.Run("debug level", func(t *testing.T) {
		var buf bytes.Buffer
		lvl := &slog.LevelVar{}
		lvl.Set(slog.LevelDebug)
		h := slgk.NewGoKitHandler(log.NewLogfmtLogger(&buf), lvl)

		results := func() []map[string]any {
			var ms []map[string]any

			// Print logs for humans.
			fmt.Println(buf.String())

			dec := logfmt.NewDecoder(&buf)
			for dec.ScanRecord() {
				m := make(map[string]any)

				for dec.ScanKeyval() {
					m[string(dec.Key())] = dec.Value()
				}
				ms = append(ms, m)
			}
			err := dec.Err()
			require.NoError(t, err, "failed to decode logfmt entry")

			return ms
		}

		err := slogtest.TestHandler(h, results)
		require.NoError(t, err, "failed slog handler test suite")
	})
	t.Run("dynamic level", func(t *testing.T) {
		var buf bytes.Buffer
		lvl := &slog.LevelVar{}

		gklogger := log.NewLogfmtLogger(&buf)
		h := slgk.NewGoKitHandler(gklogger, lvl)
		slogger := slog.New(h)

		wantedLevelCounts := map[string]int{"info": 1, "debug": 1}

		// Start at debug level.
		lvl.Set(slog.LevelDebug)
		slogger.Info("info", "hello", "world")
		slogger.Debug("debug", "hello", "world")

		// We expect to see one of each log level type in `wantedLevelCounts`
		counts := getLogEntryLevelCounts(buf.String(), logRegexp)
		require.Equal(t, wantedLevelCounts["info"], counts["info"], "info log successfully detected")
		require.Equal(t, wantedLevelCounts["debug"], counts["debug"], "debug log successfully detected")

		// Print logs for humans.
		fmt.Println(buf.String())
		buf.Reset()

		// Test that log level can be adjusted on-the-fly to info and
		// that a subsequent call to write a debug level log is _not_
		// written to the file.
		lvl.Set(slog.LevelInfo)

		slogger.Info("info", "hello", "world")
		slogger.Debug("debug", "hello", "world")

		// We expect to see one info log, and 0 debug logs.
		counts = getLogEntryLevelCounts(buf.String(), logRegexp)
		require.Equal(t, wantedLevelCounts["info"], counts["info"], "info log successfully detected")
		require.NotEqual(t, wantedLevelCounts["debug"], counts["debug"], "extra debug log detected")

		// Print logs for humans to see, if needed.
		fmt.Println(buf.String())
		buf.Reset()
	})
}

func getLogEntryLevelCounts(s string, re *regexp.Regexp) map[string]int {
	counters := make(map[string]int)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			levelIndex := re.SubexpIndex("LevelValue")

			counters[strings.ToLower(matches[levelIndex])]++
		}
	}

	return counters
}
