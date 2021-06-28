package fender

// FieldsError type
type FieldsError struct {
	Map FieldsErrorMap
}

// NewFieldsError constructor
func NewFieldsError(m FieldsErrorMap) *FieldsError {
	return &FieldsError{
		Map: m,
	}
}

// Error interface
func (e *FieldsError) Error() string {
	return e.Map.String()
}
