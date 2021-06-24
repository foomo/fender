package fender

import (
	"fmt"
	"unicode/utf8"
)

const RuleMax Rule = "max"

// NewFendMaxError constructor
func NewFendMaxError(verb rune, v interface{}) *RuleError {
	return NewFendError(RuleMax, fmt.Sprintf("%"+string(verb), v))
}

func MaxInt(v, expected int) error {
	if v > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxInt8(v, expected int8) error {
	if v > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxInt32(v, expected int32) error {
	if v > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxInt64(v, expected int64) error {
	if v > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxUInt(v, expected uint) error {
	if v > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxUInt8(v, expected uint8) error {
	if v > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxUInt32(v, expected uint32) error {
	if v > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxUInt64(v, expected uint64) error {
	if v > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxFloat32(v, expected float32) error {
	if v > expected {
		return NewFendMaxError('f', expected)
	}
	return nil
}

func MaxFloat64(v, expected float64) error {
	if v > expected {
		return NewFendMaxError('f', expected)
	}
	return nil
}

func MaxString(v string, expected int) error {
	if utf8.RuneCountInString(v) > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxInts(v []int, expected int) error {
	if len(v) > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxInt8s(v []int8, expected int) error {
	if len(v) > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxInt32s(v []int32, expected int) error {
	if len(v) > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxInt64s(v []int64, expected int) error {
	if len(v) > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxUInts(v []uint, expected int) error {
	if len(v) > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxUInt8s(v []uint8, expected int) error {
	if len(v) > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxUInt32s(v []uint32, expected int) error {
	if len(v) > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxUInt64s(v []uint64, expected int) error {
	if len(v) > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}

func MaxFloat32s(v []float32, expected int) error {
	if len(v) > expected {
		return NewFendMaxError('f', expected)
	}
	return nil
}

func MaxFloat64s(v []float64, expected int) error {
	if len(v) > expected {
		return NewFendMaxError('f', expected)
	}
	return nil
}

func MaxStrings(v []string, expected int) error {
	if len(v) > expected {
		return NewFendMaxError('d', expected)
	}
	return nil
}
