package fend

import (
	"github.com/foomo/fender/rule"
)

func All(values ...Fend) ([]*rule.Error, error) {
	var errs []*rule.Error
	for _, value := range values {
		for _, v := range value() {
			if ruleErr, err := v(); err != nil {
				return nil, err
			} else if ruleErr != nil {
				errs = append(errs, ruleErr)
			}
		}
	}
	return errs, nil
}
