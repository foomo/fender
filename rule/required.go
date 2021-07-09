package rule

import (
	"errors"
)

const NameRequired Name = "required"

var ErrRequired = errors.New(NameRequired.String())

// NewRequiredError constructor
func NewRequiredError() *Error {
	return NewError(ErrRequired)
}

func RequiredInt() IntRule {
	return func(v int) Rule {
		return func() *Error {
			if v == 0 {
				return NewRequiredError()
			}
			return nil
		}
	}
}

func RequiredInt8() Int8Rule {
	return func(v int8) Rule {
		return func() *Error {
			if v == 0 {
				return NewRequiredError()
			}
			return nil
		}
	}
}

func RequiredInt32() Int32Rule {
	return func(v int32) Rule {
		return func() *Error {
			if v == 0 {
				return NewRequiredError()
			}
			return nil
		}
	}
}

func RequiredInt64() Int64Rule {
	return func(v int64) Rule {
		return func() *Error {
			if v == 0 {
				return NewRequiredError()
			}
			return nil
		}
	}
}

func RequiredUInt() UIntRule {
	return func(v uint) Rule {
		return func() *Error {
			if v == 0 {
				return NewRequiredError()
			}
			return nil
		}
	}
}

func RequiredUInt8() UInt8Rule {
	return func(v uint8) Rule {
		return func() *Error {
			if v == 0 {
				return NewRequiredError()
			}
			return nil
		}
	}
}

func RequiredUInt32() UInt32Rule {
	return func(v uint32) Rule {
		return func() *Error {
			if v == 0 {
				return NewRequiredError()
			}
			return nil
		}
	}
}

func RequiredUInt64() UInt64Rule {
	return func(v uint64) Rule {
		return func() *Error {
			if v == 0 {
				return NewRequiredError()
			}
			return nil
		}
	}
}

func RequiredFloat32() Float32Rule {
	return func(v float32) Rule {
		return func() *Error {
			if v == 0 {
				return NewRequiredError()
			}
			return nil
		}
	}
}

func RequiredFloat64() Float64Rule {
	return func(v float64) Rule {
		return func() *Error {
			if v == 0 {
				return NewRequiredError()
			}
			return nil
		}
	}
}

func RequiredString() StringRule {
	return func(v string) Rule {
		return func() *Error {
			if v == "" {
				return NewRequiredError()
			}
			return nil
		}
	}
}
