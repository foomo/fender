package fend

import (
	"errors"
	"strings"

	"github.com/foomo/fender/config"
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

func (e *Error) Name() string {
	return e.name
}

func (e *Error) Error() string {
	errs := multierr.Errors(e.cause)
	causes := make([]string, len(errs))
	for i, cause := range errs {
		causes[i] = cause.Error()
	}
	return e.name + config.DelimiterFendName + strings.Join(causes, config.DelimiterRule)
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
