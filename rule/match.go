package rule

import (
	"context"
	"regexp"
)

const NameMatch Name = "match"

func Match(alias string, regexp *regexp.Regexp) StringRule {
	name := NameMatch
	if alias != "" {
		name = Name(alias)
	}
	return func(ctx context.Context, v string) error {
		if !regexp.MatchString(v) {
			return NewError(name, regexp.String())
		}
		return nil
	}
}

func NotMatch(alias string, regexp *regexp.Regexp) StringRule {
	name := NameMatch
	if alias != "" {
		name = Name(alias)
	}
	return func(ctx context.Context, v string) error {
		if regexp.MatchString(v) {
			return NewError(name, regexp.String())
		}
		return nil
	}
}
