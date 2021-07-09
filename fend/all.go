package fend

import "github.com/foomo/fender/rule"

func All(values ...Fend) []*rule.Error {
	var errs []*rule.Error
	for _, value := range values {
		for _, v := range value() {
			if err := v(); err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errs
}
