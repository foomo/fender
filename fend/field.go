package fend

import (
	"context"

	"github.com/foomo/fender/rule"
	"go.uber.org/multierr"
)

type FieldError struct {
	Name   string
	Errors []error
}

func (e *FieldError) Error() string {
	return e.Name
}

// Field
// should be func FendField[T any](name string, value T, rules ...rule.Rule[T]) validation.Validator {}
func Field[T any](name string, value T, rules ...func(ctx context.Context, v T) error) Fend {
	return func(ctx context.Context, mode Mode) error {
		var causes error
		for _, r := range rules {
			err := r(ctx, value)
			if e, ok := err.(*rule.Error); ok { //nolint:errorlint
				causes = multierr.Append(causes, e)
				// break if we only want the first error
				if mode == ModeFirst {
					break
				}
			} else if err != nil {
				return err
			}
		}
		if causes != nil {
			return NewError(name, causes)
		}
		return nil
	}
}
