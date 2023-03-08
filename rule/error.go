package rule

import (
	"errors"
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
		Meta: strings.Join(m, config.RuleMetaDelimiter),
	}
}

// Error interface
func (e *Error) Error() string {
	s := e.Rule.String()
	if e.Meta != "" && strings.Contains(e.Meta, config.RuleMetaDelimiter) {
		s = e.Meta
	} else if e.Meta != "" {
		s += config.RuleMetaDelimiter + e.Meta
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
