package fender_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := fender.All(
			context.TODO(),
			fend.Field("foo", "foo", rule.RequiredString, rule.MinString(1)),
			fend.Field("bar", "foo", rule.RequiredString, rule.MinString(1)),
		)
		assert.NoError(t, err)
	})

	t.Run("errors", func(t *testing.T) {
		err := fender.All(
			context.TODO(),
			fend.Field("foo", "", rule.RequiredString, rule.MinString(10)),
			fend.Field("bar", "bar", rule.RequiredString, rule.MinString(10)),
		)
		if fendErr := fender.AsError(err); assert.NotNil(t, fendErr) {
			errs := fendErr.Errors()
			assert.Len(t, errs, 2)
		}
		assert.EqualError(t, err, "foo:required:min=10;bar:min=10")
	})

	t.Run("errors combined", func(t *testing.T) {
		err := fender.All(
			context.TODO(),
			fend.Field("foo", "", rule.RequiredString, rule.MinString(10)),
			fend.Field("foo", "", rule.MinString(10), rule.RequiredString),
		)
		if fendErr := fender.AsError(err); assert.NotNil(t, fendErr) {
			errs := fendErr.Errors()
			assert.Len(t, errs, 2)
		}
		assert.EqualError(t, err, "foo:required:min=10;foo:min=10:required")
	})

	t.Run("return std error", func(t *testing.T) {
		e := errors.New("std error")
		err := fender.All(
			context.TODO(),
			fend.Field("one", "", rule.RequiredString),
			fend.Field("foo", "", func(ctx context.Context, v string) error {
				return e
			}),
		)
		assert.Nil(t, fender.AsError(err))
		assert.EqualError(t, err, e.Error())
	})
}

func ExampleAll() {
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

	u := Test{}
	err := fender.All(
		context.Background(),
		fend.Field("int", u.Int, rule.MinInt(1)),
		fend.Field("int8", u.Int8, rule.MinInt8(1)),
		fend.Field("int32", u.Int32, rule.MinInt32(1)),
		fend.Field("int64", u.Int64, rule.MinInt64(1)),
		fend.Field("uint", u.UInt, rule.MinUInt(1)),
		fend.Field("uint8", u.UInt8, rule.MinUInt8(1)),
		fend.Field("uint32", u.UInt32, rule.MinUInt32(1)),
		fend.Field("uint64", u.UInt64, rule.MinUInt64(1)),
		fend.Field("float32", u.Float32, rule.MinFloat32(1.5)),
		fend.Field("float64", u.Float64, rule.MinFloat64(1.5)),
		fend.Field("bool", u.Bool, rule.Bool(true)),
		fend.Field("string", u.String, rule.RequiredString),
	)

	// check for fender error
	if fendErr := fender.AsError(err); fendErr != nil {
		fmt.Println(err)
	} else if err != nil {
		panic(err)
	}
	// Output: int:min=1;int8:min=1;int32:min=1;int64:min=1;uint:min=1;uint8:min=1;uint32:min=1;uint64:min=1;float32:min=1.5;float64:min=1.5;bool:bool=true;string:required
}

// https://github.com/frederikhors/bench-golang-validators
func BenchmarkAll(b *testing.B) {
	type Test struct {
		Int     int     `validate:"required,min=1,max=5"`
		Int8    int8    `validate:"required,min=1,max=5"`
		Int32   int32   `validate:"required,min=1,max=5"`
		Int64   int64   `validate:"required,min=1,max=5"`
		UInt    uint    `validate:"required,min=1,max=5"`
		UInt8   uint8   `validate:"required,min=1,max=5"`
		UInt32  uint32  `validate:"required,min=1,max=5"`
		UInt64  uint64  `validate:"required,min=1,max=5"`
		Float32 float32 `validate:"required,min=1,max=5"`
		Float64 float64 `validate:"required,min=1,max=5"`
		Bool    bool    `validate:"required"`
		String  string  `validate:"required"`
	}
	v := validator.New()

	b.Run("invalid reused", func(b *testing.B) {
		u := &Test{}

		b.Run("fender", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = fender.All(
					context.TODO(),
					fend.Field("int", u.Int, rule.RequiredInt, rule.MinInt(1), rule.MaxInt(5)),
					fend.Field("int8", u.Int8, rule.RequiredInt8, rule.MinInt8(1), rule.MaxInt8(5)),
					fend.Field("int32", u.Int32, rule.RequiredInt32, rule.MinInt32(1), rule.MaxInt32(5)),
					fend.Field("int64", u.Int64, rule.RequiredInt64, rule.MinInt64(1), rule.MaxInt64(5)),
					fend.Field("uint", u.UInt, rule.RequiredUInt, rule.MinUInt(1), rule.MaxUInt(5)),
					fend.Field("uint8", u.UInt8, rule.RequiredUInt8, rule.MinUInt8(1), rule.MaxUInt8(5)),
					fend.Field("uint32", u.UInt32, rule.RequiredUInt32, rule.MinUInt32(1), rule.MaxUInt32(5)),
					fend.Field("uint64", u.UInt64, rule.RequiredUInt64, rule.MinUInt64(1), rule.MaxUInt64(5)),
					fend.Field("float32", u.Float32, rule.RequiredFloat32, rule.MinFloat32(1), rule.MaxFloat32(5)),
					fend.Field("float64", u.Float64, rule.RequiredFloat64, rule.MinFloat64(1), rule.MaxFloat64(5)),
					fend.Field("bool", u.Bool, rule.Bool(true)),
					fend.Field("string", u.String, rule.RequiredString),
				)
			}
		})
		b.Run("playground", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = v.Struct(u)
			}
		})
	})

	b.Run("success new", func(b *testing.B) {
		u := &Test{
			Int:     1,
			Int8:    1,
			Int32:   1,
			Int64:   1,
			UInt:    1,
			UInt8:   1,
			UInt32:  1,
			UInt64:  1,
			Float32: 1,
			Float64: 1,
			Bool:    true,
			String:  "true",
		}
		b.Run("fender", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = fender.All(
					context.TODO(),
					fend.Field("int", u.Int, rule.RequiredInt, rule.MinInt(1), rule.MaxInt(5)),
					fend.Field("int8", u.Int8, rule.RequiredInt8, rule.MinInt8(1), rule.MaxInt8(5)),
					fend.Field("int32", u.Int32, rule.RequiredInt32, rule.MinInt32(1), rule.MaxInt32(5)),
					fend.Field("int64", u.Int64, rule.RequiredInt64, rule.MinInt64(1), rule.MaxInt64(5)),
					fend.Field("uint", u.UInt, rule.RequiredUInt, rule.MinUInt(1), rule.MaxUInt(5)),
					fend.Field("uint8", u.UInt8, rule.RequiredUInt8, rule.MinUInt8(1), rule.MaxUInt8(5)),
					fend.Field("uint32", u.UInt32, rule.RequiredUInt32, rule.MinUInt32(1), rule.MaxUInt32(5)),
					fend.Field("uint64", u.UInt64, rule.RequiredUInt64, rule.MinUInt64(1), rule.MaxUInt64(5)),
					fend.Field("float32", u.Float32, rule.RequiredFloat32, rule.MinFloat32(1), rule.MaxFloat32(5)),
					fend.Field("float64", u.Float64, rule.RequiredFloat64, rule.MinFloat64(1), rule.MaxFloat64(5)),
					fend.Field("bool", u.Bool, rule.Bool(true)),
					fend.Field("string", u.String, rule.RequiredString),
				)
			}
		})

		b.Run("playground", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = v.Struct(u)
			}
		})
	})
}

