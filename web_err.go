package err

type WebErr struct {
	statusCode int
	Err
}

func NewWebErr(statusCode int, message string, code int) *WebErr {
	return &WebErr{
		statusCode: statusCode,
		Err: *NewErr(message, code),
	}
}

func NewWebErrWithCause(statusCode int, message string, code int, previous error) *WebErr {
	return &WebErr{
		statusCode: statusCode,
		Err: *NewErrWithCause(message, code, previous),
	}
}

func (e *WebErr) StatusCode() int {
	return e.statusCode
}

func Ok(message string, previous error) *WebErr {
	return NewWebErrWithCause(StatusCodeOk, message, CodeOk, previous)
}

func Unauthorized(message string, previous error) *WebErr {
	return NewWebErrWithCause(StatusCodeUnauthorized, message, CodeUnauthorized, previous)
}

func NotFound(message string, previous error) *WebErr {
	return NewWebErrWithCause(StatusCodeNotFound, message, CodeNotFound, previous)
}

func CustomNotFound(message string, code int, previous error) *WebErr {
	return NewWebErrWithCause(StatusCodeNotFound, message, code, previous)
}

func UnprocessableEntity(message string, previous error) *WebErr {
	return NewWebErrWithCause(StatusCodeUnprocessableEntity, message, CodeUnprocessableEntity, previous)
}

func CustomUnprocessableEntity(message string, code int, previous error) *WebErr {
	return NewWebErrWithCause(StatusCodeUnprocessableEntity, message, code, previous)
}

func InternalServerError(message string, previous error) *WebErr {
	return NewWebErrWithCause(StatusCodeInternalServerError, message, CodeInternalServerError, previous)
}

func CustomInternalServerError(message string, code int, previous error) *WebErr {
	return NewWebErrWithCause(StatusCodeInternalServerError, message, code, previous)
}
