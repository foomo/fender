package config

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

func init() {
	Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return strings.ToLower(fld.Name[:1]) + fld.Name[1:]
	})
}
