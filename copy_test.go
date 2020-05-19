package err

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyErrWithCause(t *testing.T) {
	testWebErr := CustomInternalServerError("testWebErr", 422, nil)
	testWebErrCopy := CopyErrWithCause(*testWebErr, nil)
	testWebErrCopy.isWebErr = false
	assert.Equal(t, testWebErr.Previous(), testWebErrCopy.Previous())
	assert.Equal(t, testWebErr.IsWebErr(), true)
	assert.Equal(t, testWebErrCopy.IsWebErr(), false)
}
