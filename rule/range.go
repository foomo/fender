package rule

import (
	"context"
)

const NameRange Name = "range"

func StringRange(min, max int) StringRule {
	return func(ctx context.Context, v string) error {
		if len(v) < min || len(v) > max {
			return NewError(NameRange, Meta('d', min), Meta('d', max))
		}
		return nil
	}
}

func StringNotRange(min, max int) StringRule {
	return func(ctx context.Context, v string) error {
		if len(v) >= min || len(v) <= max {
			return NewError(NameRange, Meta('d', min), Meta('d', max))
		}
		return nil
	}
}

func NumberRange[T Number](min, max T) Rule[T] {
	return func(ctx context.Context, v T) error {
		if v < min || v > max {
			return NewError(NameRange, Meta('d', min), Meta('d', max))
		}
		return nil
	}
}

func NumberNotRange[T Number](min, max T) Rule[T] {
	return func(ctx context.Context, v T) error {
		if v >= min || v <= max {
			return NewError(NameRange, Meta('d', min), Meta('d', max))
		}
		return nil
	}
}
