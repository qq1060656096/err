package err

type IErr interface {
	Error() string
	setPrevious(err error)
	Previous() error
	Message() string
	Code() int
	StatusCode() int
	IsWebErr() bool
	SetLocation()
	SetRawLocation(callDepth int)
	Location() (file string, line int)
}
