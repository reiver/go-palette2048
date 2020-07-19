package palette2048

import (
	"github.com/reiver/go-rgba32"
)

func (receiver Slice) Len() int {
	return ByteSize / rgba32.ByteSize
}
