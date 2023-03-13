package rule

import (
	"context"
	"strconv"
	"strings"

	"github.com/foomo/fender/config"
)

func StringMin(expected int) StringRule {
	return func(ctx context.Context, v string) error {
		if len(v) < expected {
			return NewErrorf(NameMin, 'd', expected)
		}
		return nil
	}
}

func StringMax(expected int) StringRule {
	return func(ctx context.Context, v string) error {
		if len(v) > expected {
			return NewErrorf(NameMax, 'd', expected)
		}
		return nil
	}
}

func StringSize(expected int) StringRule {
	return func(ctx context.Context, v string) error {
		if len(v) != expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

func StringNotSize(expected int) StringRule {
	return func(ctx context.Context, v string) error {
		if len(v) == expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

// prefix

func StringHasPrefix(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if len(expected) == 0 {
			return nil
		} else if len(v) == 0 {
			return NewError(NamePrefix)
		}

		if !strings.HasPrefix(v, expected) {
			return NewError(NamePrefix)
		}
		return nil
	}
}

func StringNotHasPrefix(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if len(expected) == 0 {
			return nil
		} else if len(v) == 0 {
			return NewError(NamePrefix)
		}

		if strings.HasPrefix(v, expected) {
			return NewError(NamePrefix)
		}
		return nil
	}
}

// suffix

func StringHasSuffix(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if len(expected) == 0 {
			return nil
		} else if len(v) == 0 {
			return NewError(NameSuffix)
		}

		if !strings.HasSuffix(v, expected) {
			return NewError(NameSuffix)
		}
		return nil
	}
}

func StringNotHasSuffix(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if len(expected) == 0 {
			return nil
		} else if len(v) == 0 {
			return NewError(NameContains)
		}

		if strings.HasSuffix(v, expected) {
			return NewError(NamePrefix)
		}
		return nil
	}
}

// contains

func StringContains(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if len(expected) == 0 {
			return nil
		} else if len(v) == 0 {
			return NewError(NameContains)
		}

		if !strings.Contains(v, expected) {
			return NewError(NameContains)
		}
		return nil
	}
}

func StringNotContains(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if len(expected) == 0 {
			return nil
		} else if len(v) == 0 {
			return NewError(NameContains)
		}

		if strings.Contains(v, expected) {
			return NewError(NameContains)
		}
		return nil
	}
}

func StringContainsAny(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if len(expected) == 0 {
			return nil
		} else if len(v) == 0 {
			return NewError(NameContains)
		}

		if !strings.ContainsAny(v, expected) {
			return NewError(NameContains)
		}
		return nil
	}
}

func StringNotContainsAny(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if len(expected) == 0 {
			return nil
		} else if len(v) == 0 {
			return NewError(NameContains)
		}

		if strings.ContainsAny(v, expected) {
			return NewError(NameContains)
		}
		return nil
	}
}

// required

func StringRequired(ctx context.Context, v string) error {
	if len(v) == 0 {
		return NewError(NameRequired)
	}
	return nil
}

func StringNotRequired(ctx context.Context, v string) error {
	if len(v) == 0 {
		return ErrBreak
	}
	return nil
}

// equal

func StringEqual(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if v != expected {
			return NewError(NameEqual)
		}
		return nil
	}
}

func StringNotEqual(expected string) StringRule {
	return func(ctx context.Context, v string) error {
		if v == expected {
			return NewError(NameEqual)
		}
		return nil
	}
}

// alpha

func StringAlpha(ctx context.Context, v string) error {
	if !config.RegexAlpha.MatchString(v) {
		return NewError(NameAlpha)
	}
	return nil
}

// numeric

func StringNumeric(ctx context.Context, v string) error {
	if !config.RegexAlphaNumeric.MatchString(v) {
		return NewError(NameNumeric)
	}
	return nil
}

// alphaNumeric

func StringAlphaNumeric(ctx context.Context, v string) error {
	if !config.RegexAlphaNumeric.MatchString(v) {
		return NewError(NameAlphaNumeric)
	}
	return nil
}

// bool

func StringBool(ctx context.Context, v string) error {
	if _, err := strconv.ParseBool(v); err != nil {
		return NewError(NameBool)
	}
	return nil
}

func StringTrue(ctx context.Context, v string) error {
	if value, err := strconv.ParseBool(v); err != nil {
		return NewError(NameBool)
	} else if !value {
		return NewError(NameBool)
	}
	return nil
}

func StringFalse(ctx context.Context, v string) error {
	if value, err := strconv.ParseBool(v); err != nil {
		return NewError(NameBool)
	} else if value {
		return NewError(NameBool)
	}
	return nil
}
