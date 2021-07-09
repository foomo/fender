package fend

import (
	"testing"

	"github.com/foomo/fender/rule"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		errs := All(String("foo", rule.RequiredString(), rule.MinString(2)))
		assert.Empty(t, errs)
	})

	t.Run("errors", func(t *testing.T) {
		errs := All(String("", rule.RequiredString(), rule.MinString(2)))
		assert.Len(t, errs, 2)
	})

}
