package rule

import (
	"context"
)

type (
	Rule[T any]   func(ctx context.Context, v T) error
	DynamicRule   func(ctx context.Context) error
	Rules[T any]  []Rule[T]
	AnyRule       = Rule[any]
	BoolRule      = Rule[bool]
	StringRule    = Rule[string]
	StringRules   = Rules[string]
	IntRule       = Rule[int]
	Int8Rule      = Rule[int8]
	Int32Rule     = Rule[int32]
	Int64Rule     = Rule[int64]
	UIntRule      = Rule[uint]
	UInt8Rule     = Rule[uint8]
	UInt32Rule    = Rule[uint32]
	UInt64Rule    = Rule[uint64]
	Float32Rule   = Rule[float32]
	Float64Rule   = Rule[float64]
	InterfaceRule = Rule[interface{}]
)

func (r Rules[T]) Append(rules ...Rule[T]) Rules[T] {
	return append(r, rules...)
}

func (r Rules[T]) Prepend(rules ...Rule[T]) Rules[T] {
	return append(rules, r...)
}
