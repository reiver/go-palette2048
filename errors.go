package palette2048

import (
	"errors"
)

var (
	errInternalError = errors.New("internal error")
	errNilReceiver   = errors.New("nil receiver")
)
