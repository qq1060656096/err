package err

func NewErr(message string, code int) *Err {
	return &Err{
		message: message,
		code:    code,
	}
}

func NewErrWithCause(message string, code int, previous error) *Err {
	return &Err{
		previous: previous,
		message:  message,
		code:     code,
	}
}

func NewWebErr(statusCode int, message string, code int) *Err {
	return &Err{
		message:    message,
		code:       code,
		statusCode: statusCode,
		isWebErr:   true,
	}
}

func NewWebErrWithCause(statusCode int, message string, code int, previous error) *Err {
	return &Err{
		previous:   previous,
		message:    message,
		code:       code,
		statusCode: statusCode,
		isWebErr:   true,
	}
}
