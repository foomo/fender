package rule

import (
	"context"
)

func IntMin(expected int) IntRule {
	return func(ctx context.Context, v int) error {
		if v < expected {
			return NewErrorf(NameMin, 'd', expected)
		}
		return nil
	}
}

func IntMax(expected int) IntRule {
	return func(ctx context.Context, v int) error {
		if v > expected {
			return NewErrorf(NameMax, 'd', expected)
		}
		return nil
	}
}

func IntRequired(ctx context.Context, v int) error {
	if v == 0 {
		return NewError(NameRequired)
	}
	return nil
}

func IntNotRequired(ctx context.Context, v int) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}

func IntSize(expected int) IntRule {
	return func(ctx context.Context, v int) error {
		if v != expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

func IntOmitEmpty(ctx context.Context, v int) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}
