package fend

type Fends []Fend

func (f Fends) Append(v ...Fend) Fends {
	return append(f, v...)
}

func (f Fends) Prepend(v ...Fend) Fends {
	return append(v, f...)
}
