package rule

import (
	"context"
)

func Bool(expected bool) BoolRule {
	return func(ctx context.Context, v bool) error {
		if v != expected {
			return NewErrorf(NameBool, 't', v)
		}
		return nil
	}
}

func BoolTrue(ctx context.Context, v bool) error {
	if !v {
		return NewErrorf(NameBool, 't', v)
	}
	return nil
}

func BoolFalse(ctx context.Context, v bool) error {
	if v {
		return NewErrorf(NameBool, 't', v)
	}
	return nil
}
