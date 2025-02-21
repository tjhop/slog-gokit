package sloggokit_test

import (
	"bytes"
	"fmt"
	"log/slog"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"testing/slogtest"
	"time"

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

			// wrap the buffer in a new reader as pre-go 1.24 the
			// slogtest package calls results() multiple times, which
			// causes the tests to pass _without running checks_.
			// See https://github.com/golang/go/issues/67605
			dec := logfmt.NewDecoder(strings.NewReader(buf.String()))
			for dec.ScanRecord() {
				m := map[string]any{}
				for dec.ScanKeyval() {
					k, v, err := parseValue(string(dec.Key()), dec.Value())
					require.NoError(t, err)
					// If it's a map, merge it with the current map
					if m[k] != nil && reflect.TypeOf(m[k]).Kind() == reflect.Map {
						m[k] = mergeMaps(m[k].(map[string]any), v.(map[string]any))
						continue
					}
					m[k] = v
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

			dec := logfmt.NewDecoder(strings.NewReader(buf.String()))
			for dec.ScanRecord() {
				m := map[string]any{}
				for dec.ScanKeyval() {
					k, v, err := parseValue(string(dec.Key()), dec.Value())
					require.NoError(t, err)
					// If it's a map, merge it with the current map
					if m[k] != nil && reflect.TypeOf(m[k]).Kind() == reflect.Map {
						m[k] = mergeMaps(m[k].(map[string]any), v.(map[string]any))
						continue
					}
					m[k] = v
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

func mergeMaps(m1, m2 map[string]any) map[string]any {
	for k, v := range m2 {
		if m1[k] != nil && reflect.TypeOf(m1[k]).Kind() == reflect.Map {
			m1[k] = mergeMaps(m1[k].(map[string]any), v.(map[string]any))
			continue
		}
		m1[k] = v
	}
	return m1
}

func parseValue(key string, value []byte) (string, any, error) {
	switch key {
	case "level":
		var l slog.Level
		err := l.UnmarshalText([]byte(value))
		if err != nil {
			return key, nil, err
		}
		return key, l, nil
	case "time":
		// parse timestamp in iso8601 2025-02-20T16:58:30.683457-05:00
		parsedTime, err := time.Parse(time.RFC3339Nano, string(value))
		if err != nil {
			return key, nil, err
		}
		return key, parsedTime, nil
	}

	groups := strings.SplitN(key, ".", 2)
	if len(groups) != 2 {
		return key, string(value), nil
	}

	k, v, err := parseValue(groups[1], value)
	if err != nil {
		return key, nil, err
	}

	return groups[0], map[string]any{k: v}, nil
}
