package fender

import "github.com/foomo/fender/fend"

func All(fields ...FendField) error {
	errorsMap := make(Errors, len(fields))
	for _, field := range fields {
		if _, ok := errorsMap[field.name]; !ok {
			if err := fend.First(field.fends...); err != nil {
				errorsMap[field.name] = err
			}
		}
	}
	if len(errorsMap) > 0 {
		return NewError(errorsMap)
	}
	return nil
}
