package fender

const RuleRequired Rule = "required"

type RequiredError struct {
	Type Rule `json:"type"`
}

// NewFendRequiredError constructor
func NewFendRequiredError() *RuleError {
	return NewRuleError(RuleRequired)
}

func RequiredInt(v int) Fend {
	return func() *RuleError {
		if v == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredInt8(v int8) Fend {
	return func() *RuleError {
		if v == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredInt32(v int32) Fend {
	return func() *RuleError {
		if v == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredInt64(v int64) Fend {
	return func() *RuleError {
		if v == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredUInt(v uint) Fend {
	return func() *RuleError {
		if v == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredUInt8(v uint8) Fend {
	return func() *RuleError {
		if v == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredUInt32(v uint32) Fend {
	return func() *RuleError {
		if v == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredUInt64(v uint64) Fend {
	return func() *RuleError {
		if v == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredFloat32(v float32) Fend {
	return func() *RuleError {
		if v == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredFloat64(v float64) Fend {
	return func() *RuleError {
		if v == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredString(v string) Fend {
	return func() *RuleError {
		if v == "" {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredInts(v []int) Fend {
	return func() *RuleError {
		if len(v) == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredInt8s(v []int8) Fend {
	return func() *RuleError {
		if len(v) == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredInt32s(v []int32) Fend {
	return func() *RuleError {
		if len(v) == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredInt64s(v []int64) Fend {
	return func() *RuleError {
		if len(v) == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredUInts(v []uint) Fend {
	return func() *RuleError {
		if len(v) == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredUInt8s(v []uint8) Fend {
	return func() *RuleError {
		if len(v) == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredUInt32s(v []uint32) Fend {
	return func() *RuleError {
		if len(v) == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredUInt64s(v []uint64) Fend {
	return func() *RuleError {
		if len(v) == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredFloat32s(v []float32) Fend {
	return func() *RuleError {
		if len(v) == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredFloat64s(v []float64) Fend {
	return func() *RuleError {
		if len(v) == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}

func RequiredStrings(v []string) Fend {
	return func() *RuleError {
		if len(v) == 0 {
			return NewFendRequiredError()
		}
		return nil
	}
}
