package rule

import (
	"context"
)

type (
	Rule[T any]   func(ctx context.Context, v T) error
	AnyRule       = Rule[any]
	BoolRule      = Rule[bool]
	StringRule    = Rule[string]
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
	DynamicRule   func(ctx context.Context) error
)
