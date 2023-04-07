package rule

type (
	Rules[T any]   []Rule[T]
	AnyRules       = Rules[any]
	BoolRules      = Rules[bool]
	StringRules    = Rules[string]
	IntRules       = Rules[int]
	Int8Rules      = Rules[int8]
	Int32Rules     = Rules[int32]
	Int64Rules     = Rules[int64]
	UIntRules      = Rules[uint]
	UInt8Rules     = Rules[uint8]
	UInt32Rules    = Rules[uint32]
	UInt64Rules    = Rules[uint64]
	Float32Rules   = Rules[float32]
	Float64Rules   = Rules[float64]
	InterfaceRules = Rules[interface{}]
)

func (r Rules[T]) Append(rules ...Rule[T]) Rules[T] {
	return append(r, rules...)
}

func (r Rules[T]) Prepend(rules ...Rule[T]) Rules[T] {
	return append(rules, r...)
}
