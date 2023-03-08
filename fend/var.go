package fend

import (
	"context"

	"github.com/foomo/fender/rule"
	"go.uber.org/multierr"
)

type VarError struct {
	Errors []error
}

func (e *VarError) Error() string {
	return ""
}

func Var[T any](value T, rules ...rule.Rule[T]) Fend {
	return func(ctx context.Context, mode Mode) error {
		var cause error
		for _, r := range rules {
			err := r(ctx, value)
			if e, ok := err.(*rule.Error); ok { //nolint:errorlint
				cause = multierr.Append(cause, e)
				// break if we only want the first error
				if mode == ModeFirst {
					break
				}
			} else if err != nil {
				return err
			}
		}
		if cause != nil {
			return NewError("", cause)
		}
		return nil
	}
}
