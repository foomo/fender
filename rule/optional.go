package rule

import (
	"context"
	"reflect"
)

func Optional[T any](ctx context.Context, v T) error {
	if reflect.ValueOf(v).IsZero() {
		return ErrBreak
	}
	return nil
}
