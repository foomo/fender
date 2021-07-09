package fend

func First(values ...Fend) error {
	for _, value := range values {
		for _, v := range value() {
			if err := v(); err != nil {
				return err
			}
		}
	}
	return nil
}
