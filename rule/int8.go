package rule

import (
	"context"
)

func Int8Min(expected int8) Int8Rule {
	return func(ctx context.Context, v int8) error {
		if v < expected {
			return NewErrorf(NameMin, 'd', expected)
		}
		return nil
	}
}

func Int8Max(expected int8) Int8Rule {
	return func(ctx context.Context, v int8) error {
		if v > expected {
			return NewErrorf(NameMax, 'd', expected)
		}
		return nil
	}
}

func Int8Required(ctx context.Context, v int8) error {
	if v == 0 {
		return NewError(NameRequired)
	}
	return nil
}

func Int8NotRequired(ctx context.Context, v int8) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}

func Int8Size(expected int8) Int8Rule {
	return func(ctx context.Context, v int8) error {
		if v != expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

func Int8OmitEmpty(ctx context.Context, v int8) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}
