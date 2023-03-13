package fender_test

import (
	"context"
	"errors"
	"testing"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := fender.All(
			context.TODO(),
			fend.Field("foo", "foo", rule.StringRequired, rule.StringMin(1)),
			fend.Field("bar", "foo", rule.StringRequired, rule.StringMin(1)),
		)
		assert.NoError(t, err)
	})

	t.Run("errors", func(t *testing.T) {
		err := fender.All(
			context.TODO(),
			fend.Field("foo", "", rule.StringRequired, rule.StringMin(10)),
			fend.Field("bar", "bar", rule.StringRequired, rule.StringMin(10)),
		)
		if fendErr := fender.AsError(err); assert.NotNil(t, fendErr) {
			errs := fendErr.Errors()
			assert.Len(t, errs, 2)
		}
		assert.EqualError(t, err, "foo:required:min=10;bar:min=10")
	})

	t.Run("errors combined", func(t *testing.T) {
		err := fender.All(
			context.TODO(),
			fend.Field("foo", "", rule.StringRequired, rule.StringMin(10)),
			fend.Field("foo", "", rule.StringMin(10), rule.StringRequired),
		)
		if fendErr := fender.AsError(err); assert.NotNil(t, fendErr) {
			errs := fendErr.Errors()
			assert.Len(t, errs, 2)
		}
		assert.EqualError(t, err, "foo:required:min=10;foo:min=10:required")
	})

	t.Run("return std error", func(t *testing.T) {
		e := errors.New("std error")
		err := fender.All(
			context.TODO(),
			fend.Field("one", "", rule.StringRequired),
			fend.Field("foo", "", func(ctx context.Context, v string) error {
				return e
			}),
		)
		assert.Nil(t, fender.AsError(err))
		assert.EqualError(t, err, e.Error())
	})
}

func TestFirst(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := fender.First(
			context.TODO(),
			fend.Field("foo", "", rule.StringRequired),
		)
		if fendErr := fender.AsError(err); assert.NotNil(t, fendErr) {
			first := fender.AsError(err).First()
			assert.NotNil(t, rule.AsError(first))
		}
		assert.EqualError(t, err, "foo:required")
	})

	t.Run("error string", func(t *testing.T) {
		err := fender.First(
			context.TODO(),
			fend.Field("foo", "", rule.StringMin(10)),
		)
		if fendErr := fender.AsError(err); assert.NotNil(t, fendErr) {
			first := fender.AsError(err).First()
			assert.NotNil(t, rule.AsError(first))
		}
		assert.EqualError(t, err, "foo:min=10")
	})

	t.Run("error var", func(t *testing.T) {
		err := fender.First(
			context.TODO(),
			fend.Field("foo", "", rule.Var("min=10")),
		)
		if fendErr := fender.AsError(err); assert.NotNil(t, fendErr) {
			first := fender.AsError(err).First()
			assert.NotNil(t, rule.AsError(first))
		}
		assert.EqualError(t, err, "foo:min=10")
	})

	t.Run("return std error", func(t *testing.T) {
		e := errors.New("std error")
		err := fender.First(
			context.TODO(),
			fend.Field("foo", "", func(ctx context.Context, v string) error {
				return e
			}),
		)
		assert.Nil(t, fender.AsError(err))
		assert.EqualError(t, err, e.Error())
	})
}
