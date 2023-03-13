package rule

import (
	"context"
	"regexp"
)

func RegexMatch(name string, regexp ...*regexp.Regexp) StringRule {
	return func(ctx context.Context, v string) error {
		if v == "" {
			return NewError(NameRegex, name)
		}
		for _, r := range regexp {
			if !r.MatchString(v) {
				return NewError(NameRegex, name)
			}
		}
		return nil
	}
}

func RegexNotMatch(name string, regexp ...*regexp.Regexp) StringRule {
	return func(ctx context.Context, v string) error {
		if v == "" {
			return NewError(NameRegex, name)
		}
		for _, r := range regexp {
			if r.MatchString(v) {
				return NewError(NameRegex, name)
			}
		}
		return nil
	}
}
