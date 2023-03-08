package rule

// NewCustomRuleError constructor
func NewCustomRuleError(rule Name, meta ...string) *Error {
	var m []string
	for _, v := range meta {
		if v != "" {
			m = append(m, v)
		}
	}
	return NewError(rule, m...)
}
