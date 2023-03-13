package rule

import (
	"context"
)

type (
	Fn           func(ctx context.Context) error
	Rule[T any]  func(ctx context.Context, v T) error
	Rules[T any] []Rule[T]

	BoolRule    = Rule[bool]
	StringRule  = Rule[string]
	StringRules = Rules[string]
	IntRule     = Rule[int]
	Int8Rule    = Rule[int8]
	Int32Rule   = Rule[int32]
	Int64Rule   = Rule[int64]
	UIntRule    = Rule[uint]
	UInt8Rule   = Rule[uint8]
	UInt32Rule  = Rule[uint32]
	UInt64Rule  = Rule[uint64]
	Float32Rule = Rule[float32]
	Float64Rule = Rule[float64]

	AnyRule       = Rule[any]
	ValidatorRule = Rule[Validator]
	InterfaceRule = Rule[interface{}]

	FloatRule[T Float]           Rule[T]
	IntegerRule[T Integer]       Rule[T]
	OrderedRule[T Ordered]       Rule[T]
	ComparableRule[T comparable] Rule[T]
)

func (r Rules[T]) Append(rules ...Rule[T]) Rules[T] {
	return append(r, rules...)
}

func (r Rules[T]) Prepend(rules ...Rule[T]) Rules[T] {
	return append(rules, r...)
}
