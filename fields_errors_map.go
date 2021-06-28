package fender

import "bytes"

// FieldsErrorsMap type
type FieldsErrorsMap map[string][]*RuleError

// String interface
func (m FieldsErrorsMap) String() string {
	var i int
	buff := bytes.NewBufferString("")
	for name, errs := range m {
		if i != 0 {
			buff.WriteString(RuleDelimiter)
		}
		buff.WriteString(name + FieldDelimiter)
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

// Error returns a std error if not empty
func (m FieldsErrorsMap) Error() error {
	if len(m) > 0 {
		return NewFieldsErrors(m)
	}
	return nil
}
