package rule

import (
	"context"
)

const NameSize Name = "size"

func StringSize(expected int) StringRule {
	return func(ctx context.Context, v string) error {
		if len(v) != expected {
			return NewError(NameSize, Meta('d', expected))
		}
		return nil
	}
}

func StringNotSize(expected int) StringRule {
	return func(ctx context.Context, v string) error {
		if len(v) == expected {
			return NewError(NameSize, Meta('d', expected))
		}
		return nil
	}
}

func NumberSize[T Number](expected T) Rule[T] {
	return func(ctx context.Context, v T) error {
		if v != expected {
			return NewError(NameSize, Meta('d', expected))
		}
		return nil
	}
}

func NumberNotSize[T Number](expected T) Rule[T] {
	return func(ctx context.Context, v T) error {
		if v == expected {
			return NewError(NameSize, Meta('d', expected))
		}
		return nil
	}
}
