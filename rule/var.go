package rule

import (
	"github.com/go-playground/validator/v10"

	"github.com/foomo/fender/config"
)

const NameVar = "var"

var ErrVar = &Error{Rule: NameVar}

// NewVarError constructor
func NewVarError(tag, meta string) *Error {
	return NewError(NameVar, tag, meta)
}

func Var(tag string) InterfaceRule {
	return func(v interface{}) Rule {
		return func() (*Error, error) {
			if err := config.Validator.Var(v, tag); err != nil {
				if validationErrors, ok := err.(validator.ValidationErrors); ok { //nolint:errorlint
					for _, err := range validationErrors {
						return NewVarError(err.Tag(), err.Param()), nil
					}
				}
			}
			return nil, nil
		}
	}
}
