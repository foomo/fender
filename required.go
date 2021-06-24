package fender

const RuleRequired Rule = "required"

type RequiredError struct {
	Type Rule `json:"type"`
}

// NewFendRequiredError constructor
func NewFendRequiredError() *RuleError {
	return NewFendError(RuleRequired, string(RuleRequired))
}

func RequiredInt(v int) error {
	if v == 0 {
		return NewFendRequiredError()
	}
	return nil
}

func RequiredInt8(v int8) error {
	if v == 0 {
		return NewFendRequiredError()
	}
	return nil
}

func RequiredInt32(v int32) error {
	if v == 0 {
		return NewFendRequiredError()
	}
	return nil
}

func RequiredInt64(v int64) error {
	if v == 0 {
		return NewFendRequiredError()
	}
	return nil
}

func RequiredFloat32(v float32) error {
	if v == 0 {
		return NewFendRequiredError()
	}
	return nil
}

func RequiredFloat64(v float64) error {
	if v == 0 {
		return NewFendRequiredError()
	}
	return nil
}

func RequiredString(v string) error {
	if v == "" {
		return NewFendRequiredError()
	}
	return nil
}

func RequiredStrings(v []string) error {
	if len(v) == 0 {
		return NewFendRequiredError()
	}
	return nil
}
