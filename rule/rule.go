package rule

type (
	Rule          func() (*Error, error)
	IntRule       func(v int) Rule
	Int8Rule      func(v int8) Rule
	Int32Rule     func(v int32) Rule
	Int64Rule     func(v int64) Rule
	UIntRule      func(v uint) Rule
	UInt8Rule     func(v uint8) Rule
	UInt32Rule    func(v uint32) Rule
	UInt64Rule    func(v uint64) Rule
	Float32Rule   func(v float32) Rule
	Float64Rule   func(v float64) Rule
	BoolRule      func(v bool) Rule
	StringRule    func(v string) Rule
	ValidatorRule func(v Validator) Rule
	InterfaceRule func(v interface{}) Rule
)
