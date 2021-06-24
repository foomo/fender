# fender

> a piece of rope or a tyre that protects the side of a boat from knocks

## Usage

```go
err := fender.NewFields().
	Add("firstName", fender.RequiredString(u.FirstName)).
	Add("lastName", fender.RequiredString(u.LastName)).
	Add("age", fender.MinInt(u.Age, 0), fender.MaxInt(u.Age, 130)).
	Add("email", fender.Email(u.Email, false)).
	Add("street", fender.RequiredString(u.Street)).
	Add("city", fender.RequiredString(u.City)).
	Add("planet", fender.RequiredString(u.Planet)).
	Add("phone", fender.RequiredString(u.Phone)).
	AllError()
if err != nil {
	panic(err)
}
```

## Rules

List of built-in fending rules:

- email
- enum
- max
- min
- required
- regex
- size

## Benchmarks

```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkFender/valid-12           	     1496362	       699 ns/op
BenchmarkFender/invalid-12               1176252 	       975 ns/op
BenchmarkPlayground/valid-12           	   42367	     25967 ns/op
BenchmarkPlayground/invalid-12         	   37567	     28034 ns/op
```

## References & alternatives

- [go-playground/validator](https://github.com/go-playground/validator)
- [go-ozzo/ozzo-validation](https://github.com/go-ozzo/ozzo-validation)

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.
