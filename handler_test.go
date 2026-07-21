package sloggokit_test

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"
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

// TestWithAttrsConcurrency specifically tests for race conditions when
// multiple goroutines call WithAttrs on the same handler concurrently.
// This reproduces the exact scenario from the bug report where multiple
// goroutines share a logger and call .With() simultaneously.
//
// The race occurs in the buggy code at: pairs = append(h.preformatted, pairs...)
// When h.preformatted has capacity to hold the additional elements, append
// writes to the shared underlying array, causing concurrent goroutines to
// write to the same memory location.
func TestWithAttrsConcurrency(t *testing.T) {
	handler := slgk.NewGoKitHandler(nil, nil)

	// Create a handler with attributes. The key is to create a scenario
	// where h.preformatted will have extra capacity.
	handler = handler.WithAttrs([]slog.Attr{
		slog.String("base", "value"),
	})

	// Call WithAttrs multiple times in quick succession to increase the
	// chance that the preformatted slice has extra capacity
	for i := 0; i < 5; i++ {
		handler = handler.WithAttrs([]slog.Attr{
			slog.String(fmt.Sprintf("key%d", i), fmt.Sprintf("val%d", i)),
		})
	}

	const numGoroutines = 50
	const iterations = 1000

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Simulate the exact pattern from the race detector report:
	// multiple goroutines calling .With() (which calls WithAttrs) concurrently
	// on the SAME handler instance
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				// This triggers the race in the buggy code:
				// append(h.preformatted, pairs...) causes multiple goroutines
				// to write to the same underlying array
				_ = handler.WithAttrs([]slog.Attr{
					slog.Int("g", id),
					slog.Int("i", j),
				})
			}
		}(i)
	}

	wg.Wait()
}

// TestWithAttrsImmutability verifies that calling WithAttrs doesn't
// modify the original handler, ensuring proper immutability.
func TestWithAttrsImmutability(t *testing.T) {
	var buf1, buf2 bytes.Buffer

	handler1 := slgk.NewGoKitHandler(log.NewLogfmtLogger(&buf1), nil)

	// Add initial attributes
	handler1 = handler1.WithAttrs([]slog.Attr{
		slog.String("key1", "value1"),
	})

	// Create a new handler with additional attributes
	handler2 := handler1.WithAttrs([]slog.Attr{
		slog.String("key2", "value2"),
	})

	// Verify both handlers exist and are different
	require.NotEqual(t, handler1, handler2, "WithAttrs should return a new handler instance")

	// Create loggers and use them
	logger1 := slog.New(handler1)
	logger2 := slog.New(slgk.NewGoKitHandler(log.NewLogfmtLogger(&buf2), nil).WithAttrs([]slog.Attr{
		slog.String("key1", "value1"),
		slog.String("key2", "value2"),
	}))

	logger1.Info("test1")
	logger2.Info("test2")

	// Both should work without interfering with each other
	require.Contains(t, buf1.String(), "key1=value1")
	require.NotContains(t, buf1.String(), "key2=value2", "Handler 1 should not have key2")
	require.Contains(t, buf2.String(), "key1=value1")
	require.Contains(t, buf2.String(), "key2=value2")
}

// TestCustomLevelMapping verifies that custom slog levels (values between or
// beyond the four standard levels) map to the correct go-kit level string.
// This is a regression test for the range-based level mapping fix: before the
// fix, any non-standard level fell through to "debug".
func TestCustomLevelMapping(t *testing.T) {
	cases := []struct {
		name      string
		level     slog.Level
		wantLevel string
	}{
		{"ExactDebug", slog.LevelDebug, "debug"},
		{"ExactInfo", slog.LevelInfo, "info"},
		{"ExactWarn", slog.LevelWarn, "warn"},
		{"ExactError", slog.LevelError, "error"},
		{"BelowDebug", slog.LevelDebug - 1, "debug"},
		{"BetweenDebugAndInfo", slog.LevelDebug + 1, "debug"},
		{"BetweenInfoAndWarn", slog.LevelInfo + 1, "info"},
		{"BetweenWarnAndError", slog.LevelWarn + 1, "warn"},
		{"AboveError", slog.LevelError + 4, "error"},
	}
	for _, tc := range cases {
		tc := tc // Needed because this library supports pre-go1.22
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			lvl := &slog.LevelVar{}
			lvl.Set(slog.LevelDebug - 1) // Allow all levels through
			h := slgk.NewGoKitHandler(log.NewLogfmtLogger(&buf), lvl)
			logger := slog.New(h)

			logger.Log(context.Background(), tc.level, "test message")

			output := buf.String()
			require.Contains(t, output, fmt.Sprintf("level=%s", tc.wantLevel),
				"level=%s: got output %q", tc.wantLevel, output)
		})
	}
}

