package rule

import (
	"context"
)

func UInt32Min(expected uint32) UInt32Rule {
	return func(ctx context.Context, v uint32) error {
		if v < expected {
			return NewErrorf(NameMin, 'd', expected)
		}
		return nil
	}
}

func UInt32Max(expected uint32) UInt32Rule {
	return func(ctx context.Context, v uint32) error {
		if v > expected {
			return NewErrorf(NameMax, 'd', expected)
		}
		return nil
	}
}

func UInt32Required(ctx context.Context, v uint32) error {
	if v == 0 {
		return NewError(NameRequired)
	}
	return nil
}

func UInt32Size(expected uint32) UInt32Rule {
	return func(ctx context.Context, v uint32) error {
		if v != expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

func UInt32OmitEmpty(ctx context.Context, v uint32) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}
