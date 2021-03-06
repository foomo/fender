package rule

import (
	"fmt"
	"unicode/utf8"
)

const NameSize = "size"

var ErrSize = &Error{Rule: NameSize}

// NewSizeError constructor
func NewSizeError(verb rune, v interface{}) *Error {
	return NewError(NameSize, fmt.Sprintf("%"+string(verb), v))
}

func SizeInt(expected int) IntRule {
	return func(v int) Rule {
		return func() (*Error, error) {
			if v != expected {
				return NewSizeError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func SizeInt8(expected int8) Int8Rule {
	return func(v int8) Rule {
		return func() (*Error, error) {
			if v != expected {
				return NewSizeError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func SizeInt32(expected int32) Int32Rule {
	return func(v int32) Rule {
		return func() (*Error, error) {
			if v != expected {
				return NewSizeError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func SizeInt64(expected int64) Int64Rule {
	return func(v int64) Rule {
		return func() (*Error, error) {
			if v != expected {
				return NewSizeError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func SizeUInt(expected uint) UIntRule {
	return func(v uint) Rule {
		return func() (*Error, error) {
			if v != expected {
				return NewSizeError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func SizeUInt8(expected uint8) UInt8Rule {
	return func(v uint8) Rule {
		return func() (*Error, error) {
			if v != expected {
				return NewSizeError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func SizeUInt32(expected uint32) UInt32Rule {
	return func(v uint32) Rule {
		return func() (*Error, error) {
			if v != expected {
				return NewSizeError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func SizeUInt64(expected uint64) UInt64Rule {
	return func(v uint64) Rule {
		return func() (*Error, error) {
			if v != expected {
				return NewSizeError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func SizeFloat32(expected float32) Float32Rule {
	return func(v float32) Rule {
		return func() (*Error, error) {
			if v != expected {
				return NewSizeError('g', expected), nil
			}
			return nil, nil
		}
	}
}

func SizeFloat64(expected float64) Float64Rule {
	return func(v float64) Rule {
		return func() (*Error, error) {
			if v != expected {
				return NewSizeError('g', expected), nil
			}
			return nil, nil
		}
	}
}

func SizeString(expected int) StringRule {
	return func(v string) Rule {
		return func() (*Error, error) {
			if utf8.RuneCountInString(v) != expected {
				return NewSizeError('d', expected), nil
			}
			return nil, nil
		}
	}
}
