package fender_test

import (
	"testing"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestWrappedError(t *testing.T) {
	t.Run("errors", func(t *testing.T) {
		if fendErr, err := fender.First(
			fender.Field("foo", fend.String("", rule.MinString(10))),
		); assert.NoError(t, err) {
			e := errors.New("validation failed")
			werr := fender.NewWrappedError(e, fendErr)
			assert.True(t, errors.Is(werr, e))
			assert.NotNil(t, errors.Is(errors.Unwrap(werr), e))

			cerr := fender.WrappedErrorCause(werr)
			assert.Error(t, cerr.First())
			assert.NotEmpty(t, cerr.Error())
			assert.NotNil(t, cerr.Errors())

			cerr = fender.WrappedErrorCause(e)
			assert.NoError(t, cerr.First())
			assert.Empty(t, cerr.Error())
			assert.Nil(t, cerr.Errors())
		}
	})
}
