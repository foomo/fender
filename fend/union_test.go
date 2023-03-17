package fend_test

import (
	"context"
	"fmt"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
)

func ExampleUnion() {
	{
		err := fender.All(context.TODO(),
			fend.Var("foo", fend.Union(
				rule.StringMin(10),
				rule.Equal("bar"),
			)),
		)
		fmt.Println(err)
	}

	{
		err := fender.All(context.TODO(),
			fend.Var("foo", fend.Union(
				rule.StringMin(10),
				rule.Equal("foo"),
			)),
		)
		fmt.Println(err)
	}
	// Output:
	// equal=bar
	// <nil>
}
