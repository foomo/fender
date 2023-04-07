package fend_test

import (
	"context"
	"fmt"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
)

func ExampleDynamic() {
	dynamicRule := func(ctx context.Context) error {
		return rule.NewError("bar")
	}
	err := fender.All(context.TODO(),
		fend.DynamicVar(dynamicRule),
		fend.DynamicField("foo", dynamicRule),
	)
	fmt.Println(err)
	// Output: bar;foo:bar
}
