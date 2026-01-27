package shared

type DomainError string

func (this DomainError) Error() string {
	return string(this)
}

func WrapErr(err error) error {
	if err == nil {
		return nil
	}

	return DomainError(err.Error())
}
