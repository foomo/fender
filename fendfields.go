package fender

type FendFields []FendField

func (f FendFields) Add(fields ...FendField) FendFields {
	return append(f, fields...)
}
