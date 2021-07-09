package fender

import (
	"bytes"

	"github.com/foomo/fender/config"
)

// Errors type
type Errors map[string][]error

// String returns the string representation
func (m Errors) String() string {
	var i int
	buff := bytes.NewBufferString("")
	for name, errs := range m {
		if i != 0 {
			buff.WriteString(config.RuleDelimiter)
		}
		buff.WriteString(name + config.FieldDelimiter)
		for j, err := range errs {
			if j != 0 {
				buff.WriteString(config.ErrorDelimiter)
			}
			buff.WriteString(err.Error())
		}
		i++
	}
	return buff.String()
}
