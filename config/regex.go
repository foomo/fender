package config

import (
	"regexp"
)

var (
	RegexEmailWeak    = regexp.MustCompile(`.+@.+\..+`)                                   // https://davidcel.is/posts/stop-validating-email-addresses-with-regex/
	RegexHostname     = regexp.MustCompile(`^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`) // https://tools.ietf.org/html/rfc952
	RegexCharsAlnum   = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	RegexCharsNumeric = regexp.MustCompile(`^[0-9]+$`)
	RegexCharsAlpha   = regexp.MustCompile(`^[a-zA-Z]+$`)
	RegexUUID         = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	RegexMD5          = regexp.MustCompile(`^[0-9a-f]{32}$`)
)
