package fender

import (
	"github.com/foomo/fender/fend"
)

func First(fields ...FendField) error {
	for _, field := range fields {
		if err := fend.First(field.fends...); err != nil {
			return NewError(Errors{field.name: err})
		}
	}
	return nil
}
