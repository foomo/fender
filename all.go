package fender

import "github.com/foomo/fender/fend"

func All(fields ...FendField) error {
	errorsMap := Errors{}
	for _, field := range fields {
		if errs := fend.All(field.fends...); errs != nil {
			errorsMap[field.name] = append(errorsMap[field.name], errs...)
		}
	}
	if len(errorsMap) > 0 {
		return NewError(errorsMap)
	}
	return nil
}
