package fender

import (
	"github.com/foomo/fender/fend"
)

func First(fields ...FendField) (*Error, error) {
	for _, field := range fields {
		if ruleErr, err := fend.First(field.fends...); err != nil {
			return nil, err
		} else if ruleErr != nil {
			return NewError(FieldErrors{field.name: ruleErr}), nil
		}
	}
	return nil, nil
}
