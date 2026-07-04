package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

// ServerConfig holds runtime server settings sourced from environment variables.
type ServerConfig struct {
	// Hosts default to empty (the unspecified address), binding IPv4-only,
	// IPv6-only, and dual-stack clusters alike; an explicit 0.0.0.0 would be
	// IPv4-only. The main port also serves /healthz and /readyz.
	ServerHost string `env:"SERVER_HOST" envDefault:""`
	ServerPort int    `env:"SERVER_PORT" envDefault:"8080"`

	// MetricsEnabled exposes Prometheus metrics at /metrics on MetricsPort.
	// Disabling it removes the metrics listener entirely; the health probe
	// endpoints live on the main port and are unaffected.
	MetricsEnabled bool   `env:"METRICS_ENABLED" envDefault:"true"`
	MetricsHost    string `env:"METRICS_HOST" envDefault:""`
	MetricsPort    int    `env:"METRICS_PORT" envDefault:"8081"`

	// ServerReadTimeout / ServerWriteTimeout bound reading a request and writing its
	// response on the public listener; the defaults harden against slow-client
	// connection holding. WriteTimeout must exceed QueryTimeout so a slow upstream
	// isn't cut off mid-render. Set either to "0" to disable (no deadline).
	ServerReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT" envDefault:"15s"`
	ServerWriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT" envDefault:"60s"`
	ServerLogging      bool          `env:"SERVER_LOGGING"`

	// QueryTimeout bounds each outbound Prometheus query.
	QueryTimeout time.Duration `env:"QUERY_TIMEOUT" envDefault:"30s"`
}

// LoadServer reads ServerConfig from the environment.
func LoadServer() (ServerConfig, error) {
	var cfg ServerConfig
	if err := env.Parse(&cfg); err != nil {
		return ServerConfig{}, err
	}
	return cfg, nil
}
