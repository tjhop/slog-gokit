package sloggokit_test

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// When SLOGGOKIT_PPROF is set (e.g. ":6060"), start a pprof HTTP server
	// so that continuous profilers like Parca can scrape the benchmark process
	// directly.
	if addr := os.Getenv("SLOGGOKIT_PPROF"); addr != "" {
		go func() {
			log.Printf("pprof endpoints available at http://localhost%s/debug/pprof/", addr)
			if err := http.ListenAndServe(addr, nil); err != nil {
				log.Printf("pprof server error: %v", err)
			}
		}()
	}

	os.Exit(m.Run())
}
