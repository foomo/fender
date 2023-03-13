package rule

import (
	"context"
)

func UInt8Min(expected uint8) UInt8Rule {
	return func(ctx context.Context, v uint8) error {
		if v < expected {
			return NewErrorf(NameMin, 'd', expected)
		}
		return nil
	}
}

func UInt8Max(expected uint8) UInt8Rule {
	return func(ctx context.Context, v uint8) error {
		if v > expected {
			return NewErrorf(NameMax, 'd', expected)
		}
		return nil
	}
}

func UInt8Required(ctx context.Context, v uint8) error {
	if v == 0 {
		return NewError(NameRequired)
	}
	return nil
}

func UInt8NotRequired(ctx context.Context, v uint8) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}

func UInt8Size(expected uint8) UInt8Rule {
	return func(ctx context.Context, v uint8) error {
		if v != expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

func UInt8OmitEmpty(ctx context.Context, v uint8) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}
