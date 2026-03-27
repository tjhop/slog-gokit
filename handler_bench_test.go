package sloggokit_test

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"testing"

	"github.com/go-kit/log"
	slgk "github.com/tjhop/slog-gokit"
)

// BenchmarkBasicLog measures simple log calls with varying numbers of attributes
func BenchmarkBasicLog(b *testing.B) {
	scenarios := []struct {
		name  string
		attrs int
	}{
		{"NoAttrs", 0},
		{"2Attrs", 2},
		{"5Attrs", 5},
		{"10Attrs", 10},
		{"20Attrs", 20},
	}

	for _, sc := range scenarios {
		sc := sc // Needed because this library supports pre-go1.22
		b.Run(sc.name, func(b *testing.B) {
			h := slgk.NewGoKitHandler(log.NewLogfmtLogger(io.Discard), nil)
			logger := slog.New(h)

			attrs := make([]any, 0, sc.attrs*2)
			for i := 0; i < sc.attrs; i++ {
				attrs = append(attrs, fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
			}

			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				logger.Info("benchmark message", attrs...)
			}
		})
	}
}

// BenchmarkLogLevels measures logging at different levels. Each sub-benchmark
// sets the handler's minimum level to match the level being tested, so all
// four cases exercise actual log emission (not the Enabled() fast-path
// rejection that would occur if Debug were run against a default Info-level
// handler).
func BenchmarkLogLevels(b *testing.B) {
	cases := []struct {
		name     string
		minLevel slog.Level
		fn       func(*slog.Logger, string, ...any)
	}{
		{"Debug", slog.LevelDebug, (*slog.Logger).Debug},
		{"Info", slog.LevelInfo, (*slog.Logger).Info},
		{"Warn", slog.LevelWarn, (*slog.Logger).Warn},
		{"Error", slog.LevelError, (*slog.Logger).Error},
	}

	for _, tc := range cases {
		tc := tc // Needed because this library supports pre-go1.22
		b.Run(tc.name, func(b *testing.B) {
			handlerLevel := &slog.LevelVar{}
			handlerLevel.Set(tc.minLevel)
			h := slgk.NewGoKitHandler(log.NewLogfmtLogger(io.Discard), handlerLevel)
			logger := slog.New(h)

			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				tc.fn(logger, "benchmark message", "key1", "value1", "key2", "value2")
			}
		})
	}
}

// BenchmarkDisabledLogs measures overhead of disabled log levels
func BenchmarkDisabledLogs(b *testing.B) {
	lvl := &slog.LevelVar{}
	lvl.Set(slog.LevelError) // Only error logs enabled

	h := slgk.NewGoKitHandler(log.NewLogfmtLogger(io.Discard), lvl)
	logger := slog.New(h)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		logger.Debug("this should be filtered", "key1", "value1", "key2", "value2")
	}
}

// BenchmarkWithAttrsChaining measures repeated WithAttrs calls
func BenchmarkWithAttrsChaining(b *testing.B) {
	depths := []int{1, 3, 5, 10, 20}

	for _, depth := range depths {
		depth := depth // Needed because this library supports pre-go1.22
		b.Run(fmt.Sprintf("Depth%d", depth), func(b *testing.B) {
			h := slgk.NewGoKitHandler(log.NewLogfmtLogger(io.Discard), nil)
			logger := slog.New(h)

			// Create a chain of loggers with attributes
			for i := 0; i < depth; i++ {
				logger = logger.With(
					slog.String(fmt.Sprintf("chain%d", i), fmt.Sprintf("value%d", i)),
				)
			}

			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				logger.Info("benchmark message", "key", "value")
			}
		})
	}
}

// BenchmarkWithGroupNesting measures nested group performance
func BenchmarkWithGroupNesting(b *testing.B) {
	depths := []int{1, 3, 5, 10}

	for _, depth := range depths {
		depth := depth // Needed because this library supports pre-go1.22
		b.Run(fmt.Sprintf("Depth%d", depth), func(b *testing.B) {
			h := slgk.NewGoKitHandler(log.NewLogfmtLogger(io.Discard), nil)
			logger := slog.New(h)

			// Create nested groups
			for i := 0; i < depth; i++ {
				logger = logger.WithGroup(fmt.Sprintf("group%d", i))
			}

			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				logger.Info("benchmark message", "key", "value")
			}
		})
	}
}

