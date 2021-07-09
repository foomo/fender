package fend_test

import (
	"fmt"
	"testing"

	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
	"github.com/go-playground/validator/v10"
)

func BenchmarkString(b *testing.B) {
	test := ""
	validate := validator.New()

	b.Run("all", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fend.All(
				fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
			)
		}
	})
	b.Run("first", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := fend.First(
				fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
				// fend.Var(test, "required,min=10"),
			); err != nil {
				fmt.Println(err)
			}
		}
	})
	b.Run("fender all", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fend.All(
				fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
			)
		}
	})
	b.Run("fender first", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fend.First(
				fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
				// fend.String(test, rule.MinString(10), rule.RequiredString),
			)
		}
	})
	b.Run("playground", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = validate.Var(test, "required,min=10")
			// _ = validate.Var(test, "required,min=10")
			// _ = validate.Var(test, "required,min=10")
			// _ = validate.Var(test, "required,min=10")
			// _ = validate.Var(test, "required,min=10")
			// _ = validate.Var(test, "required,min=10")
			// _ = validate.Var(test, "required,min=10")
			// _ = validate.Var(test, "required,min=10")
			// _ = validate.Var(test, "required,min=10")
			// _ = validate.Var(test, "required,min=10")
		}
	})
}
