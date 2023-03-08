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
		fend.Field("one", "", rule.RequiredString, rule.MinString(10)),
		fend.Field("two", "", rule.RequiredString, rule.MinString(10)),
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
		fend.Field("one", "", rule.RequiredString, rule.MinString(10)),
		fend.Field("two", "", rule.RequiredString, rule.MinString(10)),
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
		fend.Field("one", "", rule.RequiredString, rule.MinString(10)),
		fend.Field("two", "", rule.RequiredString, rule.MinString(10)),
	)

	// check for fender error
	if fendErr := fender.AsError(err); fendErr != nil {
		fmt.Println(err)
	} else if err != nil {
		panic(err)
	}
	// Output: one:required:min=10
}

func ExampleErrors() { //nolint:govet
	err := fender.All(
		context.Background(),
		fend.Field("one", "", rule.RequiredString, rule.MinString(10)),
		fend.Field("two", "", rule.RequiredString, rule.MinString(10)),
	)

	// cast fender error
	if fenderErr := fender.AsError(err); fenderErr != nil {
		// iterate fend errors
		for _, err := range fenderErr.Errors() {
			// cast fend error
			if fendErr := fend.AsError(err); fendErr != nil {
				fmt.Println(fendErr.Name())
				// iterate rule errors
				for _, err := range fendErr.Errors() {
					// cast rule error
					if ruleErr := rule.AsError(err); ruleErr != nil {
						fmt.Println(ruleErr.Error())
					}
				}
			}
		}
	} else if err != nil {
		panic(err)
	}
	// Output:
	// one
	// required
	// min=10
	// two
	// required
	// min=10
}
