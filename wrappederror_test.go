package fender_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
)

func TestWrappedError(t *testing.T) {
	t.Run("errors", func(t *testing.T) {
		if fendErr, err := fender.First(
			fender.Field("foo", fend.String("", rule.MinString(10))),
		); assert.NoError(t, err) {
			e := errors.New("validation failed")
			werr := fender.WrapError(e, fendErr)
			assert.True(t, errors.Is(werr, e))
			assert.NotNil(t, errors.Is(errors.Unwrap(werr), e))

			cerr := fender.UnwrapError(werr)
			assert.Error(t, cerr.First())
			assert.True(t, errors.Is(cerr.First(), rule.ErrMin))
			assert.NotEmpty(t, cerr.Error())
			assert.NotNil(t, cerr.Errors())

			cerr = fender.UnwrapError(e)
			assert.Nil(t, cerr)
		}
	})
}
