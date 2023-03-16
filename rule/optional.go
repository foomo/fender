package rule

import (
	"context"
	"reflect"
)

func Optional[T any](ctx context.Context, v T) error {
	if reflect.ValueOf(v).IsZero() {
		return ErrBreak
	}
	return nil
}

func BoolOptional(ctx context.Context, v bool) error {
	if !v {
		return ErrBreak
	}
	return nil
}

func StringOptional(ctx context.Context, v string) error {
	if len(v) == 0 {
		return ErrBreak
	}
	return nil
}

func NumberOptional[T Number](ctx context.Context, v T) error {
	if v == T(0) {
		return ErrBreak
	}
	return nil
}
