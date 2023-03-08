package rule

import (
	"context"
	"net"
	"net/mail"
	"regexp"
	"strings"

	"github.com/foomo/fender/config"
)

const NameEmail Name = "email"

var emailRegexWeak = regexp.MustCompile(config.EmailRegexWeak)

// NewEmailError constructor
func NewEmailError() *Error {
	return NewError(NameEmail)
}

// Email validation using go standard package
func Email(ctx context.Context, v string) error {
	if _, err := mail.ParseAddress(v); err != nil {
		return NewEmailError()
	}
	return nil
}

// EmailWeak validation using simple regex
func EmailWeak(ctx context.Context, v string) error {
	if !emailRegexWeak.MatchString(v) {
		return NewEmailError()
	}
	return nil
}

// EmailLookup validation using simple regex
func EmailLookup(ctx context.Context, v string) error {
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
