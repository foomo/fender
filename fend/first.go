package fend

import (
	"github.com/foomo/fender/rule"
)

func First(values ...Fend) (*rule.Error, error) {
	for _, value := range values {
		for _, v := range value() {
			if ruleErr, err := v(); err != nil {
				return nil, err
			} else if ruleErr != nil {
				return ruleErr, nil
			}
		}
	}
	return nil, nil
}
