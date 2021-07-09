package rule

import (
	"errors"
	"fmt"

	"github.com/foomo/fender/config"
	"github.com/go-playground/validator/v10"
)

const NameVar Name = "var"

var ErrVar = errors.New(NameVar.String())

// NewVarError constructor
func NewVarError(tag, meta string) *Error {
	return NewError(ErrVar, tag, meta)
}

func Var(tag string) InterfaceRule {
	return func(v interface{}) Rule {
		return func() *Error {
			if err := config.Validator.Var(v, tag); err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					return NewVarError(err.Tag(), err.Param())
				}
			}
			return nil
		}
	}
}
