package fender

import (
	"errors"
	"strings"

	"github.com/foomo/fender/config"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
)

type Error struct {
	FendErrs []*fend.Error
}

func NewError(fendErrs ...*fend.Error) *Error {
	return &Error{
		FendErrs: fendErrs,
	}
}

func NewFendError(name string, ruleErrs ...*rule.Error) *Error {
	return NewError(fend.NewError(name, ruleErrs...))
}

func NewFendRuleError(name string, ruleName rule.Name, ruleMeta ...string) *Error {
	return NewFendError(name, rule.NewError(ruleName, ruleMeta...))
}

func (e *Error) Error() string {
	causes := make([]string, len(e.FendErrs))
	for i, cause := range e.FendErrs {
		causes[i] = cause.Error()
	}
	return strings.Join(causes, config.DelimiterFend)
}

func (e *Error) First() error {
	if errs := e.Errors(); len(errs) > 0 {
		return errs[0]
	}
	return nil
}

func (e *Error) Errors() []error {
	causes := make([]error, len(e.FendErrs))
	for i, fendErr := range e.FendErrs {
		causes[i] = fendErr
	}
	return causes
}

// func (e *Error) Unwrap() error {
// 	return e.cause
// }

func AsError(err error) *Error {
	var fendErr *Error
	if errors.As(err, &fendErr) {
		return fendErr
	}
	return nil
}
