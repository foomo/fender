package rule

import (
	"context"
	"net/url"
	"strings"
)

func URL(ctx context.Context, v string) error {
	if i := strings.Index(v, "#"); i > -1 {
		v = v[:i]
	}
	if value, err := url.ParseRequestURI(v); err != nil || value.Scheme == "" {
		return NewError(NameURL)
	}
	return nil
}
