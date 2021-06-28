package fender

import (
	"fmt"
	"unicode/utf8"
)

const RuleMax Rule = "max"

// NewFendMaxError constructor
func NewFendMaxError(verb rune, v interface{}) *RuleError {
	return NewRuleError(RuleMax, fmt.Sprintf("%"+string(verb), v))
}

func MaxInt(v, expected int) Fend {
	return func() *RuleError {
		if v > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxInt8(v, expected int8) Fend {
	return func() *RuleError {
		if v > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxInt32(v, expected int32) Fend {
	return func() *RuleError {
		if v > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxInt64(v, expected int64) Fend {
	return func() *RuleError {
		if v > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxUInt(v, expected uint) Fend {
	return func() *RuleError {
		if v > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxUInt8(v, expected uint8) Fend {
	return func() *RuleError {
		if v > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxUInt32(v, expected uint32) Fend {
	return func() *RuleError {
		if v > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxUInt64(v, expected uint64) Fend {
	return func() *RuleError {
		if v > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxFloat32(v, expected float32) Fend {
	return func() *RuleError {
		if v > expected {
			return NewFendMaxError('f', expected)
		}
		return nil
	}
}

func MaxFloat64(v, expected float64) Fend {
	return func() *RuleError {
		if v > expected {
			return NewFendMaxError('f', expected)
		}
		return nil
	}
}

func MaxString(v string, expected int) Fend {
	return func() *RuleError {
		if utf8.RuneCountInString(v) > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxInts(v []int, expected int) Fend {
	return func() *RuleError {
		if len(v) > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxInt8s(v []int8, expected int) Fend {
	return func() *RuleError {
		if len(v) > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxInt32s(v []int32, expected int) Fend {
	return func() *RuleError {
		if len(v) > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxInt64s(v []int64, expected int) Fend {
	return func() *RuleError {
		if len(v) > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxUInts(v []uint, expected int) Fend {
	return func() *RuleError {
		if len(v) > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxUInt8s(v []uint8, expected int) Fend {
	return func() *RuleError {
		if len(v) > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxUInt32s(v []uint32, expected int) Fend {
	return func() *RuleError {
		if len(v) > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxUInt64s(v []uint64, expected int) Fend {
	return func() *RuleError {
		if len(v) > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}

func MaxFloat32s(v []float32, expected int) Fend {
	return func() *RuleError {
		if len(v) > expected {
			return NewFendMaxError('f', expected)
		}
		return nil
	}
}

func MaxFloat64s(v []float64, expected int) Fend {
	return func() *RuleError {
		if len(v) > expected {
			return NewFendMaxError('f', expected)
		}
		return nil
	}
}

func MaxStrings(v []string, expected int) Fend {
	return func() *RuleError {
		if len(v) > expected {
			return NewFendMaxError('d', expected)
		}
		return nil
	}
}
