package err

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestErr_SetLocation(t *testing.T) {
	err := NewErr("test", 500100100)
	err.SetLocation()
	file, line := err.Location()
	fmt.Print(file)
	assert.Equal(t, 12, line)
	_, file0, _, _ := runtime.Caller(0)
	assert.Equal(t, file0, file)
}
