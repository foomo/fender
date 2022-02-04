package rule

import "errors"

const NameValid Name = "valid"

type Validator interface {
	Valid() bool
}

var ErrValid = errors.New(NameValid.String())

// NewValidError constructor
func NewValidError() *Error {
	return NewError(ErrValid, NameValid.String())
}

func Valid() ValidatorRule {
	return func(v Validator) Rule {
		return func() (*Error, error) {
			if !v.Valid() {
				return NewValidError(), nil
			}
			return nil, nil
		}
	}
}
