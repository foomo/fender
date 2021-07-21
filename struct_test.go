package fender_test

import (
	"testing"

	"github.com/foomo/fender"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestStruct(t *testing.T) {
	type test struct {
		String string `validate:"required"`
		Int    int    `validate:"min=10"`
	}

	t.Run("success", func(t *testing.T) {
		inst := test{
			String: "foo",
			Int:    13,
		}
		err := fender.Struct(inst)
		assert.NoError(t, err)
	})

	t.Run("errors", func(t *testing.T) {
		err := fender.Struct(test{})
		assert.Error(t, err)
		assert.True(t, errors.Is(err, fender.Err))
		errs := err.(*fender.Error).Errors()
		assert.Len(t, errs, 2)
		assert.EqualError(t, err, "string:required;int:min=10")
	})
}
