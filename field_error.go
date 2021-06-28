package fender

// FieldError type
type FieldError struct {
	Field string
	Rule  Rule
	Meta  string
}

// NewFieldError constructor
func NewFieldError(field string, err *RuleError) *FieldError {
	return &FieldError{
		Field: field,
		Rule:  err.Rule,
		Meta:  err.Meta,
	}
}

// Error interface
func (e *FieldError) Error() string {
	s := e.Field + FieldDelimiter + string(e.Rule)
	if e.Meta != "" {
		s += MetaDelimiter + e.Meta
	}
	return s
}
