package fender

import (
	"github.com/pkg/errors"
)

type wrappedError struct {
	err   error
	cause *Error
}

func NewWrappedError(err error, cause *Error) error {
	return &wrappedError{
		err:   err,
		cause: cause,
	}
}

func (e *wrappedError) Is(target error) bool {
	return errors.Is(e.err, target) || errors.Is(e.cause, target)
}

func (e *wrappedError) Cause() error {
	return e.cause
}

func (e *wrappedError) Unwrap() error {
	return e.err
}

func (e *wrappedError) Error() string {
	return e.err.Error() + ": " + e.cause.Error()
}

func WrapError(err error, cause *Error) error {
	return NewWrappedError(err, cause)
}

func UnwrapError(err error) *Error {
	c := errors.Cause(err)
	if v, ok := c.(*Error); ok {
		return v
	}
	return nil
}
