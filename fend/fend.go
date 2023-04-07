package fend

import (
	"context"
	"errors"

	"github.com/foomo/fender/rule"
	"go.uber.org/multierr"
)

type Fend func(ctx context.Context, mode Mode) error

func fend[T any](ctx context.Context, mode Mode, meta string, value T, rules ...rule.Rule[T]) error {
	var causes error
	for _, r := range rules {
		err := r(ctx, value)
		if errors.Is(err, rule.ErrBreak) {
			break
		} else if e, ok := err.(*rule.Error); ok { //nolint:errorlint
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
		return NewError(meta, causes)
	}
	return nil
}

func fendDynamic(ctx context.Context, mode Mode, meta string, rules ...rule.DynamicRule) error {
	var causes error
	for _, r := range rules {
		err := r(ctx)
		if errors.Is(err, rule.ErrBreak) {
			break
		} else if e, ok := err.(*rule.Error); ok { //nolint:errorlint
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
		return NewError(meta, causes)
	}
	return nil
}
