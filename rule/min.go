package rule

import (
	"context"
	"fmt"
	"unicode/utf8"
)

const NameMin Name = "min"

// NewMinError constructor
func NewMinError(verb rune, v interface{}) *Error {
	return NewError(NameMin, fmt.Sprintf("%"+string(verb), v))
}

func MinInt(expected int) IntRule {
	return func(ctx context.Context, v int) error {
		if v < expected {
			return NewMinError('d', expected)
		}
		return nil
	}
}

func MinInt8(expected int8) Int8Rule {
	return func(ctx context.Context, v int8) error {
		if v < expected {
			return NewMinError('d', expected)
		}
		return nil
	}
}

func MinInt32(expected int32) Int32Rule {
	return func(ctx context.Context, v int32) error {
		if v < expected {
			return NewMinError('d', expected)
		}
		return nil
	}
}

func MinInt64(expected int64) Int64Rule {
	return func(ctx context.Context, v int64) error {
		if v < expected {
			return NewMinError('d', expected)
		}
		return nil
	}
}

func MinUInt(expected uint) UIntRule {
	return func(ctx context.Context, v uint) error {
		if v < expected {
			return NewMinError('d', expected)
		}
		return nil
	}
}

func MinUInt8(expected uint8) UInt8Rule {
	return func(ctx context.Context, v uint8) error {
		if v < expected {
			return NewMinError('d', expected)
		}
		return nil
	}
}

func MinUInt32(expected uint32) UInt32Rule {
	return func(ctx context.Context, v uint32) error {
		if v < expected {
			return NewMinError('d', expected)
		}
		return nil
	}
}

func MinUInt64(expected uint64) UInt64Rule {
	return func(ctx context.Context, v uint64) error {
		if v < expected {
			return NewMinError('d', expected)
		}
		return nil
	}
}

func MinFloat32(expected float32) Float32Rule {
	return func(ctx context.Context, v float32) error {
		if v < expected {
			return NewMinError('g', expected)
		}
		return nil
	}
}

func MinFloat64(expected float64) Float64Rule {
	return func(ctx context.Context, v float64) error {
		if v < expected {
			return NewMinError('g', expected)
		}
		return nil
	}
}

func MinString(expected int) func(ctx context.Context, v string) error {
	return func(ctx context.Context, v string) error {
		if utf8.RuneCountInString(v) < expected {
			return NewMinError('d', expected)
		}
		return nil
	}
}
