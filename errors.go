package fender

import (
	"bytes"

	"github.com/foomo/fender/config"
	"github.com/foomo/fender/rule"
)

// FieldErrors type
type FieldErrors map[string]*rule.Error

// String returns the string representation
func (m FieldErrors) String() string {
	var i int
	buff := bytes.NewBufferString("")
	for name, err := range m {
		if i != 0 {
			buff.WriteString(config.RuleDelimiter)
		}
		buff.WriteString(name + config.FieldDelimiter)
		buff.WriteString(err.Error())
		i++
	}
	return buff.String()
}
