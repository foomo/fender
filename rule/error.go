package rule

import (
	"errors"
	"fmt"
	"strings"

	"github.com/foomo/fender/config"
)

type Error struct {
	Rule Name
	Meta []string
}

// NewError constructor
func NewError(rule Name, meta ...string) *Error {
	return &Error{
		Rule: rule,
		Meta: meta,
	}
}

// Error interface
func (e *Error) Error() string {
	s := e.Rule.String()
	if len(e.Meta) > 0 {
		s += config.DelimiterRuleMeta + e.Meta[0]
	}
	if len(e.Meta) > 1 {
		s = fmt.Sprintf("%s[%s]", s, strings.Join(e.Meta[1:], ","))
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

func Meta(verb rune, value any) string {
	return fmt.Sprintf("%"+string(verb), value)
}
