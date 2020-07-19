package palette2048_test

import (
	"github.com/reiver/go-sprite32x32"

	"github.com/reiver/go-palette2048"

	"testing"
)

func TestPalette(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var x sprite32x32.Palette = palette2048.Slice{}

	if nil == x {
		t.Error("This should never happen")
	}
}
