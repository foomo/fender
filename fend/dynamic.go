package fend

import (
	"context"

	"github.com/foomo/fender/rule"
)

func DynamicField(name string, rules ...rule.DynamicRule) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fendDynamic(ctx, mode, name, rules...)
	}
}

func DynamicVar(rules ...rule.DynamicRule) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fendDynamic(ctx, mode, "", rules...)
	}
}
