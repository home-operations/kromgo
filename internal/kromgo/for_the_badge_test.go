package kromgo

import (
	"bytes"
	"encoding/xml"
	"io"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/home-operations/kromgo/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// firstGlyphBaselineY returns the y of the first "M<x> <y>" move command in svg —
// for a badge with no icon, the first such command is a text glyph, so its y locates
// the text baseline. Used to guard against a wildly-off / sign-flipped baseline.
func firstGlyphBaselineY(t *testing.T, svg string) float64 {
	t.Helper()
	m := regexp.MustCompile(`M[\d.]+ ([\d.]+)`).FindStringSubmatch(svg)
	require.Len(t, m, 2, "expected at least one glyph move command")
	y, err := strconv.ParseFloat(m[1], 64)
	require.NoError(t, err)
	return y
}

func requireWellFormed(t *testing.T, svg []byte) {
	t.Helper()
	dec := xml.NewDecoder(bytes.NewReader(svg))
	for {
		_, err := dec.Token()
		if err == io.EOF {
			break
		}
		require.NoError(t, err, "rendered SVG must be well-formed XML")
	}
}

func TestBadgeRender_ForTheBadge(t *testing.T) {
	t.Parallel()
	r, err := newBadgeRenderer(config.BadgeDefaults{})
	require.NoError(t, err)

	// Mixed-case input so the original-case accessible label is distinguishable from the
	// uppercased glyphs — pins the "never mutate spec" invariant.
	svg := r.render(badgeSpec{style: config.StyleForTheBadge, label: "build", message: "passing", color: "green", id: "b"})
	s := string(svg)

	requireWellFormed(t, svg)
	assert.Contains(t, s, `height="28"`, "for-the-badge is 28px tall")
	assert.Contains(t, s, `viewBox="0 0 `)
	assert.Regexp(t, `viewBox="0 0 \d+ 28"`, s)
	assert.Contains(t, s, `rx="0"`, "square corners")
	assert.NotContains(t, s, "<linearGradient", "no gloss gradient")
	assert.Contains(t, s, `aria-label="build: passing"`, "accessible label stays original-case")
	assert.NotContains(t, s, "<text", "text is glyph paths, not <text>")

	// Two segments (clip rect + label rect + message rect).
	assert.Equal(t, 3, strings.Count(s, "<rect"), "label + message segments (plus the clip rect)")

	// Strictly wider than the same flat badge: uppercase + 1.25 tracking + 12px margins.
	flat := r.render(badgeSpec{style: config.StyleFlat, label: "build", message: "passing", color: "green", id: "b"})
	assert.Greater(t, svgWidth(t, svg), svgWidth(t, flat), "for-the-badge is wider than flat")

	// Baseline sits inside the 28px box (catches a negated/wildly-off baseline).
	y := firstGlyphBaselineY(t, s)
	assert.Greater(t, y, 7.0)
	assert.Less(t, y, 21.0)
}

func TestBadgeRender_ForTheBadge_IconNoLabel(t *testing.T) {
	t.Parallel()
	r, err := newBadgeRenderer(config.BadgeDefaults{})
	require.NoError(t, err)
	icon, err := resolveIcon("si:kubernetes")
	require.NoError(t, err)

	// The home-ops hero shape: icon + message, no label, no labelColor → a single
	// colored segment with the logo riding on it.
	svg := r.render(badgeSpec{style: config.StyleForTheBadge, iconPath: icon, message: "v1.30.2", color: "blue", id: "k8s"})
	s := string(svg)

	requireWellFormed(t, svg)
	assert.Contains(t, s, `height="28"`)
	assert.Contains(t, s, icon, "icon path embedded")
	assert.Equal(t, 2, strings.Count(s, "<rect"), "single message segment (plus the clip rect), no label box")
	assert.Contains(t, s, `<rect x="0"`, "the single segment spans from the left edge")
	assert.Contains(t, s, "#007ec6", "blue message segment")
}

func TestBadgeRender_ForTheBadge_IconLabelColor(t *testing.T) {
	t.Parallel()
	r, err := newBadgeRenderer(config.BadgeDefaults{})
	require.NoError(t, err)
	icon, err := resolveIcon("si:kubernetes")
	require.NoError(t, err)

	// Icon + explicit labelColor but no label → a colored logo box (labelRectWidth =
	// 2*LOGO_MARGIN + logoWidth = 2*9 + 14 = 32) followed by the message segment.
	svg := r.render(badgeSpec{style: config.StyleForTheBadge, iconPath: icon, message: "v1.30.2", color: "blue", labelColor: "purple", id: "k8s"})
	s := string(svg)

	requireWellFormed(t, svg)
	assert.Equal(t, 3, strings.Count(s, "<rect"), "logo box + message segment (plus the clip rect)")
	assert.Contains(t, s, `<rect width="32" height="28"`, "logo box is 2*9+14 = 32px wide")
}

func TestBadgeRender_ForTheBadge_NoText(t *testing.T) {
	t.Parallel()
	r, err := newBadgeRenderer(config.BadgeDefaults{})
	require.NoError(t, err)
	icon, err := resolveIcon("si:kubernetes")
	require.NoError(t, err)

	// Icon, empty message, no label exercises the noText gutter = LOGO_TEXT_GUTTER -
	// LOGO_MARGIN (= -3) branch. It must still produce a well-formed, positive-width 28px
	// badge with no panic.
	svg := r.render(badgeSpec{style: config.StyleForTheBadge, iconPath: icon, message: "", color: "blue", id: "k8s"})
	requireWellFormed(t, svg)
	assert.Contains(t, string(svg), `height="28"`)
	assert.Greater(t, svgWidth(t, svg), 0)
}

func TestBadgeRender_ForTheBadge_BoldMessage(t *testing.T) {
	t.Parallel()
	r, err := newBadgeRenderer(config.BadgeDefaults{})
	require.NoError(t, err)
	icon, err := resolveIcon("si:kubernetes")
	require.NoError(t, err)

	// The message segment must be drawn with the bold companion face. For the hero shape
	// the message text is inset at TEXT_MARGIN + iconSize + gutter = 12 + 14 + 6 = 32.
	const msgX = ftbTextMargin + ftbIconSize + ftbLogoGutter
	boldPath := r.glyphPathFace(r.boldFont, "PASSING", msgX, ftbBaseline, ftbFontSize, ftbTracking)
	regularPath := r.glyphPathFace(r.font, "PASSING", msgX, ftbBaseline, ftbFontSize, ftbTracking)
	require.NotEqual(t, regularPath, boldPath, "bold and regular faces must differ for this text")

	// Pass lowercase to also prove uppercasing: the rendered glyphs equal the bold path
	// for "PASSING".
	svg := string(r.render(badgeSpec{style: config.StyleForTheBadge, iconPath: icon, message: "passing", color: "blue", id: "x"}))
	assert.Contains(t, svg, boldPath, "message rendered with the bold face (and uppercased)")
	assert.NotContains(t, svg, regularPath, "message not rendered with the regular face")
}

// svgWidth extracts the integer width="" attribute from a rendered SVG.
func svgWidth(t *testing.T, svg []byte) int {
	t.Helper()
	m := regexp.MustCompile(`width="(\d+)"`).FindSubmatch(svg)
	require.Len(t, m, 2, "svg has a width attribute")
	w, err := strconv.Atoi(string(m[1]))
	require.NoError(t, err)
	return w
}
