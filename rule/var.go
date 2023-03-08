package rule

import (
	"context"

	"github.com/go-playground/validator/v10"

	"github.com/foomo/fender/config"
)

const NameVar Name = "var"

// NewVarError constructor
func NewVarError(tag, meta string) *Error {
	return NewError(NameVar, tag, meta)
}

func Var(tag string) InterfaceRule {
	return func(ctx context.Context, v any) error {
		if err := config.Validator.VarCtx(ctx, v, tag); err != nil {
			if validationErrors, ok := err.(validator.ValidationErrors); ok { //nolint:errorlint
				for _, err := range validationErrors {
					return NewVarError(err.Tag(), err.Param())
				}
			}
		}
		return nil
	}
}
