package rule

import (
	"errors"
	"strings"

	"github.com/foomo/fender/config"
)

type Error struct {
	cause error
	rule  string
	meta  string
}

var Err = errors.New("rule violation")

// NewError constructor
func NewError(cause error, rule string, meta ...string) *Error {
	return &Error{
		cause: cause,
		rule:  rule,
		meta:  strings.Join(meta, config.MetaDelimiter),
	}
}

func (e *Error) Is(err error) bool {
	return errors.Is(err, Err)
}

// Unwrap interface
func (e *Error) Unwrap() error {
	return e.cause
}

// Cause interface
func (e *Error) Cause() error {
	return e.cause
}

// Error interface
func (e *Error) Error() string {
	s := e.rule
	if e.meta != "" {
		s += config.MetaDelimiter + e.meta
	}
	return s
}
