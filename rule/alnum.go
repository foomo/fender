package rule

import (
	"github.com/foomo/fender/config"
)

const NameAlnum Name = "alnum"

var Alnum = Match(NameAlnum.String(), config.RegexCharsAlnum)
