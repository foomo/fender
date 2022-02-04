package rule

const NameValid = "valid"

type Validator interface {
	Valid() bool
}

var ErrValid = &Error{Rule: NameValid}

// NewValidError constructor
func NewValidError() *Error {
	return NewError(NameValid)
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
