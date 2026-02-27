package rule

import (
	"cmp"

	"golang.org/x/exp/constraints" //nolint:exptostd // not all supported yet
)

type (
	Integer = constraints.Integer
	Ordered = cmp.Ordered
	Float   = constraints.Float
	Number  interface {
		Integer | Float
	}
)
