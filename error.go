package fender

import (
	"errors"
)

// Error type
type Error struct {
	fields Errors
}

var Err = errors.New("validation error")

// NewError constructor
func NewError(fields Errors) *Error {
	return &Error{
		fields: fields,
	}
}

// Is interface
func (e *Error) Is(err error) bool {
	return errors.Is(err, Err)
}

// Error interface
func (e *Error) Error() string {
	return e.fields.String()
}

func (e *Error) Errors() Errors {
	return e.fields
}

func (e *Error) First() error {
	for _, errs := range e.fields {
		if len(errs) > 0 {
			return errs[0]
		}
	}
	return nil
}
