package palette2048_test

import (
	"github.com/reiver/go-palette2048"

	"image/color"
	"math/rand"
	"time"

	"testing"
)

func TestSlice_SetColor(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	Loop: for testNumber:=0; testNumber<10; testNumber++ {

		var buffer [palette2048.ByteSize]uint8

		for i:=0; i<len(buffer); i++ {
			buffer[i] = uint8(randomness.Intn(256))
		}

		var palette palette2048.Slice = palette2048.Slice(buffer[:])

		for subTestNumber:=0; subTestNumber<5; subTestNumber++ {
			index := uint8(randomness.Intn(palette.Len()))

			r := uint8(randomness.Intn(256))
			g := uint8(randomness.Intn(256))
			b := uint8(randomness.Intn(256))
			a := uint8(255)

			c := color.NRGBA{
				R:r,
				G:g,
				B:b,
				A:a,
			}

			eR, eG, eB, eA := c.RGBA()

			if err := palette.SetColor(index, c); nil != err {
				t.Errorf("For test #%d & sub-test #%d, received an error.", testNumber, subTestNumber)
				t.Logf("INDEX: %d", index)
				t.Logf("rgba(%d,%d,%d,%d)", r,g,b,a)
				t.Logf("EXPECTED: (r,g,b,a)=(%d,%d,%d,%d)", eR, eG, eB, eA)
				t.Logf("ERROR: (%T) %q", err, err)
				continue Loop
			}

			aR, aG, aB, aA := palette.Color(index).RGBA()

			if eR != aR || eG != aG || eB != aB || eA != aA {
				t.Errorf("For test #%d & sub-test #%d, the color did not seem to be set.", testNumber, subTestNumber)
				t.Logf("INDEX: %d", index)
				t.Logf("rgba(%d,%d,%d,%d)", r,g,b,a)
				t.Logf("EXPECTED: (r,g,b,a)=(%d,%d,%d,%d)", eR, eG, eB, eA)
				t.Logf("ACTUAL:   (r,g,b,a)=(%d,%d,%d,%d)", aR, aG, aB, aA)
				continue Loop
			}
		}
	}

}
