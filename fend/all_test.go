package fend

import (
	"testing"

	"github.com/foomo/fender/rule"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		if ruleErrs, err := All(
			String("foo", rule.RequiredString, rule.MinString(2)),
		); assert.NoError(t, err) {
			assert.Empty(t, ruleErrs)
		}
	})

	t.Run("errors", func(t *testing.T) {
		if ruleErrs, err := All(
			String("", rule.RequiredString, rule.MinString(2)),
		); assert.NoError(t, err) {
			assert.Len(t, ruleErrs, 2)
		}
	})
}
