package fender

import "github.com/foomo/fender/fend"

type FendField struct {
	name  string
	fends []fend.Fend
}

func Field(name string, fends ...fend.Fend) FendField {
	return FendField{
		name:  name,
		fends: fends,
	}
}
