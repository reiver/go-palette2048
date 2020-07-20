package palette2048

import (
	"github.com/reiver/go-rgba32"

	"image/color"
)

func (receiver Slice) SetColorRGBA(index uint8, r,g,b,a uint8) error {
	if nil == receiver {
		return errNilReceiver
	}

	p := receiver.color(index)
	if nil == p {
		return errInternalError
	}

	if rgba32.ByteSize != len(p) {
		return errInternalError
	}

	p[rgba32.OffsetRed]   = r
	p[rgba32.OffsetGreen] = g
	p[rgba32.OffsetBlue]  = b
	p[rgba32.OffsetAlpha] = a

	return nil
}

func (receiver Slice) SetColorRGB(index uint8, r,g,b uint8) error {
	a := uint8(255)

	return receiver.SetColorRGBA(index, r,g,b,a)
}

func (receiver Slice) SetColor(index uint8, c color.Color) error {
	if nil == receiver {
		return errNilReceiver
	}

	rr,gg,bb,aa := c.RGBA()

	r := uint8((rr*0xff)/0xffff)
	g := uint8((gg*0xff)/0xffff)
	b := uint8((bb*0xff)/0xffff)
	a := uint8((aa*0xff)/0xffff)

	return receiver.SetColorRGBA(index, r,g,b,a)
}
