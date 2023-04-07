package rule

import (
	"context"
	"strings"
)

const NameContains Name = "contains"

func StringContains(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if !strings.Contains(v, expected) {
			return NewError(NameContains)
		}
		return nil
	}
}

func StringNotContains(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if strings.Contains(v, expected) {
			return NewError(NameContains)
		}
		return nil
	}
}
