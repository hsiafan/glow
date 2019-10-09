package errorx

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrap(t *testing.T) {
	err := errors.New("for test")
	we := Wrap(err, "error")
	assert.Equal(t, err, we.Cause())
}

func Test_wrappedError_RootCause(t *testing.T) {
	err := errors.New("for test")
	assert.Equal(t, err, RootCause(err))
	we := Wrap(err, "error")
	assert.Equal(t, err, RootCause(we))
	we2 := Wrap(err, "error")
	assert.Equal(t, err, RootCause(we2))
}
