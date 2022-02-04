package rule

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

const NameMax Name = "max"

var ErrMax = errors.New(NameMax.String())

// NewMaxError constructor
func NewMaxError(verb rune, v interface{}) *Error {
	return NewError(ErrMax, NameMax.String(), fmt.Sprintf("%"+string(verb), v))
}

func MaxInt(expected int) IntRule {
	return func(v int) Rule {
		return func() (*Error, error) {
			if v > expected {
				return NewMaxError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MaxInt8(expected int8) Int8Rule {
	return func(v int8) Rule {
		return func() (*Error, error) {
			if v > expected {
				return NewMaxError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MaxInt32(expected int32) Int32Rule {
	return func(v int32) Rule {
		return func() (*Error, error) {
			if v > expected {
				return NewMaxError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MaxInt64(expected int64) Int64Rule {
	return func(v int64) Rule {
		return func() (*Error, error) {
			if v > expected {
				return NewMaxError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MaxUInt(expected uint) UIntRule {
	return func(v uint) Rule {
		return func() (*Error, error) {
			if v > expected {
				return NewMaxError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MaxUInt8(expected uint8) UInt8Rule {
	return func(v uint8) Rule {
		return func() (*Error, error) {
			if v > expected {
				return NewMaxError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MaxUInt32(expected uint32) UInt32Rule {
	return func(v uint32) Rule {
		return func() (*Error, error) {
			if v > expected {
				return NewMaxError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MaxUInt64(expected uint64) UInt64Rule {
	return func(v uint64) Rule {
		return func() (*Error, error) {
			if v > expected {
				return NewMaxError('d', expected), nil
			}
			return nil, nil
		}
	}
}

func MaxFloat32(expected float32) Float32Rule {
	return func(v float32) Rule {
		return func() (*Error, error) {
			if v > expected {
				return NewMaxError('g', expected), nil
			}
			return nil, nil
		}
	}
}

func MaxFloat64(expected float64) Float64Rule {
	return func(v float64) Rule {
		return func() (*Error, error) {
			if v > expected {
				return NewMaxError('g', expected), nil
			}
			return nil, nil
		}
	}
}

func MaxString(expected int) StringRule {
	return func(v string) Rule {
		return func() (*Error, error) {
			if utf8.RuneCountInString(v) > expected {
				return NewMaxError('d', expected), nil
			}
			return nil, nil
		}
	}
}
