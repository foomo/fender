package fender

const RuleEnum Rule = "enum"

type Enumerator interface {
	Valid() bool
}

// NewFendEnumError constructor
func NewFendEnumError() *RuleError {
	return NewFendError(RuleEnum, string(RuleEnum))
}

func Enum(v Enumerator) error {
	if !v.Valid() {
		return NewFendEnumError()
	}
	return nil
}
