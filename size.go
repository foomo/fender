package fender

import (
	"fmt"
	"unicode/utf8"
)

const RuleSize Rule = "size"

// NewFendSizeError constructor
func NewFendSizeError(verb rune, v interface{}) *RuleError {
	return NewRuleError(RuleSize, fmt.Sprintf("%"+string(verb), v))
}

func SizeInt(v, expected int) Fend {
	return func() *RuleError {
		if v != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeInt8(v, expected int8) Fend {
	return func() *RuleError {
		if v != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeInt32(v, expected int32) Fend {
	return func() *RuleError {
		if v != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeInt64(v, expected int64) Fend {
	return func() *RuleError {
		if v != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeUInt(v, expected uint) Fend {
	return func() *RuleError {
		if v != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeUInt8(v, expected uint8) Fend {
	return func() *RuleError {
		if v != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeUInt32(v, expected uint32) Fend {
	return func() *RuleError {
		if v != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeUInt64(v, expected uint64) Fend {
	return func() *RuleError {
		if v != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeFloat32(v, expected float32) Fend {
	return func() *RuleError {
		if v != expected {
			return NewFendSizeError('f', expected)
		}
		return nil
	}
}

func SizeFloat64(v, expected float64) Fend {
	return func() *RuleError {
		if v != expected {
			return NewFendSizeError('f', expected)
		}
		return nil
	}
}

func SizeString(v string, expected int) Fend {
	return func() *RuleError {
		if utf8.RuneCountInString(v) != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeInts(v []int, expected int) Fend {
	return func() *RuleError {
		if len(v) != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeInt8s(v []int8, expected int) Fend {
	return func() *RuleError {
		if len(v) != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeInt32s(v []int32, expected int) Fend {
	return func() *RuleError {
		if len(v) != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeInt64s(v []int64, expected int) Fend {
	return func() *RuleError {
		if len(v) != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeUInts(v []uint, expected int) Fend {
	return func() *RuleError {
		if len(v) != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeUInt8s(v []uint8, expected int) Fend {
	return func() *RuleError {
		if len(v) != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeUInt32s(v []uint32, expected int) Fend {
	return func() *RuleError {
		if len(v) != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeUInt64s(v []uint64, expected int) Fend {
	return func() *RuleError {
		if len(v) != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}

func SizeFloat32s(v []float32, expected int) Fend {
	return func() *RuleError {
		if len(v) != expected {
			return NewFendSizeError('f', expected)
		}
		return nil
	}
}

func SizeFloat64s(v []float64, expected int) Fend {
	return func() *RuleError {
		if len(v) != expected {
			return NewFendSizeError('f', expected)
		}
		return nil
	}
}

func SizeStrings(v []string, expected int) Fend {
	return func() *RuleError {
		if len(v) != expected {
			return NewFendSizeError('d', expected)
		}
		return nil
	}
}
