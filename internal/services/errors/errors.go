package services_errors

type ServiceError struct {
	err error
}

func (e ServiceError) Error() string {
	return e.err.Error()
}
