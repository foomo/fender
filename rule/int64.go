package rule

import (
	"context"
)

func Int64Min(expected int64) Int64Rule {
	return func(ctx context.Context, v int64) error {
		if v < expected {
			return NewErrorf(NameMin, 'd', expected)
		}
		return nil
	}
}

func Int64Max(expected int64) Int64Rule {
	return func(ctx context.Context, v int64) error {
		if v > expected {
			return NewErrorf(NameMax, 'd', expected)
		}
		return nil
	}
}

func Int64Required(ctx context.Context, v int64) error {
	if v == 0 {
		return NewError(NameRequired)
	}
	return nil
}

func Int64Size(expected int64) Int64Rule {
	return func(ctx context.Context, v int64) error {
		if v != expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

func Int64OmitEmpty(ctx context.Context, v int64) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}
