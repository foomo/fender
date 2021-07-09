package fender

import "github.com/foomo/fender/fend"

func Field(name string, fends ...fend.Fend) FendField {
	return FendField{
		name:  name,
		fends: fends,
	}
}
