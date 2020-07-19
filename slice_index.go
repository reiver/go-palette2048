package palette2048

import (
	"image/color"
)

func (receiver Slice) Index(c color.Color) uint8 {
	const defaultIndex = 0

	if nil == c {
		return defaultIndex
	}

	if ByteSize != len(receiver) {
		return defaultIndex
	}

	r,g,b,a := c.RGBA()

	var favoredIndex uint8
	var favoredDistanceSquared int64
	{
		const maxint int64 = int64((^uint64(0)) >> 1)

		favoredIndex = defaultIndex
		favoredDistanceSquared = maxint
	}

	{
		length := receiver.Len()

		for index:=0; index<length; index++ {

			candidateR, candidateG, candidateB, candidateA := receiver.Color(uint8(index)).RGBA()

			if r == candidateR && g == candidateG && b == candidateB && a == candidateA {
				return uint8(index)
			}

			var distanceSquared int64 = 0
			{
				difference := int64(r) - int64(candidateR)
				distanceSquared += difference*difference
			}
			{
				difference := int64(g) - int64(candidateG)
				distanceSquared += difference*difference
			}
			{
				difference := int64(b) - int64(candidateB)
				distanceSquared += difference*difference
			}
			{
				difference := int64(a) - int64(candidateA)
				distanceSquared += difference*difference
			}

			if distanceSquared < favoredDistanceSquared {
				favoredIndex = uint8(index)
				favoredDistanceSquared = distanceSquared
			}
		}
	}

	return favoredIndex
}
