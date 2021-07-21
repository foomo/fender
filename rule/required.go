package rule

import (
	"errors"
)

const NameRequired Name = "required"

var ErrRequired = errors.New(NameRequired.String())

// NewRequiredError constructor
func NewRequiredError() *Error {
	return NewError(ErrRequired, NameRequired.String())
}

func RequiredInt(v int) Rule {
	return func() *Error {
		if v == 0 {
			return NewRequiredError()
		}
		return nil
	}
}

func RequiredInt8(v int8) Rule {
	return func() *Error {
		if v == 0 {
			return NewRequiredError()
		}
		return nil
	}
}

func RequiredInt32(v int32) Rule {
	return func() *Error {
		if v == 0 {
			return NewRequiredError()
		}
		return nil
	}
}

func RequiredInt64(v int64) Rule {
	return func() *Error {
		if v == 0 {
			return NewRequiredError()
		}
		return nil
	}
}

func RequiredUInt(v uint) Rule {
	return func() *Error {
		if v == 0 {
			return NewRequiredError()
		}
		return nil
	}
}

func RequiredUInt8(v uint8) Rule {
	return func() *Error {
		if v == 0 {
			return NewRequiredError()
		}
		return nil
	}
}

func RequiredUInt32(v uint32) Rule {
	return func() *Error {
		if v == 0 {
			return NewRequiredError()
		}
		return nil
	}
}

func RequiredUInt64(v uint64) Rule {
	return func() *Error {
		if v == 0 {
			return NewRequiredError()
		}
		return nil
	}
}

func RequiredFloat32(v float32) Rule {
	return func() *Error {
		if v == 0 {
			return NewRequiredError()
		}
		return nil
	}
}

func RequiredFloat64(v float64) Rule {
	return func() *Error {
		if v == 0 {
			return NewRequiredError()
		}
		return nil
	}
}

func RequiredString(v string) Rule {
	return func() *Error {
		if len(v) == 0 {
			return NewRequiredError()
		}
		return nil
	}
}
