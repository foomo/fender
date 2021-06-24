package fender

import (
	"fmt"
)

const RuleMin Rule = "min"

// NewFendMinError constructor
func NewFendMinError(verb rune, v interface{}) *RuleError {
	return NewFendError(RuleMin, fmt.Sprintf("%"+string(verb), v))
}

func MinInt(v, expected int) error {
	if v < expected {
		return NewFendMinError('d', expected)
	}
	return nil
}

func MinInt8(v, expected int8) error {
	if v < expected {
		return NewFendMinError('d', expected)
	}
	return nil
}

func MinInt32(v, expected int32) error {
	if v < expected {
		return NewFendMinError('d', expected)
	}
	return nil
}

func MinInt64(v, expected int64) error {
	if v < expected {
		return NewFendMinError('d', expected)
	}
	return nil
}

func MinFloat32(v, expected float32) error {
	if v < expected {
		return NewFendMinError('f', expected)
	}
	return nil
}

func MinFloat64(v, expected float64) error {
	if v < expected {
		return NewFendMinError('f', expected)
	}
	return nil
}

func MinString(v string, expected int) error {
	if len([]rune(v)) < expected {
		return NewFendMinError('d', expected)
	}
	return nil
}
