package fender

import (
	"github.com/foomo/fender/config"
	"github.com/foomo/fender/rule"
	"github.com/go-playground/validator/v10"
)

func Struct(s interface{}) error {
	if err := config.Validator.Struct(s); err == nil {
		return nil
	} else if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	} else if errs, ok := err.(validator.ValidationErrors); ok {
		errorsMap := make(Errors, len(errs))
		for _, err := range errs {
			errorsMap[err.Field()] = rule.NewCustomRuleError(err.Tag(), err.Param())
		}
		return NewError(errorsMap)
	} else {
		return err
	}
}
