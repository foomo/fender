package rule

import (
	"context"
)

type (
	Name          string
	Rule[T any]   func(ctx context.Context, v T) error
	IntRule       func(ctx context.Context, v int) error
	Int8Rule      func(ctx context.Context, v int8) error
	Int32Rule     func(ctx context.Context, v int32) error
	Int64Rule     func(ctx context.Context, v int64) error
	UIntRule      func(ctx context.Context, v uint) error
	UInt8Rule     func(ctx context.Context, v uint8) error
	UInt32Rule    func(ctx context.Context, v uint32) error
	UInt64Rule    func(ctx context.Context, v uint64) error
	Float32Rule   func(ctx context.Context, v float32) error
	Float64Rule   func(ctx context.Context, v float64) error
	BoolRule      func(ctx context.Context, v bool) error
	StringRule    func(ctx context.Context, v string) error
	ValidatorRule func(ctx context.Context, v Validator) error
	InterfaceRule func(ctx context.Context, v interface{}) error
)

func (n Name) String() string {
	return string(n)
}
