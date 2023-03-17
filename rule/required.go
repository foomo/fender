package rule

import (
	"context"
	"reflect"
)

const NameRequired Name = "required"

func Required[T any](ctx context.Context, v T) error {
	if reflect.ValueOf(v).IsZero() {
		return NewError(NameRequired)
	}
	return nil
}

func IsRequired[T any](expected bool) Rule[T] {
	return func(ctx context.Context, v T) error {
		if !expected {
			return ErrBreak
		}
		if reflect.ValueOf(v).IsZero() {
			return NewError(NameRequired)
		}
		return nil
	}
}

func BoolRequired(ctx context.Context, v bool) error {
	if !v {
		return NewError(NameRequired)
	}
	return nil
}

func StringRequired(ctx context.Context, v string) error {
	if len(v) == 0 {
		return NewError(NameRequired)
	}
	return nil
}

func NumberRequired[T Number](ctx context.Context, v T) error {
	if v == T(0) {
		return NewError(NameRequired)
	}
	return nil
}
