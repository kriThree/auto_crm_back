package storage_errors

type StorageError struct {
	err error
}

func (e StorageError) Error() string {
	return e.err.Error()
}
