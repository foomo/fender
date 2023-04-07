package rule

import (
	"errors"
)

var (
	ErrBreak     = errors.New("break")
	ErrUnhandled = errors.New("unhandled type")
)
