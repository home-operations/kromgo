package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/home-operations/kromgo/internal/config"
	"github.com/home-operations/kromgo/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSecureHeaders(t *testing.T) {
	t.Parallel()
	h := secureHeaders(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	assert.Equal(t, "nosniff", w.Header().Get("X-Content-Type-Options"))
	assert.Contains(t, w.Header().Get("Content-Security-Policy"), "default-src 'none'")
}

// TestWithHealth pins the pair standard: health probes ride the MAIN handler
// (so the optional metrics listener can be disabled without breaking probes),
// and the app still receives everything else.
func TestWithHealth(t *testing.T) {
	t.Parallel()
	appHit := false
	h := withHealth(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		appHit = true
		w.WriteHeader(http.StatusOK)
	}))
	for _, path := range []string{"/healthz", "/readyz", "/-/health", "/-/ready"} {
		t.Run(path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, path, nil)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "OK", w.Body.String())
		})
	}
	// Non-health paths fall through to the app.
	h.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/badges/x", nil))
	assert.True(t, appHit)
}

// The metrics mux is metrics-only: health lives on the main port.
func TestMetricsMux(t *testing.T) {
	t.Parallel()
	mux := metricsMux()
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/metrics", nil))
	assert.Equal(t, http.StatusOK, w.Code)
	w = httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestRecoverer_TurnsPanicInto500(t *testing.T) {
	t.Parallel()
	h := recoverer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		panic("boom")
	}))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/", nil))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestAccessLog_PassesThrough(t *testing.T) {
	t.Parallel()
	h := accessLog(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		_, _ = w.Write([]byte("hi"))
	}))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/", nil))
	assert.Equal(t, http.StatusTeapot, w.Code)
	assert.Equal(t, "hi", w.Body.String())
}

func TestWithMiddleware_LoggingOptional(t *testing.T) {
	t.Parallel()
	called := false
	base := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	for _, logging := range []bool{false, true} {
		called = false
		h := withMiddleware(base, config.ServerConfig{ServerLogging: logging})
		w := httptest.NewRecorder()
		h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/", nil))
		assert.Equal(t, http.StatusOK, w.Code)
		assert.True(t, called)
	}
}

func TestRun_GracefulShutdown(t *testing.T) {
	t.Parallel()
	sc := config.ServerConfig{
		ServerHost: "127.0.0.1", ServerPort: testutil.FreePort(t),
		MetricsEnabled: true,
		MetricsHost:    "127.0.0.1", MetricsPort: testutil.FreePort(t),
	}
	app := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- Run(ctx, sc, app) }()

	// Wait until the main server is serving (health rides it), then trigger
	// graceful shutdown.
	healthURL := fmt.Sprintf("http://127.0.0.1:%d/healthz", sc.ServerPort)
	require.Eventually(t, func() bool {
		resp, err := http.Get(healthURL)
		if err != nil {
			return false
		}
		_ = resp.Body.Close()
		return resp.StatusCode == http.StatusOK
	}, 3*time.Second, 20*time.Millisecond)

	cancel()
	select {
	case err := <-done:
		assert.NoError(t, err)
	case <-time.After(3 * time.Second):
		t.Fatal("Run did not return after context cancellation")
	}
}
