package rule

import (
	"strings"

	"github.com/foomo/fender/config"
)

type Error struct {
	Rule string
	Meta string
}

var Err = &Error{Rule: ""}

// NewError constructor
func NewError(rule string, meta ...string) *Error {
	return &Error{
		Rule: rule,
		Meta: strings.Join(meta, config.MetaDelimiter),
	}
}

// Is interface
func (e *Error) Is(err error) bool {
	if err == nil {
		return false
	}
	if v, ok := err.(*Error); ok && (v.Rule == e.Rule || v.Rule == "") {
		return true
	}
	return false
}

// Error interface
func (e *Error) Error() string {
	s := e.Rule
	if e.Meta != "" && strings.Contains(e.Meta, config.MetaDelimiter) {
		s = e.Meta
	} else if e.Meta != "" {
		s += config.MetaDelimiter + e.Meta
	}
	return s
}
