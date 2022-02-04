package fender

import (
	"github.com/foomo/fender/fend"
)

func All(fields ...FendField) (*Error, error) {
	errorsMap := make(FieldErrors, len(fields))
	for _, field := range fields {
		if _, ok := errorsMap[field.name]; !ok {
			if fendErr, err := fend.First(field.fends...); err != nil {
				return nil, err
			} else if fendErr != nil {
				errorsMap[field.name] = fendErr
			}
		}
	}
	if len(errorsMap) > 0 {
		return NewError(errorsMap), nil
	}
	return nil, nil
}
