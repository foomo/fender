package rule

import (
	"errors"
	"fmt"
)

var (
	ErrBreak     = errors.New("break")
	ErrUnhandled = errors.New("unhandled type")
)

func NewMinError(verb rune, v interface{}) *Error {
	return NewError(NameMin, fmt.Sprintf("%"+string(verb), v))
}
