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
		err := fender.First(fender.Field("foo", fend.String("foo", rule.RequiredString)))
		assert.NoError(t, err)
	})

	t.Run("error string", func(t *testing.T) {
		err := fender.First(fender.Field("foo", fend.String("", rule.MinString(10))))
		assert.Error(t, err)
		assert.True(t, errors.Is(err, fender.Err))
		first := err.(*fender.Error).First()
		assert.True(t, errors.Is(first, rule.Err))
		assert.True(t, errors.Is(errors.Cause(first), rule.ErrMin))
		assert.EqualError(t, err, "foo:min=10")
	})

	t.Run("error var", func(t *testing.T) {
		err := fender.First(fender.Field("foo", fend.Var("", "min=10")))
		assert.Error(t, err)
		assert.True(t, errors.Is(err, fender.Err))
		first := err.(*fender.Error).First()
		assert.True(t, errors.Is(first, rule.Err))
		assert.True(t, errors.Is(errors.Cause(first), rule.ErrVar))
		assert.EqualError(t, err, "foo:min=10")
	})
}
