package rule

import "errors"

const NameValid Name = "valid"

type Validator interface {
	Valid() bool
}

var ErrValid = errors.New(NameValid.String())

// NewValidError constructor
func NewValidError() *Error {
	return NewError(ErrValid)
}

func Valid() ValidatorRule {
	return func(v Validator) Rule {
		return func() *Error {
			if !v.Valid() {
				return NewValidError()
			}
			return nil
		}
	}
}
