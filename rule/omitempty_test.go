package rule_test

import (
	"context"
	"testing"

	"github.com/foomo/fender/rule"
)

func TestOmitEmptyString(t *testing.T) {
	tests := map[string]bool{
		"":  false,
		"v": true,
	}
	t.Run("OmitEmptyString", func(t *testing.T) {
		for value, valid := range tests {
			if err := rule.OmitEmptyString(context.TODO(), value); (err == nil) != valid {
				t.Errorf("OmitEmptyString() error = %v, wantErr %v", err, !valid)
			}
		}
	})
}
