package fend

import (
	"errors"
	"strings"

	"github.com/foomo/fender/config"
	"github.com/foomo/fender/rule"
	"go.uber.org/multierr"
)

type Error struct {
	name  string
	cause error
}

func NewError(name string, cause error) *Error {
	return &Error{
		name:  name,
		cause: cause,
	}
}

func NewRuleError(name string, ruleName rule.Name, meta ...string) *Error {
	return &Error{
		name:  name,
		cause: rule.NewError(ruleName, meta...),
	}
}

func (e *Error) Name() string {
	return e.name
}

func (e *Error) Error() string {
	var ret string
	errs := multierr.Errors(e.cause)
	causes := make([]string, len(errs))
	for i, cause := range errs {
		causes[i] = cause.Error()
	}
	if e.name != "" {
		ret += e.name + config.DelimiterFendName
	}
	return ret + strings.Join(causes, config.DelimiterRule)
}

func (e *Error) Errors() []error {
	return multierr.Errors(e.cause)
}

func (e *Error) Unwrap() error {
	return e.cause
}

func AsError(err error) *Error {
	var fendErr *Error
	if errors.As(err, &fendErr) {
		return fendErr
	}
	return nil
}
