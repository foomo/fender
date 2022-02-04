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
		if fendErr, err := fender.Struct(inst); assert.NoError(t, err) {
			assert.Nil(t, fendErr)
		}
	})

	t.Run("errors", func(t *testing.T) {
		if fendErr, err := fender.Struct(test{}); assert.NoError(t, err) {
			assert.True(t, errors.Is(fendErr, fender.Err))
			errs := fendErr.Errors()
			assert.Len(t, errs, 2)
			assert.Len(t, fendErr.Error(), len("string:required;int:min=10"))
		}
	})
}
