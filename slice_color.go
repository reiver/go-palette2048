package palette2048

import (
	"github.com/reiver/go-rgba32"

	"image/color"
)

func (receiver Slice) Color(index uint8) color.Color {
	if nil == receiver {
		return nil
	}

	if ByteSize != len(receiver) {
		return nil
	}

	low  := int(index) * rgba32.ByteSize
	high := low        + rgba32.ByteSize

	p := receiver[low:high]
	if (high-low) != len(p) {
		return nil
	}

	var rgba rgba32.Slice = rgba32.Slice(p)

	return rgba
}
