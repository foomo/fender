package fender_test

import (
	"context"
	"testing"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
	"github.com/go-playground/validator/v10"
)

// https://github.com/frederikhors/bench-golang-validators

/*
benchmark                           iter        time/iter   bytes alloc          allocs
---------                           ----        ---------   -----------          ------
BenchmarkAll/playground/v1-10    5310297     214.30 ns/op      208 B/op     6 allocs/op
BenchmarkAll/playground/v2-10      77941   14893.00 ns/op    15394 B/op   185 allocs/op
BenchmarkAll/playground/v3-10   10654539     113.20 ns/op        8 B/op     1 allocs/op
BenchmarkAll/playground/v4-10      79444   14544.00 ns/op    15200 B/op   180 allocs/op
BenchmarkAll/fender/v1-10        3328400     364.40 ns/op      360 B/op    11 allocs/op
BenchmarkAll/fender/v2-10        3569961     332.50 ns/op      328 B/op     9 allocs/op
BenchmarkAll/fender/v3-10        9554830     126.90 ns/op      136 B/op     4 allocs/op
BenchmarkAll/fender/v4-10       12320528      97.84 ns/op      104 B/op     2 allocs/op    919 ns/op
*/

func BenchmarkAll(b *testing.B) {
	type Test struct {
		Int int `validate:"required,min=1,max=5"`
	}
	b.Run("playground", func(b *testing.B) {
		b.Run("v1", func(b *testing.B) {
			u := &Test{}
			v := validator.New()
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = v.Struct(u)
			}
		})
		b.Run("v2", func(b *testing.B) {
			u := &Test{}
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = validator.New().Struct(u)
			}
		})
		b.Run("v3", func(b *testing.B) {
			u := &Test{Int: 3}
			v := validator.New()
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = v.Struct(u)
			}
		})
		b.Run("v4", func(b *testing.B) {
			u := &Test{Int: 3}
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = validator.New().Struct(u)
			}
		})
	})

	b.Run("fender", func(b *testing.B) {
		b.Run("v1", func(b *testing.B) {
			u := &Test{}
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = fender.All(context.TODO(),
					fend.Field("int", u.Int, rule.Required[int], rule.NumberMin[int](1), rule.NumberMax[int](5)),
				)
			}
		})
		b.Run("v2", func(b *testing.B) {
			u := &Test{}
			rules := fend.NewRules(rule.Required[int], rule.NumberMin[int](1), rule.NumberMax[int](5))
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = fender.All(context.TODO(),
					rules.Field("int", u.Int),
				)
			}
		})
		b.Run("v3", func(b *testing.B) {
			u := &Test{Int: 3}
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = fender.All(context.TODO(),
					fend.Field("int", u.Int, rule.Required[int], rule.NumberMin[int](1), rule.NumberMax[int](5)),
				)
			}
		})
		b.Run("v4", func(b *testing.B) {
			u := &Test{Int: 3}
			rules := fend.NewRules(rule.Required[int], rule.NumberMin[int](1), rule.NumberMax[int](5))
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = fender.All(context.TODO(),
					rules.Field("int", u.Int),
				)
			}
		})
	})
}
