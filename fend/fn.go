package fend

import (
	"context"

	"github.com/foomo/fender/rule"
)

func Fn(rules ...rule.Fn) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fendFn(ctx, mode, "", rules...)
	}
}
