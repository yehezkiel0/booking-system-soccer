package error

func ErrMapping(err error) bool {
	allErrors := make([]error, 0)
	allErrors = append(allErrors, GeneralError...)
	allErrors = append(allErrors, UserError...)

	for _, e := range allErrors {
		if e.Error() == err.Error() {
			return true
		}
	}
	return false
}