/*
v0
BenchmarkSimpleStruct/invalid_reused/fender-12  	  				  220735	      4943 ns/op
BenchmarkSimpleStruct/invalid_reused/playground-12         	  405120	      2849 ns/op
BenchmarkSimpleStruct/success_new/fender-12                	  177958	      5747 ns/op
BenchmarkSimpleStruct/success_new/playground-12            	  719839	      1578 ns/op

v1
BenchmarkSimpleStruct/invalid_reused/fender-10  	            316605	      4050 ns/op
BenchmarkSimpleStruct/invalid_reused/playground-10         	  683263	      1741 ns/op
BenchmarkSimpleStruct/success_new/fender-10                	 1000000	      1036 ns/op
BenchmarkSimpleStruct/success_new/playground-10            	 1312262	       919 ns/op
*/

func TestFirst(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := fender.First(
			context.TODO(),
			fend.Field("foo", "", rule.RequiredString),
		)
		if fendErr := fender.AsError(err); assert.NotNil(t, fendErr) {
			first := fender.AsError(err).First()
			assert.NotNil(t, rule.AsError(first))
		}
		assert.EqualError(t, err, "foo:required")
	})

	t.Run("error string", func(t *testing.T) {
		err := fender.First(
			context.TODO(),
			fend.Field("foo", "", rule.MinString(10)),
		)
		if fendErr := fender.AsError(err); assert.NotNil(t, fendErr) {
			first := fender.AsError(err).First()
			assert.NotNil(t, rule.AsError(first))
		}
		assert.EqualError(t, err, "foo:min=10")
	})

	t.Run("error var", func(t *testing.T) {
		err := fender.First(
			context.TODO(),
			fend.Field("foo", "", rule.Var("min=10")),
		)
		if fendErr := fender.AsError(err); assert.NotNil(t, fendErr) {
			first := fender.AsError(err).First()
			assert.NotNil(t, rule.AsError(first))
		}
		assert.EqualError(t, err, "foo:min=10")
	})

	t.Run("return std error", func(t *testing.T) {
		e := errors.New("std error")
		err := fender.First(
			context.TODO(),
			fend.Field("foo", "", func(ctx context.Context, v string) error {
				return e
			}),
		)
		assert.Nil(t, fender.AsError(err))
		assert.EqualError(t, err, e.Error())
	})
}

func ExampleFirst() {
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

	u := Test{}
	err := fender.First(
		context.Background(),
		fend.Field("int", u.Int, rule.MinInt(1)),
		fend.Field("int8", u.Int8, rule.MinInt8(1)),
		fend.Field("int32", u.Int32, rule.MinInt32(1)),
		fend.Field("int64", u.Int64, rule.MinInt64(1)),
		fend.Field("uint", u.UInt, rule.MinUInt(1)),
		fend.Field("uint8", u.UInt8, rule.MinUInt8(1)),
		fend.Field("uint32", u.UInt32, rule.MinUInt32(1)),
		fend.Field("uint64", u.UInt64, rule.MinUInt64(1)),
		fend.Field("float32", u.Float32, rule.MinFloat32(1.5)),
		fend.Field("float64", u.Float64, rule.MinFloat64(1.5)),
		fend.Field("bool", u.Bool, rule.Bool(true)),
		fend.Field("string", u.String, rule.RequiredString),
	)

	// check for fender error
	if fendErr := fender.AsError(err); fendErr != nil {
		fmt.Println(err)
	} else if err != nil {
		panic(err)
	}
	// Output: int:min=1
}
