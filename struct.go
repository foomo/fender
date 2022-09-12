package fender

import (
	"github.com/go-playground/validator/v10"

	"github.com/foomo/fender/config"
	"github.com/foomo/fender/rule"
)

func Struct(s interface{}) (*Error, error) {
	if err := config.Validator.Struct(s); err == nil {
		return nil, nil
	} else if errs, ok := err.(validator.ValidationErrors); ok { //nolint:errorlint
		errorsMap := make(FieldErrors, len(errs))
		for _, err := range errs {
			errorsMap[err.Field()] = rule.NewCustomRuleError(err.Tag(), err.Param())
		}
		return NewError(errorsMap), nil
	} else {
		return nil, err
	}
}
