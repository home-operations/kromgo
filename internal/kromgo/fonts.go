package kromgo

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/golang/freetype/truetype"
)

// Fonts are compiled into the binary, never read from disk — the image is scratch
// and we control the set. DejaVu Sans is the default for badges and graphs (the free,
// metric-compatible stand-in for the Verdana shields.io renders with); Comic Neue is a
// second selectable face. Both are vendored via npm and generated into assets/ by
// cmd/genassets (regular + bold). Add a face by vendoring it (npm) and PRing it here.

//go:embed assets/dejavu-sans.ttf
var dejavuSansTTF []byte

//go:embed assets/dejavu-sans-bold.ttf
var dejavuSansBoldTTF []byte

//go:embed assets/comic-neue.ttf
var comicNeueTTF []byte

//go:embed assets/comic-neue-bold.ttf
var comicNeueBoldTTF []byte

var embeddedFonts = map[string][]byte{
	"dejavu-sans":      dejavuSansTTF,
	"dejavu-sans-bold": dejavuSansBoldTTF,
	"comic-neue":       comicNeueTTF,
	"comic-neue-bold":  comicNeueBoldTTF,
}

// resolveBadgeFont returns the TTF bytes for a badge font name (empty = the default
// DejaVu Sans face). The badge renderer parses these bytes with sfnt to draw glyph paths.
func resolveBadgeFont(name string) ([]byte, error) {
	if name == "" {
		return dejavuSansTTF, nil
	}
	if data := embeddedFonts[name]; data != nil {
		return data, nil
	}
	return nil, fmt.Errorf("unknown font %q", name)
}

// resolveBadgeBoldFont returns the TTF bytes for the bold companion of a badge font
// name — the face the for-the-badge style draws its (bold) message segment with. The
// companion is "<name>-bold" (empty = dejavu-sans-bold, the default face's bold). A
// name that is already a "-bold" face is its own companion. A face with no bundled
// bold companion degrades gracefully to its regular bytes rather than erroring, so a
// future regular-only face still renders for-the-badge (just not bold).
func resolveBadgeBoldFont(name string) ([]byte, error) {
	if name == "" {
		return dejavuSansBoldTTF, nil
	}
	if strings.HasSuffix(name, "-bold") {
		return resolveBadgeFont(name)
	}
	if data := embeddedFonts[name+"-bold"]; data != nil {
		return data, nil
	}
	return resolveBadgeFont(name) // no bold companion — fall back to the regular face
}

// resolveGraphFont returns the parsed font for a graph font name (empty = the default
// DejaVu Sans face). Graphs render their text through the chart library with this font.
func resolveGraphFont(name string) (*truetype.Font, error) {
	if name == "" {
		return truetype.Parse(dejavuSansTTF)
	}
	if data := embeddedFonts[name]; data != nil {
		return truetype.Parse(data)
	}
	return nil, fmt.Errorf("unknown font %q", name)
}
