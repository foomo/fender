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
		err := fender.All(
			fender.Field("foo", fend.String("foo", rule.RequiredString(), rule.MinString(1))),
			fender.Field("bar", fend.String("foo", rule.RequiredString(), rule.MinString(1))),
		)
		assert.NoError(t, err)
	})

	t.Run("errors", func(t *testing.T) {
		err := fender.All(
			fender.Field("foo", fend.String("", rule.RequiredString(), rule.MinString(10))),
			fender.Field("bar", fend.String("", rule.RequiredString(), rule.MinString(10))),
		)
		assert.Error(t, err)
		assert.True(t, errors.Is(err, fender.Err))
		errs := err.(*fender.Error).Errors()
		assert.Len(t, errs, 2)
		assert.Len(t, errs["foo"], 2)
		assert.Len(t, errs["bar"], 2)
		assert.EqualError(t, err, "foo:required,min|10;bar:required,min|10")
	})

	t.Run("errors combined", func(t *testing.T) {
		err := fender.All(
			fender.Field("foo", fend.String("", rule.RequiredString(), rule.MinString(10))),
			fender.Field("foo", fend.String("", rule.MinString(10), rule.RequiredString())),
		)
		assert.Error(t, err)
		assert.True(t, errors.Is(err, fender.Err))
		errs := err.(*fender.Error).Errors()
		assert.Len(t, errs, 1)
		assert.Len(t, errs["foo"], 4)
		assert.EqualError(t, err, "foo:required,min|10,min|10,required")
	})
}
