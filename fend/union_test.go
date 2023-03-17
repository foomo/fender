package fend_test

import (
	"context"
	"fmt"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
)

func ExampleUnion() {
	email := fend.NewRules(rule.StringMin(5), rule.Email)
	phone := fend.NewRules(rule.StringMin(3), rule.Numeric)
	emailOrPhone := fend.NewRules(fend.Union(email, phone))

	{
		err := fender.All(context.TODO(),
			fend.Var("foo", fend.Union(email, phone)),
		)
		fmt.Println(err)
	}
	{
		err := fender.All(context.TODO(),
			emailOrPhone.Var("foo"),
		)
		fmt.Println(err)
	}

	{
		err := fender.All(context.TODO(),
			fend.Var("foo@bar.com", fend.Union(email, phone)),
		)
		fmt.Println(err)
	}
	{
		err := fender.All(context.TODO(),
			emailOrPhone.Var("foo@bar.com"),
		)
		fmt.Println(err)
	}

	{
		err := fender.All(context.TODO(),
			fend.Var("123456", fend.Union(email, phone)),
		)
		fmt.Println(err)
	}
	{
		err := fender.All(context.TODO(),
			emailOrPhone.Var("123456"),
		)
		fmt.Println(err)
	}

	// Output:
	// numeric=^[0-9]+$
	// numeric=^[0-9]+$
	// <nil>
	// <nil>
	// <nil>
	// <nil>
}
