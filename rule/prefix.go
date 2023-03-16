package rule

import (
	"context"
	"strings"
)

const NamePrefix Name = "prefix"

func Prefix(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if !strings.HasPrefix(v, expected) {
			return NewError(NamePrefix, Meta('d', expected))
		}
		return nil
	}
}

func NotPrefix(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if strings.HasPrefix(v, expected) {
			return NewError(NamePrefix, Meta('d', expected))
		}
		return nil
	}
}
