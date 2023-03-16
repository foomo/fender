package rule

import (
	"context"
	"strings"
)

const NameSuffix Name = "suffix"

func Suffix(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if !strings.HasSuffix(v, expected) {
			return NewError(NameSuffix, Meta('d', expected))
		}
		return nil
	}
}

func NotSuffix(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if strings.HasSuffix(v, expected) {
			return NewError(NameSuffix, Meta('d', expected))
		}
		return nil
	}
}
