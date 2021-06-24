package fender_test

import (
	"testing"

	"github.com/foomo/fender"
	"github.com/go-playground/validator/v10"
)

// https://github.com/frederikhors/bench-golang-validators

type User struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Age       int    `json:"age" validate:"gte=0,lte=130"`
	Email     string `json:"email" validate:"required,email"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	Planet    string `json:"planet" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
}

var user = &User{
	FirstName: "Badger",
	LastName:  "Smith",
	Age:       80,
	City:      "string",
	Email:     "Badger.Smith@email.com",
	Street:    "Eavesdown Docks",
	Planet:    "Persphone",
	Phone:     "none",
}

func BenchmarkFender(b *testing.B) {
	u := &User{}
	b.Run("invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := fender.NewFields().
				Add("firstName", fender.RequiredString(u.FirstName), fender.MinString(u.FirstName, 100)).
				Add("lastName", fender.RequiredString(u.LastName)).
				Add("age", fender.MinInt(u.Age, 0), fender.MaxInt(u.Age, 130)).
				Add("email", fender.Email(u.Email, false)).
				Add("street", fender.RequiredString(u.Street)).
				Add("city", fender.RequiredString(u.City)).
				Add("planet", fender.RequiredString(u.Planet)).
				Add("phone", fender.RequiredString(u.Phone)).
				AllError(); err != nil {
				b.Log("Fender: " + err.Error())
			}
		}
	})
	u = user
	b.Run("valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := fender.NewFields().
				Add("firstName", fender.RequiredString(u.FirstName)).
				Add("lastName", fender.RequiredString(u.LastName)).
				Add("age", fender.MinInt(u.Age, 0), fender.MaxInt(u.Age, 130)).
				Add("email", fender.Email(u.Email, false)).
				Add("street", fender.RequiredString(u.Street)).
				Add("city", fender.RequiredString(u.City)).
				Add("planet", fender.RequiredString(u.Planet)).
				Add("phone", fender.RequiredString(u.Phone)).
				AllError(); err != nil {
				b.Log("Fender: " + err.Error())
			}
		}
	})
}

func BenchmarkPlayground(b *testing.B) {
	u := &User{}
	b.Run("invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := validator.New()
			if err := v.Struct(u); err != nil {
				b.Log("Playground: " + err.Error())
			}
		}
	})
	u = user
	b.Run("valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := validator.New()
			if err := v.Struct(u); err != nil {
				b.Log("Playground: " + err.Error())
			}
		}
	})
}