// BenchmarkMixedWithAttrsAndGroups measures realistic mixed usage
func BenchmarkMixedWithAttrsAndGroups(b *testing.B) {
	h := slgk.NewGoKitHandler(log.NewLogfmtLogger(io.Discard), nil)

	// Simulate realistic application logger setup:
	// Base logger -> with service attrs -> with request group -> with request attrs
	logger := slog.New(h).
		With(slog.String("service", "api"), slog.String("version", "1.0.0")).
		WithGroup("request").
		With(slog.String("id", "req-123"), slog.String("method", "POST"))

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		logger.Info("processing request", "endpoint", "/api/users", "status", 200)
	}
}

// BenchmarkAttributeTypes measures impact of different value types
func BenchmarkAttributeTypes(b *testing.B) {
	// Built once outside the table so the allocation does not pollute per-iteration measurements.
	largeString := strings.Repeat("x", 1024) // 1KB string

	types := []struct {
		name string
		fn   func(*slog.Logger)
	}{
		{"Strings", func(l *slog.Logger) {
			l.Info("msg", "k1", "v1", "k2", "v2", "k3", "v3")
		}},
		{"Ints", func(l *slog.Logger) {
			l.Info("msg", "k1", 1, "k2", 2, "k3", 3)
		}},
		{"Mixed", func(l *slog.Logger) {
			l.Info("msg", "str", "value", "int", 42, "bool", true, "float", 3.14)
		}},
		{"LargeStrings", func(l *slog.Logger) {
			l.Info("msg", "k1", largeString, "k2", largeString)
		}},
		{"GroupAttr", func(l *slog.Logger) {
			l.Info("msg", slog.Group("g", slog.String("k1", "v1"), slog.Int("k2", 2)))
		}},
		{"NestedGroups", func(l *slog.Logger) {
			l.Info("msg",
				slog.Group("outer",
					slog.String("k1", "v1"),
					slog.Group("inner",
						slog.String("k2", "v2"),
						slog.Int("k3", 3),
					),
				),
			)
		}},
	}

	for _, tc := range types {
		tc := tc // Needed because this library supports pre-go1.22
		b.Run(tc.name, func(b *testing.B) {
			h := slgk.NewGoKitHandler(log.NewLogfmtLogger(io.Discard), nil)
			logger := slog.New(h)

			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				tc.fn(logger)
			}
		})
	}
}

// BenchmarkConcurrentLogging measures performance under concurrent access.
// b.SetParallelism(scale) is a GOMAXPROCS multiplier, not an absolute goroutine
// count. Sub-benchmark names reflect this: Scale1x spawns 1*GOMAXPROCS
// goroutines, Scale2x spawns 2*GOMAXPROCS, and so on.
func BenchmarkConcurrentLogging(b *testing.B) {
	scalingFactors := []int{1, 2, 4, 8, 16}

	for _, scale := range scalingFactors {
		scale := scale // Needed because this library supports pre-go1.22
		b.Run(fmt.Sprintf("Scale%dx", scale), func(b *testing.B) {
			h := slgk.NewGoKitHandler(log.NewLogfmtLogger(io.Discard), nil)
			logger := slog.New(h).With("service", "test")

			b.SetParallelism(scale)
			b.ResetTimer()
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				i := 0
				for pb.Next() {
					logger.Info("concurrent message", "worker", i, "key", "value")
					i++
				}
			})
		})
	}
}

// BenchmarkConcurrentWithAttrsThenLog measures the combined cost of WithAttrs followed by a Log
// call under concurrent load. The logger.Info() call is intentional: it acts as an escape sink
// that prevents the compiler from eliminating the WithAttrs allocation, making this a realistic
// measure of the goroutine-per-request pattern where each goroutine derives a child logger and
// immediately logs with it.
func BenchmarkConcurrentWithAttrsThenLog(b *testing.B) {
	h := slgk.NewGoKitHandler(log.NewLogfmtLogger(io.Discard), nil)
	baseLogger := slog.New(h)

	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			logger := baseLogger.With(
				slog.Int("worker", i),
				slog.String("key", "value"),
			)
			logger.Info("message")
			i++
		}
	})
}

