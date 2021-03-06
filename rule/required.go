package rule

const NameRequired = "required"

var ErrRequired = &Error{Rule: NameRequired}

// NewRequiredError constructor
func NewRequiredError() *Error {
	return NewError(NameRequired)
}

func RequiredInt(v int) Rule {
	return func() (*Error, error) {
		if v == 0 {
			return NewRequiredError(), nil
		}
		return nil, nil
	}
}

func RequiredInt8(v int8) Rule {
	return func() (*Error, error) {
		if v == 0 {
			return NewRequiredError(), nil
		}
		return nil, nil
	}
}

func RequiredInt32(v int32) Rule {
	return func() (*Error, error) {
		if v == 0 {
			return NewRequiredError(), nil
		}
		return nil, nil
	}
}

func RequiredInt64(v int64) Rule {
	return func() (*Error, error) {
		if v == 0 {
			return NewRequiredError(), nil
		}
		return nil, nil
	}
}

func RequiredUInt(v uint) Rule {
	return func() (*Error, error) {
		if v == 0 {
			return NewRequiredError(), nil
		}
		return nil, nil
	}
}

func RequiredUInt8(v uint8) Rule {
	return func() (*Error, error) {
		if v == 0 {
			return NewRequiredError(), nil
		}
		return nil, nil
	}
}

func RequiredUInt32(v uint32) Rule {
	return func() (*Error, error) {
		if v == 0 {
			return NewRequiredError(), nil
		}
		return nil, nil
	}
}

func RequiredUInt64(v uint64) Rule {
	return func() (*Error, error) {
		if v == 0 {
			return NewRequiredError(), nil
		}
		return nil, nil
	}
}

func RequiredFloat32(v float32) Rule {
	return func() (*Error, error) {
		if v == 0 {
			return NewRequiredError(), nil
		}
		return nil, nil
	}
}

func RequiredFloat64(v float64) Rule {
	return func() (*Error, error) {
		if v == 0 {
			return NewRequiredError(), nil
		}
		return nil, nil
	}
}

func RequiredString(v string) Rule {
	return func() (*Error, error) {
		if len(v) == 0 {
			return NewRequiredError(), nil
		}
		return nil, nil
	}
}
