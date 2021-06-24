package fender

import (
	"fmt"
	"unicode/utf8"
)

const RuleSize Rule = "size"

// NewFendSizeError constructor
func NewFendSizeError(verb rune, v interface{}) *RuleError {
	return NewFendError(RuleSize, fmt.Sprintf("%"+string(verb), v))
}

func SizeInt(v, expected int) error {
	if v != expected {
		return NewFendSizeError('d', expected)
	}
	return nil
}

func SizeInt8(v, expected int8) error {
	if v != expected {
		return NewFendSizeError('d', expected)
	}
	return nil
}

func SizeInt32(v, expected int32) error {
	if v != expected {
		return NewFendSizeError('d', expected)
	}
	return nil
}

func SizeInt64(v, expected int64) error {
	if v != expected {
		return NewFendSizeError('d', expected)
	}
	return nil
}

func SizeFloat32(v, expected float32) error {
	if v != expected {
		return NewFendSizeError('f', expected)
	}
	return nil
}

func SizeFloat64(v, expected float64) error {
	if v != expected {
		return NewFendSizeError('f', expected)
	}
	return nil
}

func SizeString(v string, expected int) error {
	if utf8.RuneCountInString(v) != expected {
		return NewFendSizeError('d', expected)
	}
	return nil
}

func SizeStrings(v []string, expected int) error {
	if len(v) != expected {
		return NewFendSizeError('d', expected)
	}
	return nil
}
