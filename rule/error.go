package rule

import (
	"errors"
	"fmt"
	"strings"

	"github.com/foomo/fender/config"
)

type Error struct {
	Rule Name
	Meta string
}

// NewError constructor
func NewError(rule Name, meta ...string) *Error {
	var m []string
	for _, v := range meta {
		if strings.TrimSpace(v) != "" {
			m = append(m, v)
		}
	}
	return &Error{
		Rule: rule,
		Meta: strings.Join(m, config.DelimiterRuleMeta),
	}
}

func NewErrorf(name Name, verb rune, v ...any) *Error {
	meta := make([]string, len(v))
	for i, a := range v {
		meta[i] = fmt.Sprintf("%"+string(verb), a)
	}
	return NewError(name, meta...)
}

// Error interface
func (e *Error) Error() string {
	s := e.Rule.String()
	if e.Meta != "" && strings.Contains(e.Meta, config.DelimiterRuleMeta) {
		s = e.Meta
	} else if e.Meta != "" {
		s += config.DelimiterRuleMeta + e.Meta
	}
	return s
}

func AsError(err error) *Error {
	var fendErr *Error
	if errors.As(err, &fendErr) {
		return fendErr
	}
	return nil
}