// TestZeroValueHandlerPanics verifies that a zero-value GoKitHandler panics
// when Enabled or Handle is called, matching the stdlib pattern where
// slog.JSONHandler and slog.TextHandler also panic when not constructed via
// their constructors. This documents that NewGoKitHandler is required.
//
// WithAttrs and WithGroup don't panic because they only copy fields into a
// new struct without dereferencing pointers. The resulting child handler
// will panic when Enabled or Handle is called on it.
func TestZeroValueHandlerPanics(t *testing.T) {
	t.Run("Enabled", func(t *testing.T) {
		h := &slgk.GoKitHandler{}
		require.Panics(t, func() {
			h.Enabled(context.Background(), slog.LevelInfo)
		})
	})
	t.Run("Handle", func(t *testing.T) {
		h := &slgk.GoKitHandler{}
		require.Panics(t, func() {
			record := slog.NewRecord(time.Now(), slog.LevelInfo, "msg", 0)
			_ = h.Handle(context.Background(), record)
		})
	})
	t.Run("WithAttrs_then_Enabled", func(t *testing.T) {
		h := &slgk.GoKitHandler{}
		child := h.WithAttrs([]slog.Attr{slog.String("k", "v")})
		require.Panics(t, func() {
			child.Enabled(context.Background(), slog.LevelInfo)
		})
	})
	t.Run("WithGroup_then_Handle", func(t *testing.T) {
		h := &slgk.GoKitHandler{}
		child := h.WithGroup("g")
		require.Panics(t, func() {
			record := slog.NewRecord(time.Now(), slog.LevelInfo, "msg", 0)
			_ = child.Handle(context.Background(), record)
		})
	})
}

// passthroughHandler forwards records unchanged, simulating slog middleware
// wrapping the handler. Fixed-depth caller unwinding reports the wrong frame
// under this wrapping; resolution from record.PC must not.
type passthroughHandler struct{ next slog.Handler }

func (p passthroughHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return p.next.Enabled(ctx, l)
}

func (p passthroughHandler) Handle(ctx context.Context, r slog.Record) error {
	return p.next.Handle(ctx, r)
}

func (p passthroughHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return passthroughHandler{next: p.next.WithAttrs(attrs)}
}

func (p passthroughHandler) WithGroup(name string) slog.Handler {
	return passthroughHandler{next: p.next.WithGroup(name)}
}

// TestCaller verifies that the "caller" key is resolved from the record's PC
// and points at the actual log call site.
func TestCaller(t *testing.T) {
	t.Run("direct call site", func(t *testing.T) {
		var buf bytes.Buffer
		slogger := slog.New(slgk.NewGoKitHandler(log.NewLogfmtLogger(&buf), nil))

		_, file, line, _ := runtime.Caller(0)
		slogger.Info("hello") // must stay on the line after the runtime.Caller call
		require.Contains(t, buf.String(), fmt.Sprintf("caller=%s:%d", filepath.Base(file), line+1))
	})
	t.Run("wrapped in middleware handler", func(t *testing.T) {
		var buf bytes.Buffer
		h := slgk.NewGoKitHandler(log.NewLogfmtLogger(&buf), nil)
		slogger := slog.New(passthroughHandler{next: h})

		_, file, line, _ := runtime.Caller(0)
		slogger.Info("hello") // must stay on the line after the runtime.Caller call
		require.Contains(t, buf.String(), fmt.Sprintf("caller=%s:%d", filepath.Base(file), line+1))
	})
	t.Run("record without PC omits caller", func(t *testing.T) {
		var buf bytes.Buffer
		h := slgk.NewGoKitHandler(log.NewLogfmtLogger(&buf), nil)

		record := slog.NewRecord(time.Now(), slog.LevelInfo, "direct handle call", 0)
		require.NoError(t, h.Handle(context.Background(), record))
		require.NotContains(t, buf.String(), "caller=")
	})
}

