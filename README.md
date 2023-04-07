# fender

> a piece of rope or a tyre that protects the side of a boat from knocks

Fender provides a unified way to validate data and on the backend & frontend side.

## Usage

### All

```go
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
```

### AllFirst

```go
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
```

### First

```go
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
```

### Handle errors

```go
func ExampleErrors() {
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
```

## Benchmarks

```text
go test -bench=. | prettybench
goos: darwin
goarch: arm64
pkg: github.com/foomo/fender
PASS
benchmark                                      iter       time/iter
---------                                      ----       ---------
BenchmarkAll/invalid_reused/fender-10        282154   3771.00 ns/op
BenchmarkAll/invalid_reused/playground-10    678679   1741.00 ns/op
BenchmarkAll/success_new/fender-10          1000000   1026.00 ns/op
BenchmarkAll/success_new/playground-10      1308327    937.60 ns/op
ok  	github.com/foomo/fender	5.619s
```

## References & alternatives

- [go-playground/validator](https://github.com/go-playground/validator)
- [go-ozzo/ozzo-validation](https://github.com/go-ozzo/ozzo-validation)

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.
