package fender_test

import (
	"context"
	"fmt"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
)

func ExampleAll() {
	err := fender.All(
		context.Background(),
		fend.Field("one", "", rule.Required[string], rule.StringMin(10)),
		fend.Field("two", "", rule.Required[string], rule.StringMin(10)),
	)
	// check for fender error
	if fendErr := fender.AsError(err); fendErr != nil {
		fmt.Println(err)
	} else if err != nil {
		panic(err)
	}
	// Output: one:required:min=10;two:required:min=10
}

func ExampleFirst() {
	err := fender.First(
		context.Background(),
		fend.Field("one", "", rule.Required[string], rule.StringMin(10)),
		fend.Field("two", "", rule.Required[string], rule.StringMin(10)),
	)

	// check for fender error
	if fendErr := fender.AsError(err); fendErr != nil {
		fmt.Println(err)
	} else if err != nil {
		panic(err)
	}
	// Output: one:required
}

func ExampleAllFirst() {
	err := fender.AllFirst(
		context.Background(),
		fend.Field("one", "", rule.Required[string], rule.StringMin(10)),
		fend.Field("two", "", rule.Required[string], rule.StringMin(10)),
	)

	// check for fender error
	if fendErr := fender.AsError(err); fendErr != nil {
		fmt.Println(err)
	} else if err != nil {
		panic(err)
	}
	// Output: one:required:min=10
}
