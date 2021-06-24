package fender

type RuleError struct {
	Rule Rule
	Meta string
}

// NewFendError constructor
func NewFendError(fend Rule, meta string) *RuleError {
	return &RuleError{
		Rule: fend,
		Meta: meta,
	}
}

// RuleError interface
func (e *RuleError) Error() string {
	s := string(e.Rule)
	if e.Meta != "" {
		s += MetaDelimiter + e.Meta
	}
	return s
}
