package rule

import (
	"context"
	"reflect"
)

func Required[T any](ctx context.Context, v T) error {
	if reflect.ValueOf(v).IsZero() {
		return NewError(NameRequired)
	}
	return nil
}
