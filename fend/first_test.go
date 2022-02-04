package fend

import (
	"testing"

	"github.com/foomo/fender/rule"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestFirstError(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		if fendErr, err := First(String("foo", rule.RequiredString)); assert.NoError(t, err) {
			assert.Nil(t, fendErr)
		}
	})

	t.Run("error", func(t *testing.T) {
		if fendErr, err := First(String("", rule.RequiredString)); assert.NoError(t, err) {
			assert.True(t, errors.Is(fendErr, rule.Err))
			assert.True(t, errors.Is(errors.Cause(fendErr), rule.ErrRequired))
			assert.EqualError(t, fendErr, "required")
		}
	})

	t.Run("second error", func(t *testing.T) {
		if fendErr, err := First(String("foo", rule.RequiredString, rule.MinString(10))); assert.NoError(t, err) {
			assert.True(t, errors.Is(fendErr, rule.Err))
			assert.True(t, errors.Is(errors.Cause(fendErr), rule.ErrMin))
			assert.EqualError(t, fendErr, "min=10")
		}
	})

	t.Run("third error", func(t *testing.T) {
		if fendErr, err := First(
			String("foo", rule.RequiredString),
			String("bar", rule.MinString(10)),
		); assert.NoError(t, err) {
			assert.True(t, errors.Is(fendErr, rule.Err))
			assert.True(t, errors.Is(errors.Cause(fendErr), rule.ErrMin))
			assert.EqualError(t, fendErr, "min=10")
		}
	})
}
