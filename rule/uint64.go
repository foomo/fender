package rule

import (
	"context"
)

func UInt64Min(expected uint64) UInt64Rule {
	return func(ctx context.Context, v uint64) error {
		if v < expected {
			return NewErrorf(NameMin, 'd', expected)
		}
		return nil
	}
}

func UInt64Max(expected uint64) UInt64Rule {
	return func(ctx context.Context, v uint64) error {
		if v > expected {
			return NewErrorf(NameMax, 'd', expected)
		}
		return nil
	}
}

func UInt64Required(ctx context.Context, v uint64) error {
	if v == 0 {
		return NewError(NameRequired)
	}
	return nil
}

func UInt64NotRequired(ctx context.Context, v uint64) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}

func UInt64Size(expected uint64) UInt64Rule {
	return func(ctx context.Context, v uint64) error {
		if v != expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

func UInt64OmitEmpty(ctx context.Context, v uint64) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}
