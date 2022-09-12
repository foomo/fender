package fend_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
)

func TestAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		if ruleErrs, err := fend.All(
			fend.String("foo", rule.RequiredString, rule.MinString(2)),
		); assert.NoError(t, err) {
			assert.Empty(t, ruleErrs)
		}
	})

	t.Run("errors", func(t *testing.T) {
		if ruleErrs, err := fend.All(
			fend.String("", rule.RequiredString, rule.MinString(2)),
		); assert.NoError(t, err) {
			assert.Len(t, ruleErrs, 2)
		}
	})
}
