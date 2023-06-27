package fender

import (
	"context"

	"github.com/foomo/fender/fend"
)

func All(ctx context.Context, fends ...fend.Fend) error {
	return Mode(ctx, fend.ModeAll, fends...)
}

func First(ctx context.Context, fends ...fend.Fend) error {
	return Mode(ctx, fend.ModeFirst, fends...)
}

func AllFirst(ctx context.Context, fends ...fend.Fend) error {
	return Mode(ctx, fend.ModeAllFirst, fends...)
}

func Mode(ctx context.Context, mode fend.Mode, fends ...fend.Fend) error {
	var cause []*fend.Error
	for _, validator := range fends {
		err := validator(ctx, mode)
		if e, ok := err.(*fend.Error); ok { //nolint:errorlint
			// append error
			cause = append(cause, e)
			// break if we only want the first errors
			if mode == fend.ModeFirst || mode == fend.ModeAllFirst {
				break
			}
		} else if err != nil {
			return err
		}
	}
	if cause != nil {
		return NewError(cause...)
	}
	return nil
}
