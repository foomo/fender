package rule

import (
	"context"
)

const NameMin Name = "min"

func StringMin(expected int) StringRule {
	return func(ctx context.Context, v string) error {
		if len(v) < expected {
			return NewError(NameMin, Meta('d', expected))
		}
		return nil
	}
}

func NumberMin[T Number](expected T) Rule[T] {
	return func(ctx context.Context, v T) error {
		if v < expected {
			return NewError(NameMin, Meta('d', expected))
		}
		return nil
	}
}
