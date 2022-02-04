package rule

import (
	"fmt"
	"unicode/utf8"
)

const NameMin = "min"

var ErrMin = &Error{Rule: NameMin}

// NewMinError constructor
func NewMinError(verb rune, v interface{}) *Error {
	return NewError(NameMin, fmt.Sprintf("%"+string(verb), v))
}

func MinInt(expected int) IntRule {
	return func(v int) Rule {
		return func() (*Error, error) {
			if v < expected {
				return NewMinError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MinInt8(expected int8) Int8Rule {
	return func(v int8) Rule {
		return func() (*Error, error) {
			if v < expected {
				return NewMinError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MinInt32(expected int32) Int32Rule {
	return func(v int32) Rule {
		return func() (*Error, error) {
			if v < expected {
				return NewMinError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MinInt64(expected int64) Int64Rule {
	return func(v int64) Rule {
		return func() (*Error, error) {
			if v < expected {
				return NewMinError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MinUInt(expected uint) UIntRule {
	return func(v uint) Rule {
		return func() (*Error, error) {
			if v < expected {
				return NewMinError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MinUInt8(expected uint8) UInt8Rule {
	return func(v uint8) Rule {
		return func() (*Error, error) {
			if v < expected {
				return NewMinError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MinUInt32(expected uint32) UInt32Rule {
	return func(v uint32) Rule {
		return func() (*Error, error) {
			if v < expected {
				return NewMinError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MinUInt64(expected uint64) UInt64Rule {
	return func(v uint64) Rule {
		return func() (*Error, error) {
			if v < expected {
				return NewMinError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MinFloat32(expected float32) Float32Rule {
	return func(v float32) Rule {
		return func() (*Error, error) {
			if v < expected {
				return NewMinError('g', expected), nil
			}
			return nil, nil
		}
	}
}

func MinFloat64(expected float64) Float64Rule {
	return func(v float64) Rule {
		return func() (*Error, error) {
			if v < expected {
				return NewMinError('g', expected), nil
			}
			return nil, nil
		}
	}
}

func MinString(expected int) StringRule {
	return func(v string) Rule {
		return func() (*Error, error) {
			if utf8.RuneCountInString(v) < expected {
				return NewMinError('d', expected), nil
			}
			return nil, nil
		}
	}
}
