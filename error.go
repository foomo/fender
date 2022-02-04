package fender

import (
	"errors"
)

// Error type
type Error struct {
	fields FieldErrors
}

var Err = errors.New("validation error")

func IsError(err error) bool {
	if v, ok := err.(*Error); ok && v == nil {
		return false
	} else if err == nil {
		return false
	}
	return true
}

// NewError constructor
func NewError(fields FieldErrors) *Error {
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

func (e *Error) Errors() FieldErrors {
	return e.fields
}

func (e *Error) First() error {
	for _, err := range e.fields {
		return err
	}
	return nil
}
