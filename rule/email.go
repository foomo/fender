package rule

import (
	"context"
	"net"
	"net/mail"
	"strings"

	"github.com/foomo/fender/config"
)

const NameEmail Name = "email"

// Email validation using go standard package
func Email(ctx context.Context, v string) error {
	if _, err := mail.ParseAddress(v); err != nil {
		return NewError(NameEmail, "parse")
	}
	return nil
}

// EmailWeak validation using simple regex
func EmailWeak(ctx context.Context, v string) error {
	if !config.RegexEmailWeak.MatchString(v) {
		return NewError(NameEmail, "weak")
	}
	return nil
}

// EmailLookup validation using simple regex
func EmailLookup(ctx context.Context, v string) error {
	at := strings.LastIndex(v, "@")
	if at < 0 {
		return NewError(NameEmail, "lookup")
	}
	host := v[at+1:]
	if _, err := net.LookupMX(host); err != nil {
		if _, err := net.LookupIP(host); err != nil {
			return NewError(NameEmail, "lookup")
		}
	}
	return nil
}
