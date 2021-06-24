//nolint:lll
package fender

import (
	"net"
	"net/mail"
	"regexp"
	"strings"
)

const (
	RuleEmail Rule = "email"
	// EmailRegexWeak as https://davidcel.is/posts/stop-validating-email-addresses-with-regex/
	EmailRegexWeak = ".+@.+\\..+"
	// EmailRegexRFC5322 as https://stackoverflow.com/questions/201323/how-to-validate-an-email-address-using-a-regular-expression/201378#201378
	EmailRegexRFC5322 = "^(?i)(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])*$"
)

var (
	emailRegexWeak    = regexp.MustCompile(EmailRegexWeak)
	emailRegexRFC5322 = regexp.MustCompile(EmailRegexRFC5322)
)

// NewFendEmailError constructor
func NewFendEmailError() *RuleError {
	return NewFendError(RuleEmail, string(RuleEmail))
}

// Email validation using go standard package
func Email(v string, lookup bool) error {
	a, err := mail.ParseAddress(v)
	if err != nil {
		return NewFendEmailError()
	}

	if lookup {
		return emailHostLookup(a.Address)
	}

	return nil
}

// EmailWeak validation using simple regex
func EmailWeak(v string) error {
	if !emailRegexWeak.MatchString(v) {
		return NewFendEmailError()
	}

	return nil
}

// EmailRFC5322 validation using RFC5322 regex
func EmailRFC5322(v string, lookup bool) error {
	if !emailRegexRFC5322.MatchString(v) {
		return NewFendEmailError()
	}

	if lookup {
		return emailHostLookup(v)
	}

	return nil
}

func emailHostLookup(v string) error {
	at := strings.LastIndex(v, "@")
	if at < 0 {
		return NewFendEmailError()
	}
	host := v[at+1:]
	if _, err := net.LookupMX(host); err != nil {
		if _, err := net.LookupIP(host); err != nil {
			return NewFendEmailError()
		}
	}
	return nil
}
