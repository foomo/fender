package rule

import (
	"github.com/foomo/fender/config"
)

const (
	NameNumeric Name = "numeric"
)

var Numeric = Match(NameNumeric.String(), config.RegexNumeric)
