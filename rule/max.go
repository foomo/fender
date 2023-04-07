package rule

import (
	"context"
)

const NameMax Name = "max"

func StringMax(expected int) StringRule {
	return func(ctx context.Context, v string) error {
		if len(v) > expected {
			return NewError(NameMax, Meta('d', expected))
		}
		return nil
	}
}

func NumberMax[T Number](expected T) Rule[T] {
	return func(ctx context.Context, v T) error {
		if v > expected {
			return NewError(NameMax, Meta('d', expected))
		}
		return nil
	}
}
