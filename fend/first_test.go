package fend

import (
	"testing"

	"github.com/foomo/fender/rule"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestFirstError(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := First(String("foo", rule.RequiredString))
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		err := First(String("", rule.RequiredString))
		assert.Error(t, err)
		assert.True(t, errors.Is(err, rule.Err))
		assert.True(t, errors.Is(errors.Cause(err), rule.ErrRequired))
		assert.EqualError(t, err, "required")
	})

	t.Run("second error", func(t *testing.T) {
		err := First(String("foo", rule.RequiredString, rule.MinString(10)))
		assert.Error(t, err)
		assert.True(t, errors.Is(err, rule.Err))
		assert.True(t, errors.Is(errors.Cause(err), rule.ErrMin))
		assert.EqualError(t, err, "min|10")
	})

	t.Run("third error", func(t *testing.T) {
		err := First(
			String("foo", rule.RequiredString),
			String("bar", rule.MinString(10)),
		)
		assert.Error(t, err)
		assert.True(t, errors.Is(err, rule.Err))
		assert.True(t, errors.Is(errors.Cause(err), rule.ErrMin))
		assert.EqualError(t, err, "min|10")
	})
}
