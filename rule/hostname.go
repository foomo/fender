package rule

import (
	"context"

	"github.com/foomo/fender/config"
)

const NameHostname Name = "hostname"

func Hostname(ctx context.Context, v string) error {
	if !config.RegexHostname.MatchString(v) {
		return NewError(NameHostname)
	}
	return nil
}
