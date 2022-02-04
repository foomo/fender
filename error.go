package fender

// Error type
type Error struct {
	Fields FieldErrors
}

var Err = &Error{}

// NewError constructor
func NewError(fields FieldErrors) *Error {
	return &Error{
		Fields: fields,
	}
}

// Is interface
func (e *Error) Is(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*Error)
	return ok
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
