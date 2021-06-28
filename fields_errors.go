package fender

// FieldsErrors type
type FieldsErrors struct {
	Map FieldsErrorsMap
}

// NewFieldsErrors constructor
func NewFieldsErrors(m FieldsErrorsMap) *FieldsErrors {
	return &FieldsErrors{
		Map: m,
	}
}

// Error interface
func (e *FieldsErrors) Error() string {
	return e.Map.String()
}
