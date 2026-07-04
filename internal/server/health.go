package server

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// healthPaths are the liveness/readiness endpoints, served on the MAIN port
// (the org pair standard: /healthz = liveness, /readyz = readiness — aliases
// here, kromgo has no serving condition beyond being up).
var healthPaths = []string{"/healthz", "/readyz"}

// withHealth overlays the health endpoints onto the application handler, so
// probes work regardless of whether the optional metrics listener is enabled.
func withHealth(app http.Handler) http.Handler {
	mux := http.NewServeMux()
	for _, path := range healthPaths {
		mux.HandleFunc("GET "+path, ok)
	}
	mux.Handle("/", app)
	return mux
}

// metricsMux serves Prometheus metrics on the dedicated, optional metrics
// listener. Health probes are NOT served here — they live on the main port, so
// this whole listener can be disabled without breaking probes.
func metricsMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /metrics", promhttp.Handler())
	return mux
}

func ok(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}
