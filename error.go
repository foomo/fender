package fender

import (
	"errors"
)

// Error type
type Error struct {
	Fields FieldErrors
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
		Fields: fields,
	}
}

// Is interface
func (e *Error) Is(err error) bool {
	return e != nil && err != nil && err.Error() == e.Error()
}

// Error interface
func (e *Error) Error() string {
	return e.Fields.String()
}

func (e *Error) Errors() FieldErrors {
	return e.Fields
}

func (e *Error) First() error {
	for _, err := range e.Fields {
		return err
	}
	return nil
}

func (e *Error) Wrap(err error) error {
	return NewWrappedError(err, e)
}
