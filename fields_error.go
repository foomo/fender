package fender

import "bytes"

// FieldsError type
type FieldsError struct {
	Errors map[string]error
}

// NewFieldsError constructor
func NewFieldsError(e map[string]error) *FieldsError {
	return &FieldsError{
		Errors: e,
	}
}

// Error interface
func (e FieldsError) Error() string {
	var i int
	buff := bytes.NewBufferString("")
	for name, err := range e.Errors {
		if i != 0 {
			buff.WriteString(RuleDelimiter)
		}
		buff.WriteString(name + NameDelimiter)
		buff.WriteString(err.Error())
		i++
	}
	return buff.String()
}
