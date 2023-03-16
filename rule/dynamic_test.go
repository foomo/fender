package rule_test

import (
	"context"
	"fmt"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
)

func ExampleDynamic() {
	err := fender.All(context.TODO(),
		fend.Field("demo", "foo", rule.Dynamic[string](
			func(ctx context.Context) error {
				return rule.NewError("bar")
			},
		)),
	)
	fmt.Println(err)
	// Output: demo:bar
}
