package rule

import (
	"errors"
	"net"
	"net/mail"
	"regexp"
	"strings"

	"github.com/foomo/fender/config"
)

const (
	NameEmail Name = "email"
)

var ErrEmail = errors.New(NameEmail.String())

var emailRegexWeak = regexp.MustCompile(config.EmailRegexWeak)

// NewEmailError constructor
func NewEmailError() *Error {
	return NewError(ErrEmail, NameEmail.String())
}

// Email validation using go standard package
func Email(v string) Rule {
	return func() *Error {
		if _, err := mail.ParseAddress(v); err != nil {
			return NewEmailError()
		}

		return nil
	}
}

// EmailWeak validation using simple regex
func EmailWeak(v string) Rule {
	return func() *Error {
		if !emailRegexWeak.MatchString(v) {
			return NewEmailError()
		}

		return nil
	}
}

// EmailLookup validation using simple regex
func EmailLookup(v string) Rule {
	return func() *Error {
		at := strings.LastIndex(v, "@")
		if at < 0 {
			return NewEmailError()
		}
		host := v[at+1:]
		if _, err := net.LookupMX(host); err != nil {
			if _, err := net.LookupIP(host); err != nil {
				return NewEmailError()
			}
		}
		return nil
	}
}
