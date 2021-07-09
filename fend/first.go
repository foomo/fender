package fend

import "github.com/foomo/fender/rule"

func First(values ...Fend) *rule.Error {
	for _, value := range values {
		for _, v := range value() {
			if err := v(); err != nil {
				return err
			}
		}
	}
	return nil
}
