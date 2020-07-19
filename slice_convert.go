package palette2048

import (
	"image/color"
)

func (receiver Slice) Convert(c color.Color) color.Color {
	if nil == receiver {
		return nil
	}

	if nil == c {
		return nil
	}

	if ByteSize != len(receiver) {
		return nil
	}

	index := receiver.Index(c)
	if receiver.Len() <= int(index) {
		return nil
	}

	return receiver.Color(index)
}
