package fend

import (
	"testing"

	"github.com/foomo/fender/rule"
	"github.com/go-playground/validator/v10"
)

func BenchmarkString(b *testing.B) {
	test := "foo"

	b.Run(test, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			String("", rule.RequiredString())()
			String("", rule.RequiredString())()
			String("", rule.RequiredString())()
			String("", rule.RequiredString())()
			String("", rule.RequiredString())()
			String("", rule.RequiredString())()
			String("", rule.RequiredString())()
			String("", rule.RequiredString())()
			String("", rule.RequiredString())()
			String("", rule.RequiredString())()
		}
	})
	b.Run("playground", func(b *testing.B) {
		validate := validator.New()
		for i := 0; i < b.N; i++ {
			_ = validate.Var(test, "required")
			_ = validate.Var(test, "required")
			_ = validate.Var(test, "required")
			_ = validate.Var(test, "required")
			_ = validate.Var(test, "required")
			_ = validate.Var(test, "required")
			_ = validate.Var(test, "required")
			_ = validate.Var(test, "required")
			_ = validate.Var(test, "required")
			_ = validate.Var(test, "required")
		}
	})
}
