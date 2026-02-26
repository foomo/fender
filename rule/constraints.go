package rule

import (
	"cmp"

	"cmp"

	"golang.org/x/exp/constraints"
)

type (
	Integer = constraints.Integer
	Ordered = cmp.Ordered
	Float   = constraints.Float
	Number  interface {
		Integer | Float
	}
)
