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
			// assert.True(t, errors.Is(fendErr, fender.Err)) FIXME
			// first := fendErr.First() FIXME
			// assert.True(t, errors.Is(first, rule.ErrMin)) FIXME
			// assert.True(t, errors.Is(errors.Cause(first), rule.ErrMin)) FIXME
			assert.EqualError(t, fendErr, "foo:min=10")
		}
	})

	t.Run("error var", func(t *testing.T) {
		if fendErr, err := fender.First(fender.Field("foo", fend.Var("", "min=10"))); assert.NoError(t, err) && assert.NotNil(t, fendErr) {
			assert.True(t, fender.IsError(fendErr))
			// first := fendErr.First()
			// assert.True(t, errors.Is(first, rule.Err)) FIXME
			// assert.True(t, errors.Is(errors.Cause(first), rule.ErrVar)) FIXME
			assert.EqualError(t, fendErr, "foo:min=10")
		}
	})

	t.Run("return std error", func(t *testing.T) {
		e := errors.New("std error")
		if fendErr, err := fender.First(fender.Field("foo", fend.Custom(func() (*rule.Error, error) {
			return nil, e
		}))); assert.Error(t, err) && assert.Nil(t, fendErr) {
			assert.False(t, fender.IsError(fendErr))
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
			assert.False(t, fender.IsError(fendErr))
			assert.True(t, errors.Is(err, e))
		}
	})
}
