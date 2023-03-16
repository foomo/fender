package rule

import (
	"context"
	"strconv"
)

const NameBool Name = "bool"

func Bool(expectd bool) BoolRule {
	return func(ctx context.Context, v bool) error {
		if v != expectd {
			return NewError(NameBool, Meta('T', expectd))
		}
		return nil
	}
}

func StringBool(expectd bool) StringRule {
	return func(ctx context.Context, v string) error {
		if b, err := strconv.ParseBool(v); err != nil {
			return NewError(NameBool)
		} else if b != expectd {
			return NewError(NameBool, Meta('T', expectd))
		}
		return nil
	}
}
