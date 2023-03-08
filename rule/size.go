package rule

import (
	"context"
	"fmt"
	"unicode/utf8"
)

const NameSize Name = "size"

// NewSizeError constructor
func NewSizeError(verb rune, v interface{}) *Error {
	return NewError(NameSize, fmt.Sprintf("%"+string(verb), v))
}

func SizeInt(expected int) IntRule {
	return func(ctx context.Context, v int) error {
		if v != expected {
			return NewSizeError('d', expected)
		}
		return nil
	}
}

func SizeInt8(expected int8) Int8Rule {
	return func(ctx context.Context, v int8) error {
		if v != expected {
			return NewSizeError('d', expected)
		}
		return nil
	}
}

func SizeInt32(expected int32) Int32Rule {
	return func(ctx context.Context, v int32) error {
		if v != expected {
			return NewSizeError('d', expected)
		}
		return nil
	}
}

func SizeInt64(expected int64) Int64Rule {
	return func(ctx context.Context, v int64) error {
		if v != expected {
			return NewSizeError('d', expected)
		}
		return nil
	}
}

func SizeUInt(expected uint) UIntRule {
	return func(ctx context.Context, v uint) error {
		if v != expected {
			return NewSizeError('d', expected)
		}
		return nil
	}
}

func SizeUInt8(expected uint8) UInt8Rule {
	return func(ctx context.Context, v uint8) error {
		if v != expected {
			return NewSizeError('d', expected)
		}
		return nil
	}
}

func SizeUInt32(expected uint32) UInt32Rule {
	return func(ctx context.Context, v uint32) error {
		if v != expected {
			return NewSizeError('d', expected)
		}
		return nil
	}
}

func SizeUInt64(expected uint64) UInt64Rule {
	return func(ctx context.Context, v uint64) error {
		if v != expected {
			return NewSizeError('d', expected)
		}
		return nil
	}
}

func SizeFloat32(expected float32) Float32Rule {
	return func(ctx context.Context, v float32) error {
		if v != expected {
			return NewSizeError('g', expected)
		}
		return nil
	}
}

func SizeFloat64(expected float64) Float64Rule {
	return func(ctx context.Context, v float64) error {
		if v != expected {
			return NewSizeError('g', expected)
		}
		return nil
	}
}

func SizeString(expected int) StringRule {
	return func(ctx context.Context, v string) error {
		if utf8.RuneCountInString(v) != expected {
			return NewSizeError('d', expected)
		}
		return nil
	}
}