// BenchmarkHandleOnly isolates Handle() with varying amounts of pre-flattened
// attributes and zero record-level attributes. This highlights the cost of the
// bulk-append path (append(pairs, h.preformatted...)) without noise from
// per-attr processing on the record side.
func BenchmarkHandleOnly(b *testing.B) {
	preformattedCounts := []int{0, 5, 20}

	for _, n := range preformattedCounts {
		n := n // Needed because this library supports pre-go1.22
		b.Run(fmt.Sprintf("Preformatted%d", n), func(b *testing.B) {
			h := slgk.NewGoKitHandler(log.NewLogfmtLogger(io.Discard), nil)
			logger := slog.New(h)

			// Build up pre-flattened attrs via With()
			for i := 0; i < n; i++ {
				logger = logger.With(
					slog.String(fmt.Sprintf("pre%d", i), fmt.Sprintf("val%d", i)),
				)
			}

			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				// Log with no record-level attrs -- only the
				// pre-flattened bulk-append path is exercised.
				logger.Info("benchmark message")
			}
		})
	}
}

// BenchmarkEnabledCheck micro-benchmarks the Enabled() guard at matching and
// non-matching levels. The enabled case adds the overhead of the subsequent
// Handle() call; the disabled case should be nearly free.
func BenchmarkEnabledCheck(b *testing.B) {
	cases := []struct {
		name     string
		logLevel slog.Level
		minLevel slog.Level
	}{
		{"Enabled_DebugAtDebug", slog.LevelDebug, slog.LevelDebug},
		{"Enabled_InfoAtInfo", slog.LevelInfo, slog.LevelInfo},
		{"Disabled_DebugAtInfo", slog.LevelDebug, slog.LevelInfo},
		{"Disabled_DebugAtError", slog.LevelDebug, slog.LevelError},
		{"Disabled_InfoAtWarn", slog.LevelInfo, slog.LevelWarn},
		{"Disabled_WarnAtError", slog.LevelWarn, slog.LevelError},
	}

	for _, tc := range cases {
		tc := tc // Needed because this library supports pre-go1.22
		b.Run(tc.name, func(b *testing.B) {
			lvl := &slog.LevelVar{}
			lvl.Set(tc.minLevel)
			h := slgk.NewGoKitHandler(log.NewLogfmtLogger(io.Discard), lvl)
			logger := slog.New(h)

			// Pre-build the method ref to ensure we're comparing
			// apples to apples across levels.
			var logFn func(string, ...any)
			switch tc.logLevel {
			case slog.LevelDebug:
				logFn = logger.Debug
			case slog.LevelInfo:
				logFn = logger.Info
			case slog.LevelWarn:
				logFn = logger.Warn
			case slog.LevelError:
				logFn = logger.Error
			}

			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				logFn("benchmark message", "key", "value")
			}
		})
	}
}

// BenchmarkLevelSelection measures the cost of levelLoggerCache.get() across
// all four levels. With the level caching optimization, this should show zero
// allocations (the leveled loggers are pre-built at handler construction time).
func BenchmarkLevelSelection(b *testing.B) {
	levels := []struct {
		name  string
		level slog.Level
	}{
		{"Debug", slog.LevelDebug},
		{"Info", slog.LevelInfo},
		{"Warn", slog.LevelWarn},
		{"Error", slog.LevelError},
	}

	for _, tc := range levels {
		tc := tc // Needed because this library supports pre-go1.22
		b.Run(tc.name, func(b *testing.B) {
			lvl := &slog.LevelVar{}
			lvl.Set(tc.level)
			h := slgk.NewGoKitHandler(log.NewLogfmtLogger(io.Discard), lvl)
			logger := slog.New(h)

			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				// Each iteration exercises the full path: Enabled() + Handle()
				// (including levelLoggers.get()), but with minimal attrs to
				// isolate level selection cost.
				logger.Log(context.Background(), tc.level, "msg")
			}
		})
	}
}
