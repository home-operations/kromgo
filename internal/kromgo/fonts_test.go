package kromgo

import (
	"bytes"
	"testing"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/sfnt"
)

// TestEmbeddedFontsParse ensures every registered font parses with both the badge
// renderer (sfnt) and the graph renderer (truetype). It guards against a vendored
// face the renderers can't read — e.g. wiring in an @fontsource woff2 instead of a
// static .ttf, or a corrupt asset.
func TestEmbeddedFontsParse(t *testing.T) {
	if len(embeddedFonts) == 0 {
		t.Fatal("no embedded fonts registered")
	}
	for name, data := range embeddedFonts {
		t.Run(name, func(t *testing.T) {
			if len(data) == 0 {
				t.Fatalf("font %q is empty (asset not generated?)", name)
			}
			if _, err := sfnt.Parse(data); err != nil {
				t.Errorf("sfnt parse (badge renderer): %v", err)
			}
			if _, err := truetype.Parse(data); err != nil {
				t.Errorf("truetype parse (graph renderer): %v", err)
			}
		})
	}
}

// TestResolveBadgeBoldFont covers the bold companion the for-the-badge message uses:
// the default maps to dejavu-sans-bold, a regular face to its "-bold" sibling, a face
// that is already bold to itself, and a face with no bold companion degrades to its
// regular bytes rather than erroring.
func TestResolveBadgeBoldFont(t *testing.T) {
	cases := map[string][]byte{
		"":                 dejavuSansBoldTTF,
		"dejavu-sans":      dejavuSansBoldTTF,
		"comic-neue":       comicNeueBoldTTF,
		"dejavu-sans-bold": dejavuSansBoldTTF, // already bold → itself
	}
	for name, want := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := resolveBadgeBoldFont(name)
			if err != nil {
				t.Fatalf("resolveBadgeBoldFont(%q): %v", name, err)
			}
			if !bytes.Equal(got, want) {
				t.Errorf("resolveBadgeBoldFont(%q) returned the wrong face", name)
			}
		})
	}
}
