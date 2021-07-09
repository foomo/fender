package fender_test

import (
	"testing"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)



func TestFirst(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		err := fender.First(fender.Field("foo", fend.String("foo", rule.RequiredString())))
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		err := fender.First(fender.Field("foo", fend.String("", rule.RequiredString())))
		assert.Error(t, err)
		assert.True(t, errors.Is(err, fender.Err))
		assert.True(t, errors.Is(errors.Unwrap(err), rule.Err))
		assert.True(t, errors.Is(errors.Cause(err), rule.ErrRequired))
		assert.EqualError(t, err, "foo:required")
	})

	t.Run("error meta", func(t *testing.T) {
		err := fender.First(fender.Field("foo", fend.String("", rule.MinString(10))))
		assert.Error(t, err)
		assert.True(t, errors.Is(err, fender.Err))
		assert.True(t, errors.Is(errors.Unwrap(err), rule.Err))
		assert.True(t, errors.Is(errors.Cause(err), rule.ErrMin))
		assert.EqualError(t, err, "foo:min|10")
	})
}

