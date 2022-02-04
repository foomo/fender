# fender

> a piece of rope or a tyre that protects the side of a boat from knocks

Fender provides a unified way to validate data and make the errors passable to the frontend e.g. through gotsrpc.

## Usage

Fender provides different ways to validate your data.

### struct validation using go-playground validator

```go
package main

import (
	"fmt"

	"github.com/foomo/fender"
)

type Test struct {
	Int     int     `validate:"min=1"`
	Int8    int8    `validate:"min=1"`
	Int32   int32   `validate:"min=1"`
	Int64   int64   `validate:"min=1"`
	UInt    uint    `validate:"min=1"`
	UInt8   uint8   `validate:"min=1"`
	UInt32  uint32  `validate:"min=1"`
	UInt64  uint64  `validate:"min=1"`
	Float32 float32 `validate:"min=1.5"`
	Float64 float64 `validate:"min=1.5"`
	Bool    bool    `validate:"required"`
	String  string  `validate:"required"`
}

func main() {
	u := Test{}
	fendErr, err := fender.Struct(u)
	if err != nil {
		panic("internal error: " + err.Error())
	}
	fmt.Println(fendErr) // bool:required;string:required;uInt:min=1;uInt8:min=1;uInt32:min=1;uInt64:min=1;float32:min=1.5;int:min=1;int8:min=1;int32:min=1;int64:min=1;float64:min=1.5
}
```

### validate by code

```go
package main

import (
	"fmt"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
)

type Test struct {
	Int     int
	Int8    int8
	Int32   int32
	Int64   int64
	UInt    uint
	UInt8   uint8
	UInt32  uint32
	UInt64  uint64
	Float32 float32
	Float64 float64
	Bool    bool
	String  string
}

func main() {
	u := Test{}
	fendErr, err := fender.All(
		fender.Field("int", fend.Int(u.Int, rule.MinInt(1))),
		fender.Field("int8", fend.Int8(u.Int8, rule.MinInt8(1))),
		fender.Field("int32", fend.Int32(u.Int32, rule.MinInt32(1))),
		fender.Field("int64", fend.Int64(u.Int64, rule.MinInt64(1))),
		fender.Field("uint", fend.UInt(u.UInt, rule.MinUInt(1))),
		fender.Field("uint8", fend.UInt8(u.UInt8, rule.MinUInt8(1))),
		fender.Field("uint32", fend.UInt32(u.UInt32, rule.MinUInt32(1))),
		fender.Field("uint64", fend.UInt64(u.UInt64, rule.MinUInt64(1))),
		fender.Field("float32", fend.Float32(u.Float32, rule.MinFloat32(1.5))),
		fender.Field("float64", fend.Float64(u.Float64, rule.MinFloat64(1.5))),
		fender.Field("bool", fend.Bool(u.Bool, rule.Bool(true))),
		fender.Field("string", fend.String(u.String, rule.RequiredString)),
	)
	if err != nil {
		panic("internal error: " + err.Error())
	}
	fmt.Println(fendErr) // uint:min=1;uint64:min=1;float32:min=1.5;int:min=1;int8:min=1;uint8:min=1;uint32:min=1;float64:min=1.5;bool:bool=true;string:required;int32:min=1;int64:min=1
}
```

### validate by code and stop on first error

```go
package main

import (
	"fmt"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
)

type Test struct {
	Int     int
	Int8    int8
	Int32   int32
	Int64   int64
	UInt    uint
	UInt8   uint8
	UInt32  uint32
	UInt64  uint64
	Float32 float32
	Float64 float64
	Bool    bool
	String  string
}

func main() {
	u := Test{}
	fendErr, err := fender.First(
		fender.Field("int", fend.Int(u.Int, rule.MinInt(1))),
		fender.Field("int8", fend.Int8(u.Int8, rule.MinInt8(1))),
		fender.Field("int32", fend.Int32(u.Int32, rule.MinInt32(1))),
		fender.Field("int64", fend.Int64(u.Int64, rule.MinInt64(1))),
		fender.Field("uint", fend.UInt(u.UInt, rule.MinUInt(1))),
		fender.Field("uint8", fend.UInt8(u.UInt8, rule.MinUInt8(1))),
		fender.Field("uint32", fend.UInt32(u.UInt32, rule.MinUInt32(1))),
		fender.Field("uint64", fend.UInt64(u.UInt64, rule.MinUInt64(1))),
		fender.Field("float32", fend.Float32(u.Float32, rule.MinFloat32(1.5))),
		fender.Field("float64", fend.Float64(u.Float64, rule.MinFloat64(1.5))),
		fender.Field("bool", fend.Bool(u.Bool, rule.Bool(true))),
		fender.Field("string", fend.String(u.String, rule.RequiredString)),
	)
	if err != nil {
		panic("internal error: " + err.Error())
	}
	fmt.Println(fendErr) // int:min=1
}
```

## References & alternatives

- [go-playground/validator](https://github.com/go-playground/validator)
- [go-ozzo/ozzo-validation](https://github.com/go-ozzo/ozzo-validation)

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.

## Notes

We'll take another refactoring round as soon as go supports generics
