package err

func Ok(message string, previous error) *Err {
	return NewWebErrWithCause(StatusCodeOk, message, CodeOk, previous)
}

func Unauthorized(message string, previous error) *Err {
	return NewWebErrWithCause(StatusCodeUnauthorized, message, CodeUnauthorized, previous)
}

func NotFound(message string, previous error) *Err {
	return NewWebErrWithCause(StatusCodeNotFound, message, CodeNotFound, previous)
}

func CustomNotFound(message string, code int, previous error) *Err {
	return NewWebErrWithCause(StatusCodeNotFound, message, code, previous)
}

func UnprocessableEntity(message string, previous error) *Err {
	return NewWebErrWithCause(StatusCodeUnprocessableEntity, message, CodeUnprocessableEntity, previous)
}

func CustomUnprocessableEntity(message string, code int, previous error) *Err {
	return NewWebErrWithCause(StatusCodeUnprocessableEntity, message, code, previous)
}

func InternalServerError(message string, previous error) *Err {
	return NewWebErrWithCause(StatusCodeInternalServerError, message, CodeInternalServerError, previous)
}

func CustomInternalServerError(message string, code int, previous error) *Err {
	return NewWebErrWithCause(StatusCodeInternalServerError, message, code, previous)
}
