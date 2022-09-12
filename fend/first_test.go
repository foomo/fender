package fend_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
)

func TestFirstError(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		if fendErr, err := fend.First(fend.String("foo", rule.RequiredString)); assert.NoError(t, err) {
			assert.Nil(t, fendErr)
		}
	})
	t.Run("success", func(t *testing.T) {
		if fendErr, err := fend.First(fend.Bool(true, rule.True)); assert.NoError(t, err) {
			assert.Nil(t, fendErr)
		}
	})

	t.Run("error", func(t *testing.T) {
		if fendErr, err := fend.First(fend.String("", rule.RequiredString)); assert.NoError(t, err) {
			assert.True(t, errors.Is(fendErr, rule.Err))
			assert.True(t, errors.Is(errors.Cause(fendErr), rule.ErrRequired))
			assert.EqualError(t, fendErr, "required")
		}
	})

	t.Run("second error", func(t *testing.T) {
		if fendErr, err := fend.First(fend.String("foo", rule.RequiredString, rule.MinString(10))); assert.NoError(t, err) {
			assert.True(t, errors.Is(fendErr, rule.Err))
			assert.True(t, errors.Is(errors.Cause(fendErr), rule.ErrMin))
			assert.EqualError(t, fendErr, "min=10")
		}
	})

	t.Run("third error", func(t *testing.T) {
		if fendErr, err := fend.First(
			fend.String("foo", rule.RequiredString),
			fend.String("bar", rule.MinString(10)),
		); assert.NoError(t, err) {
			assert.True(t, errors.Is(fendErr, rule.Err))
			assert.True(t, errors.Is(errors.Cause(fendErr), rule.ErrMin))
			assert.EqualError(t, fendErr, "min=10")
		}
	})
}
