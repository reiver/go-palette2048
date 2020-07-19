package palette2048_test

import (
	"github.com/reiver/go-palette2048"

	"image/color"

	"testing"
)

func TestColorModel(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var x color.Model = palette2048.Slice{}

	if nil == x {
		t.Error("This should never happen")
	}
}
