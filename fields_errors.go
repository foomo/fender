package fender

import "bytes"

// FieldsErrors type
type FieldsErrors struct {
	Errors map[string][]error
}

// NewFieldsErrors constructor
func NewFieldsErrors(e map[string][]error) *FieldsErrors {
	return &FieldsErrors{
		Errors: e,
	}
}

// String interface
func (e FieldsErrors) Error() string {
	var i int
	buff := bytes.NewBufferString("")
	for name, errs := range e.Errors {
		if i != 0 {
			buff.WriteString(RuleDelimiter)
		}
		buff.WriteString(name + NameDelimiter)
		for j, err := range errs {
			if j != 0 {
				buff.WriteString(ErrorDelimiter)
			}
			buff.WriteString(err.Error())
		}
		i++
	}
	return buff.String()
}
