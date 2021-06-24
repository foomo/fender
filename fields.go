package fender

// Fields type
type Fields struct {
	errors map[string][]error
}

// NewFields constructor
func NewFields() *Fields {
	return &Fields{
		errors: map[string][]error{},
	}
}

// Add adds an error fields if a validation fails
func (f Fields) Add(name string, errs ...error) Fields {
	for _, err := range errs {
		if err != nil {
			f.errors[name] = append(f.errors[name], err)
		}
	}
	return f
}

// All return all fields errors
func (f Fields) All() map[string][]error {
	return f.errors
}

// First returns only all fields first error
func (f Fields) First() map[string]error {
	m := make(map[string]error, len(f.errors))
	for k, v := range f.errors {
		m[k] = v[0]
	}
	return m
}

// AllError returns all fields errors as error type
func (f Fields) AllError() error {
	if f.HasErrors() {
		return NewFieldsErrors(f.All())
	}
	return nil
}

// FirstError returns only all fields first errors as error type
func (f Fields) FirstError() error {
	if f.HasErrors() {
		return NewFieldsError(f.First())
	}
	return nil
}

// HasErrors return true if there are no errors
func (f Fields) HasErrors() bool {
	return len(f.errors) > 0
}
