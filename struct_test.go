package fender_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/foomo/fender"
)

func TestStruct(t *testing.T) {
	type test struct {
		Name string `validate:"required"`
		Age  int    `validate:"min=10"`
	}

	t.Run("success", func(t *testing.T) {
		inst := test{
			Name: "foo",
			Age:  13,
		}
		if fendErr, err := fender.Struct(inst); assert.NoError(t, err) {
			assert.Nil(t, fendErr)
		}
	})

	t.Run("errors", func(t *testing.T) {
		if fendErr, err := fender.Struct(test{}); assert.NoError(t, err) {
			assert.True(t, errors.Is(fendErr, fender.Err))
			errs := fendErr.Errors()
			assert.Len(t, errs, 2)
			assert.Equal(t, len("name:required;age:min=10"), len(fendErr.Error()))
		}
	})
}
