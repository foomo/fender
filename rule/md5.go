package rule

import (
	"context"

	"github.com/foomo/fender/config"
)

const NameMD5 Name = "md5"

func MD5(ctx context.Context, v string) error {
	if !config.RegexMD5.MatchString(v) {
		return NewError(NameMD5)
	}
	return nil
}
