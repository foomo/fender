package rule_test

import (
	"context"
	"testing"

	"github.com/foomo/fender/rule"
)

func TestURLRequiresHostForHTTP(t *testing.T) {
	t.Parallel()

	cases := map[string]bool{
		"http://example.com":  true,
		"https://example.com": true,
		"http://example.com/": true,
		"ftp://example.com":   true,
		"http://":             false,
		"https:":              false,
		"http:///":            false,
		"":                    false,
	}

	for input, wantOK := range cases {
		err := rule.URL(context.Background(), input)
		gotOK := err == nil
		if gotOK != wantOK {
			t.Errorf("URL(%q) ok=%v want %v (err=%v)", input, gotOK, wantOK, err)
		}
	}
}
