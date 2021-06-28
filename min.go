package fender

import (
	"fmt"
	"unicode/utf8"
)

const RuleMin Rule = "min"

// NewFendMinError constructor
func NewFendMinError(verb rune, v interface{}) *RuleError {
	return NewRuleError(RuleMin, fmt.Sprintf("%"+string(verb), v))
}

func MinInt(v, expected int) Fend {
	return func() *RuleError {
		if v < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinInt8(v, expected int8) Fend {
	return func() *RuleError {
		if v < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinInt32(v, expected int32) Fend {
	return func() *RuleError {
		if v < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinInt64(v, expected int64) Fend {
	return func() *RuleError {
		if v < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinUInt(v, expected uint) Fend {
	return func() *RuleError {
		if v < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinUInt8(v, expected uint8) Fend {
	return func() *RuleError {
		if v < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinUInt32(v, expected uint32) Fend {
	return func() *RuleError {
		if v < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinUInt64(v, expected uint64) Fend {
	return func() *RuleError {
		if v < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinFloat32(v, expected float32) Fend {
	return func() *RuleError {
		if v < expected {
			return NewFendMinError('f', expected)
		}
		return nil
	}
}

func MinFloat64(v, expected float64) Fend {
	return func() *RuleError {
		if v < expected {
			return NewFendMinError('f', expected)
		}
		return nil
	}
}

func MinString(v string, expected int) Fend {
	return func() *RuleError {
		if utf8.RuneCountInString(v) < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinInts(v []int, expected int) Fend {
	return func() *RuleError {
		if len(v) < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinInt8s(v []int8, expected int) Fend {
	return func() *RuleError {
		if len(v) < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinInt32s(v []int32, expected int) Fend {
	return func() *RuleError {
		if len(v) < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinInt64s(v []int64, expected int) Fend {
	return func() *RuleError {
		if len(v) < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinUInts(v []uint, expected int) Fend {
	return func() *RuleError {
		if len(v) < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinUInt8s(v []uint8, expected int) Fend {
	return func() *RuleError {
		if len(v) < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinUInt32s(v []uint32, expected int) Fend {
	return func() *RuleError {
		if len(v) < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinUInt64s(v []uint64, expected int) Fend {
	return func() *RuleError {
		if len(v) < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}

func MinFloat32s(v []float32, expected int) Fend {
	return func() *RuleError {
		if len(v) < expected {
			return NewFendMinError('f', expected)
		}
		return nil
	}
}

func MinFloat64s(v []float64, expected int) Fend {
	return func() *RuleError {
		if len(v) < expected {
			return NewFendMinError('f', expected)
		}
		return nil
	}
}

func MinStrings(v []string, expected int) Fend {
	return func() *RuleError {
		if len(v) < expected {
			return NewFendMinError('d', expected)
		}
		return nil
	}
}
