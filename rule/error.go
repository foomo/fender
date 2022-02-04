package rule

import (
	"strings"

	"github.com/foomo/fender/config"
	"github.com/pkg/errors"
)

type Error struct {
	Rule     string
	Meta     string
	CauseTxt string
}

var Err = errors.New("rule violation")

func IsError(err error) bool {
	if v, ok := err.(*Error); ok && v == nil {
		return false
	} else if err == nil {
		return false
	}
	return true
}

// NewError constructor
func NewError(cause error, rule string, meta ...string) *Error {
	return &Error{
		Rule:     rule,
		Meta:     strings.Join(meta, config.MetaDelimiter),
		CauseTxt: cause.Error(),
	}
}

// Is interface
func (e *Error) Is(err error) bool {
	return e != nil && err != nil && err.Error() == e.Error()
}

// Unwrap interface
func (e *Error) Unwrap() error {
	return errors.New(e.CauseTxt)
}

// Cause interface
func (e *Error) Cause() error {
	return errors.New(e.CauseTxt)
}

// Error interface
func (e *Error) Error() string {
	s := e.Rule
	if e.Meta != "" {
		s += config.MetaDelimiter + e.Meta
	}
	return s
}
