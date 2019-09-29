package glow

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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
	//	retry := NewRandomIntervalRetry(3, time.Duration(0), time.Second*3)
	//	err := retry.Run(func() error {
	//		return errors.New("for test")
	//	})
	//	assert.Error(t, err)
	retry := NewRandomIntervalRetry(3, time.Duration(-10), time.Second*-100)
	err := retry.Run(func() error {
		return errors.New("for test")
	})
	errors.Unwrap()
	assert.Error(t, err)
}

func TestNewBinaryExponentialBackOff(t *testing.T) {
	//retry := NewBinaryExponentialBackOff(3, time.Second*1)
	//err := retry.Run(func() error {
	//	return errors.New("for test")
	//})
	//assert.Error(t, err)

	retry := NewBinaryExponentialBackOff(3, time.Duration(-1))
	err := retry.Run(func() error {
		return errors.New("for test")
	})
	assert.Error(t, err)
}
