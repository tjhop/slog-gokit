package sloggokit

import (
	"context"
	"io"
	"log/slog"
	"runtime"
	"testing"
	"time"

	"github.com/go-kit/log"
)

// BenchmarkCallerCacheMiss forces caller resolution on every iteration by
// deleting the cached PC entry before each Handle() call, approximating the
// pre-cache per-call cost (the CallersFrames symbolization walk plus string
// build and Store). The sync.Map Delete is included in the measurement, but
// it is cheap relative to the resolution it re-triggers.
//
// This benchmark must live in the internal test package: forcing a miss
// requires evicting from the unexported callerCache, and there is no public
// eviction API (the nearest public approximation, a unique synthesized PC
// per iteration, would grow the cache without bound mid-benchmark and
// measure map growth instead of the miss path). It is a same-version
// diagnostic only -- keep it out of cross-version comparisons; the
// public-API counterparts (BenchmarkCallerCache Hit/NoPC and
// BenchmarkCallerCacheSites) live in handler_bench_test.go and are mirrored
// in the benchmarks/ module.
func BenchmarkCallerCacheMiss(b *testing.B) {
	h := NewGoKitHandler(log.NewLogfmtLogger(io.Discard), nil)
	ctx := context.Background()

	pcs := make([]uintptr, 1)
	runtime.Callers(1, pcs)
	pc := pcs[0]

	record := slog.NewRecord(time.Now(), slog.LevelInfo, "benchmark message", pc)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		callerCache.Delete(pc)
		_ = h.Handle(ctx, record)
	}
}
