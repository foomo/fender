package rule

import (
	"context"
)

func Int32Min(expected int32) Int32Rule {
	return func(ctx context.Context, v int32) error {
		if v < expected {
			return NewErrorf(NameMin, 'd', expected)
		}
		return nil
	}
}

func Int32Max(expected int32) Int32Rule {
	return func(ctx context.Context, v int32) error {
		if v > expected {
			return NewErrorf(NameMax, 'd', expected)
		}
		return nil
	}
}

func Int32Required(ctx context.Context, v int32) error {
	if v == 0 {
		return NewError(NameRequired)
	}
	return nil
}

func Int32Size(expected int32) Int32Rule {
	return func(ctx context.Context, v int32) error {
		if v != expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

func Int32OmitEmpty(ctx context.Context, v int32) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}
