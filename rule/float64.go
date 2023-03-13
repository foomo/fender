package rule

import (
	"context"
)

func Float64Min(expected float64) Float64Rule {
	return func(ctx context.Context, v float64) error {
		if v < expected {
			return NewErrorf(NameMin, 'd', expected)
		}
		return nil
	}
}

func Float64Max(expected float64) Float64Rule {
	return func(ctx context.Context, v float64) error {
		if v > expected {
			return NewErrorf(NameMax, 'd', expected)
		}
		return nil
	}
}

func Float64Required(ctx context.Context, v float64) error {
	if v == 0 {
		return NewError(NameRequired)
	}
	return nil
}

func Float64NotRequired(ctx context.Context, v float64) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}

func Float64Size(expected float64) Float64Rule {
	return func(ctx context.Context, v float64) error {
		if v != expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

func Float64OmitEmpty(ctx context.Context, v float64) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}
