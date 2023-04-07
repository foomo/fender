package rule

import (
	"github.com/foomo/fender/config"
)

const NameAlpha Name = "alpha"

var Alpha = Match(NameAlpha.String(), config.RegexAlpha)
