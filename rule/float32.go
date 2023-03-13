package rule

import (
	"context"
)

func Float32Min(expected float32) Float32Rule {
	return func(ctx context.Context, v float32) error {
		if v < expected {
			return NewErrorf(NameMin, 'd', expected)
		}
		return nil
	}
}

func Float32Max(expected float32) Float32Rule {
	return func(ctx context.Context, v float32) error {
		if v > expected {
			return NewErrorf(NameMax, 'd', expected)
		}
		return nil
	}
}

func Float32Required(ctx context.Context, v float32) error {
	if v == 0 {
		return NewError(NameRequired)
	}
	return nil
}

func Float32NotRequired(ctx context.Context, v float32) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}

func Float32Size(expected float32) Float32Rule {
	return func(ctx context.Context, v float32) error {
		if v != expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

func Float32OmitEmpty(ctx context.Context, v float32) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}
