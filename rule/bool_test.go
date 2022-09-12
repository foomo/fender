package rule_test

import (
	"testing"

	"github.com/foomo/fender/rule"
)

func TestTrue(t *testing.T) {
	tests := map[bool]bool{
		true:  true,
		false: false,
	}
	t.Run("True", func(t *testing.T) {
		for value, valid := range tests {
			if ruleErr, _ := rule.True(value)(); (ruleErr == nil) != valid {
				t.Errorf("True() error = %v, wantErr %v", ruleErr, !valid)
			}
		}
	})
}

func TestFalse(t *testing.T) {
	tests := map[bool]bool{
		true:  false,
		false: true,
	}
	t.Run("False", func(t *testing.T) {
		for value, valid := range tests {
			if ruleErr, _ := rule.False(value)(); (ruleErr == nil) != valid {
				t.Errorf("False() error = %v, wantErr %v", ruleErr, !valid)
			}
		}
	})
}
