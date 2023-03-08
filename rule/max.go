package rule

import (
	"context"
	"fmt"
	"unicode/utf8"
)

const NameMax Name = "max"

// NewMaxError constructor
func NewMaxError(verb rune, v interface{}) *Error {
	return NewError(NameMax, fmt.Sprintf("%"+string(verb), v))
}

func MaxInt(expected int) IntRule {
	return func(ctx context.Context, v int) error {
		if v > expected {
			return NewMaxError('d', expected)
		}
		return nil
	}
}

func MaxInt8(expected int8) Int8Rule {
	return func(ctx context.Context, v int8) error {
		if v > expected {
			return NewMaxError('d', expected)
		}
		return nil
	}
}

func MaxInt32(expected int32) Int32Rule {
	return func(ctx context.Context, v int32) error {
		if v > expected {
			return NewMaxError('d', expected)
		}
		return nil
	}
}

func MaxInt64(expected int64) Int64Rule {
	return func(ctx context.Context, v int64) error {
		if v > expected {
			return NewMaxError('d', expected)
		}
		return nil
	}
}

func MaxUInt(expected uint) UIntRule {
	return func(ctx context.Context, v uint) error {
		if v > expected {
			return NewMaxError('d', expected)
		}
		return nil
	}
}

func MaxUInt8(expected uint8) UInt8Rule {
	return func(ctx context.Context, v uint8) error {
		if v > expected {
			return NewMaxError('d', expected)
		}
		return nil
	}
}

func MaxUInt32(expected uint32) UInt32Rule {
	return func(ctx context.Context, v uint32) error {
		if v > expected {
			return NewMaxError('d', expected)
		}
		return nil
	}
}

func MaxUInt64(expected uint64) UInt64Rule {
	return func(ctx context.Context, v uint64) error {
		if v > expected {
			return NewMaxError('d', expected)
		}
		return nil
	}
}

func MaxFloat32(expected float32) Float32Rule {
	return func(ctx context.Context, v float32) error {
		if v > expected {
			return NewMaxError('g', expected)
		}
		return nil
	}
}

func MaxFloat64(expected float64) Float64Rule {
	return func(ctx context.Context, v float64) error {
		if v > expected {
			return NewMaxError('g', expected)
		}
		return nil
	}
}

func MaxString(expected int) StringRule {
	return func(ctx context.Context, v string) error {
		if utf8.RuneCountInString(v) > expected {
			return NewMaxError('d', expected)
		}
		return nil
	}
}
