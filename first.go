package fender

import (
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
)

func First(fields ...FendField) error {
	for _, field := range fields {
		if err := fend.First(field.fends...); err != nil {
			return NewError(Errors{field.name: []*rule.Error{err}})
		}
	}
	return nil
}
