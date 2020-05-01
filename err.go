package err

import "runtime"

type Err struct {
	previous error
	message string
	code int
	file string
	line int
}

func NewErr(message string, code int) *Err {
	return &Err{
		message: message,
		code: code,
	}
}

func NewErrWithCause(message string, code int, previous error) *Err {
	return &Err{
		previous: previous,
		message: message,
		code: code,
	}
}

func CopyErr(err Err) *Err {
	newErr := err
	return &newErr
}

func CopyErrWithCause(err Err, previous error) *Err {
	newErr := err
	newErr.previous = previous
	return &newErr
}

func (e *Err) Error() string {
	return e.message
}

func (e *Err) Previous() error {
	return e.previous
}

func (e *Err) Message() string {
	return e.message
}

func (e *Err) Code() int {
	return e.code
}

func (e *Err) SetLocation() {
	e.SetRawLocation(1)
}

func (e *Err) SetRawLocation(callDepth int) {
	_, file, line, _ := runtime.Caller(callDepth + 1)
	e.file = file
	e.line = line
}

func (e *Err) Location() (file string, line int) {
	return e.file, e.line
}

