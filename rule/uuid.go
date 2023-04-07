package rule

import (
	"context"

	"github.com/foomo/fender/config"
)

const NameUUID Name = "uuid"

func UUID(ctx context.Context, v string) error {
	if !config.RegexUUID.MatchString(v) {
		return NewError(NameUUID)
	}
	return nil
}
