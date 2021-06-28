package fender

import "bytes"

// FieldsErrorMap type
type FieldsErrorMap map[string]*RuleError

func (m FieldsErrorMap) String() string {
	var i int
	buff := bytes.NewBufferString("")
	for name, err := range m {
		if i != 0 {
			buff.WriteString(RuleDelimiter)
		}
		buff.WriteString(name + FieldDelimiter)
		buff.WriteString(err.Error())
		i++
	}
	return buff.String()
}

// Error returns a std error if not empty
func (m FieldsErrorMap) Error() error {
	if len(m) > 0 {
		return NewFieldsError(m)
	}
	return nil
}
