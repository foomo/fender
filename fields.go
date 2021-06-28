package fender

// Fields type
type Fields struct {
	fieldFends map[string][]Fend
}

// NewFields constructor
func NewFields() *Fields {
	return &Fields{
		fieldFends: map[string][]Fend{},
	}
}

// Add adds an error fields if a validation fails
func (f Fields) Add(name string, fends ...Fend) Fields {
	f.fieldFends[name] = append(f.fieldFends[name], fends...)
	return f
}

// AllFieldsErrors return all fields errors
func (f Fields) AllFieldsErrors() FieldsErrorsMap {
	errs := FieldsErrorsMap{}
	for field, fends := range f.fieldFends {
		for _, fend := range fends {
			if err := fend(); err != nil {
				errs[field] = append(errs[field], err)
			}
		}
	}
	return nil
}

// AllFieldsFirstError returns only the first error on all fields
func (f Fields) AllFieldsFirstError() FieldsErrorMap {
	errs := FieldsErrorMap{}
	for field, fends := range f.fieldFends {
		for _, fend := range fends {
			if err := fend(); err != nil {
				errs[field] = err
				break
			}
		}
	}
	return errs
}

// FirstFieldError returns only the first error
func (f Fields) FirstFieldError() error {
	for field, fends := range f.fieldFends {
		for _, fend := range fends {
			if err := fend(); err != nil {
				return NewFieldError(field, err)
			}
		}
	}
	return nil
}
