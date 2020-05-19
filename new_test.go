package err

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewErr(t *testing.T) {
	nErr := NewErr("NewErr", 1)
	err := Err{}
	err.message = "NewErr"
	err.code = 1
	assert.Equal(t, err, *nErr)
}

func TestNewErrWithCause(t *testing.T) {
	causeErr := errors.New("causeErr")
	nErr := NewErrWithCause("NewErr", 2, causeErr)
	err := Err{}
	err.message = "NewErr"
	err.code = 2
	err.previous = causeErr
	assert.Equal(t, err, *nErr)
}

func TestNewWebErr(t *testing.T) {
	nErr := NewWebErr(500, "NewWebErr", 3)
	err := Err{}
	err.message = "NewWebErr"
	err.code = 3
	err.statusCode = 500
	err.isWebErr = true
	assert.Equal(t, err, *nErr)
}

func TestNewWebErrWithCause(t *testing.T) {
	causeErr := errors.New("causeWebErr")
	nErr := NewWebErrWithCause(504, "NewWebErrWithCause", 4, causeErr)
	err := Err{}
	err.message = "NewWebErrWithCause"
	err.code = 4
	err.statusCode = 504
	err.isWebErr = true
	err.previous = causeErr
	assert.Equal(t, err, *nErr)
}