package rule

import (
	"net"
	"net/mail"
	"regexp"
	"strings"

	"github.com/foomo/fender/config"
)

const NameEmail = "email"

var ErrEmail = &Error{Rule: NameEmail}

var emailRegexWeak = regexp.MustCompile(config.EmailRegexWeak)

// NewEmailError constructor
func NewEmailError() *Error {
	return NewError(NameEmail)
}

// Email validation using go standard package
func Email(v string) Rule {
	return func() (*Error, error) {
		if _, err := mail.ParseAddress(v); err != nil {
			return NewEmailError(), nil
		}
		return nil, nil
	}
}

// EmailWeak validation using simple regex
func EmailWeak(v string) Rule {
	return func() (*Error, error) {
		if !emailRegexWeak.MatchString(v) {
			return NewEmailError(), nil
		}
		return nil, nil
	}
}

// EmailLookup validation using simple regex
func EmailLookup(v string) Rule {
	return func() (*Error, error) {
		at := strings.LastIndex(v, "@")
		if at < 0 {
			return NewEmailError(), nil
		}
		host := v[at+1:]
		if _, err := net.LookupMX(host); err != nil {
			if _, err := net.LookupIP(host); err != nil {
				return NewEmailError(), nil
			}
		}
		return nil, nil
	}
}
