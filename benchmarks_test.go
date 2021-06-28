package fender_test

import (
	"testing"

	"github.com/foomo/fender"
	"github.com/go-playground/validator/v10"
)

// https://github.com/frederikhors/bench-golang-validators

func BenchmarkSimpleStruct(b *testing.B) {
	type Test struct {
		FirstName string `validate:"required"`
		LastName  string `validate:"required"`
		Age       int    `validate:"required,min=18"`
		Email     string `validate:"email"`
		Street    string `validate:"required"`
		City      string `validate:"required"`
		Planet    string `validate:"required"`
		Phone     string `validate:"required"`
	}

	b.Run("failure", func(b *testing.B) {
		u := &Test{}

		b.Run("fender", func(b *testing.B) {
			fields := fender.NewFields().
				Add("firstName", fender.RequiredString(u.FirstName)).
				Add("lastName", fender.RequiredString(u.LastName)).
				Add("age", fender.RequiredInt(u.Age), fender.MinInt(u.Age, 18)).
				Add("email", fender.Email(u.Email, false)).
				Add("street", fender.RequiredString(u.Street)).
				Add("city", fender.RequiredString(u.City)).
				Add("planet", fender.RequiredString(u.Planet)).
				Add("phone", fender.RequiredString(u.Phone))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if err := fields.AllFieldsErrors().Error(); err != nil {
					b.Log("Fender: " + err.Error())
				}
			}
		})

		b.Run("playground", func(b *testing.B) {
			v := validator.New()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if err := v.Struct(u); err != nil {
					b.Log("Playground: " + err.Error())
				}
			}
		})
	})

	b.Run("success", func(b *testing.B) {
		u := &Test{
			FirstName: "Badger",
			LastName:  "Smith",
			Age:       80,
			City:      "string",
			Email:     "Badger.Smith@email.com",
			Street:    "Eavesdown Docks",
			Planet:    "Persphone",
			Phone:     "none",
		}
		b.Run("fender", func(b *testing.B) {
			fields := fender.NewFields().
				Add("firstName", fender.RequiredString(u.FirstName)).
				Add("lastName", fender.RequiredString(u.LastName)).
				Add("age", fender.MinInt(u.Age, 0), fender.MaxInt(u.Age, 130)).
				Add("email", fender.Email(u.Email, false)).
				Add("street", fender.RequiredString(u.Street)).
				Add("city", fender.RequiredString(u.City)).
				Add("planet", fender.RequiredString(u.Planet)).
				Add("phone", fender.RequiredString(u.Phone))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if err := fields.AllFieldsErrors().Error(); err != nil {
					b.Log("Fender: " + err.Error())
				}
			}
		})

		b.Run("playground", func(b *testing.B) {
			v := validator.New()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if err := v.Struct(u); err != nil {
					b.Log("Playground: " + err.Error())
				}
			}
		})
	})
}
