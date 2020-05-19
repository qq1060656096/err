package err

import "runtime"

type Err struct {
	previous   error
	message    string
	code       int
	statusCode int
	isWebErr   bool
	file       string
	line       int
}

func (e *Err) Error() string {
	return e.message
}

func (e *Err) setPrevious(err error) {
	e.previous = err
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

func (e *Err) StatusCode() int {
	return e.statusCode
}

func (e *Err) IsWebErr() bool {
	return e.isWebErr
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
