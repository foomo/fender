package fend

import (
	"context"

	"github.com/foomo/fender/rule"
)

func DynamicField(path string, rules ...rule.DynamicRule) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fendDynamic(ctx, mode, path, rules...)
	}
}

func DynamicVar(rules ...rule.DynamicRule) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fendDynamic(ctx, mode, "", rules...)
	}
}
