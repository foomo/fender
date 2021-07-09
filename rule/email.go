//nolint:lll
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

var	ErrEmail = errors.New(NameEmail.String())

var (
	emailRegexWeak    = regexp.MustCompile(config.EmailRegexWeak)
	emailRegexRFC5322 = regexp.MustCompile(config.EmailRegexRFC5322)
)

// NewEmailError constructor
func NewEmailError() *Error {
	return NewError(ErrEmail)
}

// Email validation using go standard package
func Email(lookup bool) StringRule {
	return func(v string) Rule {
		return func() *Error {
			a, err := mail.ParseAddress(v)
			if err != nil {
				return NewEmailError()
			}

			if lookup {
				return emailLookup(a.Address)
			}

			return nil
		}
	}
}

// EmailWeak validation using simple regex
func EmailWeak() StringRule {
	return func(v string) Rule {
		return func() *Error {
			if !emailRegexWeak.MatchString(v) {
				return NewEmailError()
			}

			return nil
		}
	}
}

// EmailRFC5322 validation using RFC5322 regex
func EmailRFC5322(lookup bool) StringRule {
	return func(v string) Rule {
		return func() *Error {
			if !emailRegexRFC5322.MatchString(v) {
				return NewEmailError()
			}

			if lookup {
				return emailLookup(v)
			}

			return nil
		}
	}
}

func emailLookup(v string) *Error {
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
