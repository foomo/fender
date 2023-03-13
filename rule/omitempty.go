package rule

import (
	"context"
)

func OmitEmptyString(ctx context.Context, v string) error {
	if len(v) == 0 {
		return ErrBreak
	}
	return nil
}
