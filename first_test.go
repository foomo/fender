package fender_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
)

func TestFirst(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		if fendErr, err := fender.First(
			fender.Field("foo", fend.String("foo", rule.RequiredString)),
		); assert.NoError(t, err) {
			assert.Nil(t, fendErr)
		}
	})

	t.Run("error string", func(t *testing.T) {
		if fendErr, err := fender.First(
			fender.Field("foo", fend.String("", rule.MinString(10))),
		); assert.NoError(t, err) && assert.NotNil(t, fendErr) {
			assert.True(t, errors.Is(fendErr, fender.Err))
			first := fendErr.First()
			assert.True(t, errors.Is(first, rule.Err))
			assert.True(t, errors.Is(first, rule.ErrMin))
			assert.False(t, errors.Is(first, rule.ErrMax))
			assert.EqualError(t, fendErr, "foo:min=10")
		}
	})

	t.Run("error var", func(t *testing.T) {
		if fendErr, err := fender.First(
			fender.Field("foo", fend.Var("", "min=10")),
		); assert.NoError(t, err) && assert.NotNil(t, fendErr) {
			first := fendErr.First()
			assert.True(t, errors.Is(first, rule.Err))
			assert.True(t, errors.Is(first, rule.ErrVar))
			assert.EqualError(t, fendErr, "foo:min=10")
		}
	})

	t.Run("return std error", func(t *testing.T) {
		e := errors.New("std error")
		if fendErr, err := fender.First(fender.Field("foo", fend.Custom(func() (*rule.Error, error) {
			return nil, e
		}))); assert.Error(t, err) && assert.Nil(t, fendErr) {
			assert.True(t, errors.Is(err, e))
		}
	})

	t.Run("return std errors", func(t *testing.T) {
		e := errors.New("std error")
		if fendErr, err := fender.First(
			fender.Field("foo", fend.Custom(func() (*rule.Error, error) {
				return nil, e
			})),
			fender.Field("foo", fend.String("", rule.RequiredString)),
		); assert.Error(t, err) && assert.Nil(t, fendErr) {
			assert.True(t, errors.Is(err, e))
		}
	})
}
