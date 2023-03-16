package rule

import (
	"context"
)

const Max Name = "min"

func StringMax(expected int) StringRule {
	return func(ctx context.Context, v string) error {
		if len(v) > expected {
			return NewError(Max, Meta('d', expected))
		}
		return nil
	}
}

func NumberMax[T Number](expected T) Rule[T] {
	return func(ctx context.Context, v T) error {
		if v > expected {
			return NewError(Max, Meta('d', expected))
		}
		return nil
	}
}
