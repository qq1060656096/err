package err

func CopyErr(err Err) *Err {
	return &err
}

func CopyErrWithCause(err Err, previous error) *Err {
	err.setPrevious(previous)
	return &err
}
