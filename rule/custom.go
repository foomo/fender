package rule

const NameCustom = "custom"

// NewCustomRuleError constructor
func NewCustomRuleError(rule string, meta ...string) *Error {
	var m []string
	for _, v := range meta {
		if v != "" {
			m = append(m, v)
		}
	}
	return NewError(rule, m...)
}
