package rule

import (
	"context"

	"github.com/foomo/fender/config"
)

func MD5(ctx context.Context, v string) error {
	if !config.RegexMD5.MatchString(v) {
		return NewError(NameMD5)
	}
	return nil
}
