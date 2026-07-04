package config

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func writeConfig(t *testing.T, body string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "config.yaml")
	require.NoError(t, os.WriteFile(path, []byte(body), 0o600))
	return path
}

func TestLoad_Valid(t *testing.T) {
	t.Parallel()
	path := writeConfig(t, `
prometheus: http://prom:9090
gallery:
  enabled: true
defaults:
  graph:
    maxDuration: 7d
    gallery:
      hidden: true
badges:
  - id: cpu
    query: node_cpu
    valueExpr: string(result) + "%"
    gallery:
      hidden: false
graphs:
  - id: cpu
    query: node_cpu
    width: 400
`)
	cfg, err := Load(path)
	require.NoError(t, err)
	assert.Equal(t, "http://prom:9090", cfg.Prometheus)
	require.NotNil(t, cfg.Gallery.Enabled)
	assert.True(t, *cfg.Gallery.Enabled)
	require.Len(t, cfg.Badges, 1)
	assert.Equal(t, "cpu", cfg.Badges[0].ID)
	require.NotNil(t, cfg.Badges[0].Gallery.Hidden)
	assert.False(t, *cfg.Badges[0].Gallery.Hidden, "per-badge gallery.hidden")
	require.NotNil(t, cfg.Defaults.Graph.Gallery.Hidden)
	assert.True(t, *cfg.Defaults.Graph.Gallery.Hidden, "per-type default gallery.hidden")
	require.Len(t, cfg.Graphs, 1)
	assert.Equal(t, 400, cfg.Graphs[0].Width)
	assert.Equal(t, "7d", cfg.Defaults.Graph.MaxDuration)
}

func TestLoad_MissingFile(t *testing.T) {
	t.Parallel()
	_, err := Load(filepath.Join(t.TempDir(), "nope.yaml"))
	assert.Error(t, err)
}

func TestLoad_InvalidYAML(t *testing.T) {
	t.Parallel()
	_, err := Load(writeConfig(t, "badges: [: bad"))
	assert.Error(t, err)
}

func TestLoad_RejectsUnknownKey(t *testing.T) {
	t.Parallel()
	// A non-legacy typo'd key must error under strict decoding.
	_, err := Load(writeConfig(t, "badges: []\nbogus: true\n"))
	assert.Error(t, err)
}

func TestLoad_LegacyConfigErrors(t *testing.T) {
	t.Parallel()
	// A pre-0.12 config (top-level metrics:) gets a pointed migration error.
	_, err := Load(writeConfig(t, "metrics:\n  - name: cpu\n    query: q\n"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "0.12")
}

func TestLoad_InvalidDuration(t *testing.T) {
	t.Parallel()
	_, err := Load(writeConfig(t, "defaults:\n  graph:\n    maxDuration: bogus\n"))
	assert.Error(t, err)
}

func TestLoad_DuplicateID(t *testing.T) {
	t.Parallel()
	_, err := Load(writeConfig(t, "badges:\n  - id: cpu\n    query: q\n  - id: cpu\n    query: q\n"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate")
}

func TestLoad_RangeBadgeRequiresLast(t *testing.T) {
	t.Parallel()
	_, err := Load(writeConfig(t, "badges:\n  - id: cpu\n    query: q\n    type: range\n"))
	assert.Error(t, err)
}

func TestLoad_InvalidID(t *testing.T) {
	t.Parallel()
	// ids are URL path segments and gallery Markdown; reject unsafe characters.
	_, err := Load(writeConfig(t, "badges:\n  - id: a/b\n    query: q\n"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "id must match")

	_, err = Load(writeConfig(t, "graphs:\n  - id: 'a b'\n    query: q\n"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "id must match")
}

func TestLoad_Style(t *testing.T) {
	t.Parallel()
	// Every supported style passes validation, both per-badge and as a default.
	for _, style := range []string{StyleFlat, StyleFlatSquare, StylePlastic, StyleForTheBadge} {
		_, err := Load(writeConfig(t, "badges:\n  - id: cpu\n    query: q\n    style: "+style+"\n"))
		require.NoErrorf(t, err, "badge style %q should be valid", style)
		_, err = Load(writeConfig(t, "defaults:\n  badge:\n    style: "+style+"\nbadges:\n  - id: cpu\n    query: q\n"))
		require.NoErrorf(t, err, "default badge style %q should be valid", style)
	}

	// An unknown style is still rejected (guards ValidStyle staying authoritative).
	_, err := Load(writeConfig(t, "badges:\n  - id: cpu\n    query: q\n    style: for-the-win\n"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unknown style")
}

func TestLoadServer_Defaults(t *testing.T) {
	// Clear any inherited env so envDefault applies. t.Setenv registers the
	// restore; os.Unsetenv then removes the var for the duration of the test.
	for _, k := range []string{"SERVER_PORT", "METRICS_ENABLED", "METRICS_PORT", "QUERY_TIMEOUT", "SERVER_LOGGING", "SERVER_READ_TIMEOUT", "SERVER_WRITE_TIMEOUT"} {
		t.Setenv(k, "")
		_ = os.Unsetenv(k)
	}
	sc, err := LoadServer()
	require.NoError(t, err)
	// Empty host = the unspecified address (dual-stack), not IPv4-only 0.0.0.0.
	assert.Empty(t, sc.ServerHost)
	assert.Equal(t, 8080, sc.ServerPort)
	assert.True(t, sc.MetricsEnabled)
	assert.Equal(t, 8081, sc.MetricsPort)
	assert.Equal(t, 30*time.Second, sc.QueryTimeout)
	assert.Equal(t, 15*time.Second, sc.ServerReadTimeout, "read timeout defaults to a bounded value, not 0")
	assert.Equal(t, 60*time.Second, sc.ServerWriteTimeout, "write timeout defaults above QueryTimeout")
	assert.False(t, sc.ServerLogging)
}

func TestLoadServer_Overrides(t *testing.T) {
	t.Setenv("SERVER_PORT", "9000")
	t.Setenv("QUERY_TIMEOUT", "5s")
	t.Setenv("SERVER_LOGGING", "true")
	t.Setenv("SERVER_READ_TIMEOUT", "5s")
	t.Setenv("SERVER_WRITE_TIMEOUT", "0") // 0 disables the deadline
	sc, err := LoadServer()
	require.NoError(t, err)
	assert.Equal(t, 9000, sc.ServerPort)
	assert.Equal(t, 5*time.Second, sc.QueryTimeout)
	assert.True(t, sc.ServerLogging)
	assert.Equal(t, 5*time.Second, sc.ServerReadTimeout)
	assert.Equal(t, time.Duration(0), sc.ServerWriteTimeout, "0 disables the write deadline")
}
