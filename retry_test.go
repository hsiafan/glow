package glow

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewFixIntervalRetry(t *testing.T) {
	retry := NewFixIntervalRetry(3, time.Duration(0))
	var i = 0
	err := retry.Run(func() error {
		i++
		if i < 2 {
			return errors.New("test error")
		}
		return nil
	})
	assert.NoError(t, err)
	assert.Equal(t, 2, i)

	i = 0
	err = retry.Run(func() error {
		i++
		return nil
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, i)
}

func TestNewRandomIntervalRetry(t *testing.T) {
	retry := NewRandomIntervalRetry(3, time.Duration(0), time.Microsecond*3)
	err := retry.Run(func() error {
		return errors.New("for test")
	})
	assert.Error(t, err)
	// retry := NewRandomIntervalRetry(3, time.Duration(-10), time.Second*-100)
	// err := retry.Run(func() error {
	// 	return errors.New("for test")
	// })
	// assert.Error(t, err)
}

func TestNewExponentialBackOff(t *testing.T) {
	retry := NewExponentialBackOff(3, time.Microsecond*1, time.Microsecond*1, time.Microsecond*10)
	err := retry.Run(func() error {
		return errors.New("for test")
	})
	assert.Error(t, err)

	// retry := NewExponentialBackOff(3, time.Duration(-1), time.Duration(10), time.Duration(1000))
	// err := retry.Run(func() error {
	// 	return errors.New("for test")
	// })
	// assert.Error(t, err)
}
