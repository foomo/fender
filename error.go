package fender

import (
	"errors"
	"strings"

	"github.com/foomo/fender/config"
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

func (e *Error) Error() string {
	errs := multierr.Errors(e.cause)
	causes := make([]string, len(errs))
	for i, cause := range errs {
		causes[i] = cause.Error()
	}
	return strings.Join(causes, config.FendDelimiter)
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
