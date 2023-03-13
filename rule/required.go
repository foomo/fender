package rule

// NewRequiredError constructor
func NewRequiredError() *Error {
	return NewError(NameRequired)
}
