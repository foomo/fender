package rule

import (
	"context"
	"net/url"
	"strings"
)

func URI(ctx context.Context, v string) error {
	if i := strings.Index(v, "#"); i > -1 {
		v = v[:i]
	}
	if _, err := url.ParseRequestURI(v); err != nil {
		return NewError(NameURI)
	}
	return nil
}
