package fender_test

import (
	"testing"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		if fendErr, err := fender.All(
			fender.Field("foo", fend.String("foo", rule.RequiredString, rule.MinString(1))),
			fender.Field("bar", fend.String("foo", rule.RequiredString, rule.MinString(1))),
		); assert.NoError(t, err) {
			assert.Nil(t, fendErr)
		}
	})

	t.Run("errors", func(t *testing.T) {
		if fendErr, err := fender.All(
			fender.Field("foo", fend.String("", rule.RequiredString, rule.MinString(10))),
			fender.Field("bar", fend.String("bar", rule.RequiredString, rule.MinString(10))),
		); assert.NoError(t, err) && assert.NotNil(t, fendErr) {
			assert.True(t, errors.Is(fendErr, fender.Err))
			errs := fendErr.Errors()
			assert.Len(t, errs, 2)
			assert.EqualError(t, fendErr, "foo:required;bar:min=10")
		}
	})

	t.Run("errors combined", func(t *testing.T) {
		if fendErr, err := fender.All(
			fender.Field("foo", fend.String("", rule.RequiredString, rule.MinString(10))),
			fender.Field("foo", fend.String("", rule.MinString(10), rule.RequiredString)),
		); assert.NoError(t, err) && assert.NotNil(t, fendErr) {
			assert.True(t, errors.Is(fendErr, fender.Err))
			errs := fendErr.Errors()
			assert.Len(t, errs, 1)
			assert.EqualError(t, fendErr, "foo:required")
		}
	})

	t.Run("return std error", func(t *testing.T) {
		e := errors.New("std error")
		if fendErr, err := fender.All(
			fender.Field("foo", fend.Custom(func() (*rule.Error, error) {
				return nil, e
			})),
		); assert.Error(t, err) && assert.Nil(t, fendErr) {
			assert.True(t, errors.Is(err, e))
		}
	})

	t.Run("return std errors", func(t *testing.T) {
		e := errors.New("std error")
		if fendErr, err := fender.All(
			fender.Field("one", fend.String("", rule.RequiredString)),
			fender.Field("two", fend.Custom(func() (*rule.Error, error) {
				return nil, e
			})),
			fender.Field("three", fend.String("", rule.RequiredString)),
		); assert.Error(t, err) && assert.Nil(t, fendErr) {
			assert.True(t, errors.Is(err, e))
		}
	})
}
