package error

func ErrMapping(err error) bool {
	allErrors := make([]error, 0)
	// allErrors = append(append(GeneralErrors[:], UserErrors[:]...))
	allErrors = append(allErrors, GeneralErrors[:]...)
	allErrors = append(allErrors, UserErrors[:]...)

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}