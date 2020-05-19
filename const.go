package err

import "net/http"

const (
	StatusCodeOk = http.StatusOK
	CodeOk       = http.StatusOK

	StatusCodeUnauthorized = http.StatusUnauthorized
	CodeUnauthorized       = http.StatusUnauthorized

	StatusCodeNotFound = http.StatusNotFound
	CodeNotFound       = http.StatusNotFound

	StatusCodeUnprocessableEntity = http.StatusUnprocessableEntity
	CodeUnprocessableEntity       = http.StatusUnprocessableEntity

	StatusCodeInternalServerError = http.StatusInternalServerError
	CodeInternalServerError       = http.StatusInternalServerError
)
