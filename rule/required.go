package rule

import (
	"context"
)

const NameRequired Name = "required"

// NewRequiredError constructor
func NewRequiredError() *Error {
	return NewError(NameRequired)
}

func RequiredInt(ctx context.Context, v int) error {
	if v == 0 {
		return NewRequiredError()
	}
	return nil
}

func RequiredInt8(ctx context.Context, v int8) error {
	if v == 0 {
		return NewRequiredError()
	}
	return nil
}

func RequiredInt32(ctx context.Context, v int32) error {
	if v == 0 {
		return NewRequiredError()
	}
	return nil
}

func RequiredInt64(ctx context.Context, v int64) error {
	if v == 0 {
		return NewRequiredError()
	}
	return nil
}

func RequiredUInt(ctx context.Context, v uint) error {
	if v == 0 {
		return NewRequiredError()
	}
	return nil
}

func RequiredUInt8(ctx context.Context, v uint8) error {
	if v == 0 {
		return NewRequiredError()
	}
	return nil
}

func RequiredUInt32(ctx context.Context, v uint32) error {
	if v == 0 {
		return NewRequiredError()
	}
	return nil
}

func RequiredUInt64(ctx context.Context, v uint64) error {
	if v == 0 {
		return NewRequiredError()
	}
	return nil
}

func RequiredFloat32(ctx context.Context, v float32) error {
	if v == 0 {
		return NewRequiredError()
	}
	return nil
}

func RequiredFloat64(ctx context.Context, v float64) error {
	if v == 0 {
		return NewRequiredError()
	}
	return nil
}

func RequiredString(ctx context.Context, v string) error {
	if len(v) == 0 {
		return NewRequiredError()
	}
	return nil
}
