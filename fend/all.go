package fend

func All(values ...Fend) []error {
	var errs []error
	for _, value := range values {
		for _, v := range value() {
			if err := v(); err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errs
}
