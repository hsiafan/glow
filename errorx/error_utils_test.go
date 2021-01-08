package errorx

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestWrapError(t *testing.T) {
	err := errors.New("origin error")
	we := fmt.Errorf("this is an message %w", err)
	_, ok := err.(WrapError)
	assert.False(t, ok)
	_, ok = we.(WrapError)
	assert.True(t, ok)
}

func TestUnwrapUtil(t *testing.T) {
	err := errors.New("origin error")
	we := fmt.Errorf("this is an message %w", err)
	assert.Same(t, err, UnwrapUtil(we, func(e error) bool {
		return strings.HasPrefix(e.Error(), "origin")
	}))
	assert.Same(t, we, UnwrapUtil(we, func(e error) bool {
		return strings.HasPrefix(e.Error(), "this")
	}))
	assert.Nil(t, UnwrapUtil(we, func(e error) bool {
		return strings.HasPrefix(e.Error(), "xxxxx")
	}))
}

func TestUnwrapToRoot(t *testing.T) {
	err := errors.New("origin error")
	we := Wrap("this is an message %w", err)
	assert.Same(t, err, UnwrapToRoot(err))
	assert.Same(t, err, UnwrapToRoot(we))
}