// TestCallerCache verifies that cached caller resolution stays keyed to the
// right call site: repeated logs from one site (cache hits after the first
// resolution) and logs from distinct sites must each report their own
// file:line.
func TestCallerCache(t *testing.T) {
	var buf bytes.Buffer
	slogger := slog.New(slgk.NewGoKitHandler(log.NewLogfmtLogger(&buf), nil))

	_, file, line, _ := runtime.Caller(0)
	slogger.Info("first")  // cache miss, resolves and stores this PC
	slogger.Info("second") // distinct call site, its own PC and line
	for i := 0; i < 3; i++ {
		slogger.Info("loop") // misses once, then hits the cached PC
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	require.Len(t, lines, 5)

	base := filepath.Base(file)
	require.Contains(t, lines[0], fmt.Sprintf("caller=%s:%d", base, line+1))
	require.Contains(t, lines[1], fmt.Sprintf("caller=%s:%d", base, line+2))
	for _, logLine := range lines[2:] {
		require.Contains(t, logLine, fmt.Sprintf("caller=%s:%d", base, line+4))
	}
}

// stringValuer is a LogValuer that resolves to a plain string value.
type stringValuer struct{ s string }

func (v stringValuer) LogValue() slog.Value { return slog.StringValue(v.s) }

// groupValuer is a LogValuer that resolves to a group, the case where
// resolve-then-expand order is observable: stdlib handlers expand the
// resolved group into individual keys rather than emitting one value.
type groupValuer struct{}

func (groupValuer) LogValue() slog.Value {
	return slog.GroupValue(slog.String("name", "gopher"), slog.Int("age", 42))
}

// TestLogValuerResolution verifies that LogValuer attrs are resolved before
// being emitted, and that a LogValuer resolving to a group is expanded into
// individual dotted keys exactly like a literal group attr, matching the
// resolve-then-expand order of the stdlib handlers.
func TestLogValuerResolution(t *testing.T) {
	t.Run("resolves to scalar", func(t *testing.T) {
		var buf bytes.Buffer
		slogger := slog.New(slgk.NewGoKitHandler(log.NewLogfmtLogger(&buf), nil))

		slogger.Info("hello", "who", stringValuer{s: "gopher"})
		require.Contains(t, buf.String(), "who=gopher")
	})
	t.Run("resolves to group", func(t *testing.T) {
		var buf bytes.Buffer
		slogger := slog.New(slgk.NewGoKitHandler(log.NewLogfmtLogger(&buf), nil))

		slogger.Info("hello", "req", groupValuer{})
		require.Contains(t, buf.String(), "req.name=gopher")
		require.Contains(t, buf.String(), "req.age=42")
	})
	t.Run("resolves to group under WithGroup prefix", func(t *testing.T) {
		var buf bytes.Buffer
		slogger := slog.New(slgk.NewGoKitHandler(log.NewLogfmtLogger(&buf), nil)).WithGroup("outer")

		slogger.Info("hello", "req", groupValuer{})
		require.Contains(t, buf.String(), "outer.req.name=gopher")
		require.Contains(t, buf.String(), "outer.req.age=42")
	})
	t.Run("resolves through WithAttrs", func(t *testing.T) {
		var buf bytes.Buffer
		slogger := slog.New(slgk.NewGoKitHandler(log.NewLogfmtLogger(&buf), nil).WithAttrs([]slog.Attr{
			slog.Any("who", stringValuer{s: "gopher"}),
			slog.Any("req", groupValuer{}),
		}))

		slogger.Info("hello")
		require.Contains(t, buf.String(), "who=gopher")
		require.Contains(t, buf.String(), "req.name=gopher")
		require.Contains(t, buf.String(), "req.age=42")
	})
}

// TestConcurrentHandle drives Handle() concurrently through a single handler
// from multiple call sites. Under -race this covers the caller cache's
// concurrent paths: goroutines racing to first-resolve the same PC and
// lock-free reads on cache hits. TestWithAttrsConcurrency only exercises
// WithAttrs, so without this test the Handle path has no race coverage at
// all. Every emitted line is checked against the set of known call sites.
func TestConcurrentHandle(t *testing.T) {
	var buf bytes.Buffer
	// SyncWriter serializes the underlying writes so concurrently logged
	// lines cannot interleave mid-line.
	slogger := slog.New(slgk.NewGoKitHandler(log.NewLogfmtLogger(log.NewSyncWriter(&buf)), nil))

	_, file, line, _ := runtime.Caller(0)
	sites := []func(){
		func() { slogger.Info("concurrent site a") },
		func() { slogger.Info("concurrent site b") },
		func() { slogger.Info("concurrent site c") },
	}

	const numGoroutines = 8
	const iterations = 200

	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	for g := 0; g < numGoroutines; g++ {
		go func(id int) {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				sites[(id+i)%len(sites)]()
			}
		}(g)
	}
	wg.Wait()

	base := filepath.Base(file)
	want := []string{
		fmt.Sprintf("caller=%s:%d", base, line+2),
		fmt.Sprintf("caller=%s:%d", base, line+3),
		fmt.Sprintf("caller=%s:%d", base, line+4),
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	require.Len(t, lines, numGoroutines*iterations)
	for _, logLine := range lines {
		matched := false
		for _, w := range want {
			if strings.Contains(logLine, w) {
				matched = true
				break
			}
		}
		require.True(t, matched, "log line has unexpected caller: %s", logLine)
	}
}
