package fender

const RuleEnum Rule = "enum"

type Enumerator interface {
	Valid() bool
}

// NewFendEnumError constructor
func NewFendEnumError() *RuleError {
	return NewRuleError(RuleEnum)
}

func Enum(v Enumerator) Fend {
	return func() *RuleError {
		if !v.Valid() {
			return NewFendEnumError()
		}
		return nil
	}
}
