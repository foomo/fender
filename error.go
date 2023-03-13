package fender

import (
	"errors"
	"strings"

	"github.com/foomo/fender/config"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
	"go.uber.org/multierr"
)

type Error struct {
	cause error
}

func NewError(cause error) *Error {
	return &Error{
		cause: cause,
	}
}

func NewFendError(name string, cause error) *Error {
	return NewError(fend.NewError(name, cause))
}

func NewFendRuleError(name string, ruleName rule.Name, ruleMeta ...string) *Error {
	return NewFendError(name, rule.NewError(ruleName, ruleMeta...))
}

func (e *Error) Error() string {
	errs := multierr.Errors(e.cause)
	causes := make([]string, len(errs))
	for i, cause := range errs {
		causes[i] = cause.Error()
	}
	return strings.Join(causes, config.DelimiterFend)
}

func (e *Error) First() error {
	if errs := e.Errors(); len(errs) > 0 {
		return errs[0]
	}
	return nil
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
