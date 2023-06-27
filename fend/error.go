package fend

import (
	"errors"
	"strings"

	"github.com/foomo/fender/config"
	"github.com/foomo/fender/rule"
)

type Error struct {
	Path     string
	RuleErrs []*rule.Error
}

func NewError(path string, ruleErrs ...*rule.Error) *Error {
	return &Error{
		Path:     path,
		RuleErrs: ruleErrs,
	}
}

func NewRuleError(path string, ruleName rule.Name, meta ...string) *Error {
	return &Error{
		Path:     path,
		RuleErrs: []*rule.Error{rule.NewError(ruleName, meta...)},
	}
}

func (e *Error) Name() string {
	return e.Path
}

func (e *Error) Error() string {
	var ret string
	causes := make([]string, len(e.RuleErrs))
	for i, ruleErr := range e.RuleErrs {
		causes[i] = ruleErr.Error()
	}
	if e.Path != "" {
		ret += e.Path + config.DelimiterFendName
	}
	return ret + strings.Join(causes, config.DelimiterRule)
}

func (e *Error) Errors() []error {
	causes := make([]error, len(e.RuleErrs))
	for i, ruleErr := range e.RuleErrs {
		causes[i] = ruleErr
	}
	return causes
}

// func (e *Error) Unwrap() error {
// 	return e.rule
// }

func AsError(err error) *Error {
	var fendErr *Error
	if errors.As(err, &fendErr) {
		return fendErr
	}
	return nil
}
