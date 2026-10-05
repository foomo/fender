package rule

import (
	"context"
	"net/url"
	"strings"
)

const NameURL Name = "url"

func URL(ctx context.Context, v string) error {
	if i := strings.Index(v, "#"); i > -1 {
		v = v[:i]
	}

	value, err := url.ParseRequestURI(v)
	if err != nil || value.Scheme == "" {
		return NewError(NameURL)
	}

	// http(s) URLs without a host are not usable destinations (e.g. "http://", "https:")
	if (value.Scheme == "http" || value.Scheme == "https") && value.Host == "" {
		return NewError(NameURL)
	}

	return nil
}
