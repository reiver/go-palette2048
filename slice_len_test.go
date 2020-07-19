package palette2048_test

import (
	"github.com/reiver/go-palette2048"

	"testing"
)

func TestSlice_Len(t *testing.T) {

	const expected = 2048 / 8

	var actual int
	{
		var buffer [palette2048.ByteSize]uint8

		var palette palette2048.Slice = palette2048.Slice(buffer[:])

		actual = palette.Len()
	}

	if expected != actual {
		t.Errorf("The actual length is not what was expected.")
		t.Logf("EXPECTED: %d", expected)
		t.Logf("ACTUAL:   %d", actual)
		return
	}
}
