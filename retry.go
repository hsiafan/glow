package glow

import (
	"fmt"
	"time"

	"github.com/hsiafan/sugar/randx"
)

// Retry is a utils for retrying execution.
// This type does not mutable status, can be reused.
type Retry struct {
	MaxTimes     int                               // Retry times, including the first execution. Negative or zero value causes no retry, which has same effect as value 1.
	IntervalFunc func(retryTime int) time.Duration // A function provide interval between retries. A negative or zero duration causes next retry immediately. The retryTime param start with 1.
}

// NewFixIntervalRetry return a new Retry with fixed interval between Retry.
// If interval is negative, will panic
func NewFixIntervalRetry(times int, interval time.Duration) *Retry {
	if interval < time.Duration(0) {
		panic(fmt.Sprint("invalid interval valus:", interval))
	}
	return &Retry{MaxTimes: times, IntervalFunc: func(retryTime int) time.Duration {
		return interval
	}}
}

// NewRandomIntervalRetry return a new Retry with random interval between minInterval(inclusive) and maxInterval(inclusive).
// If minInterval or maxInterval is negative, or minInterval equals or larger than maxInterval, will panic.
func NewRandomIntervalRetry(times int, minInterval time.Duration, maxInterval time.Duration) *Retry {
	if minInterval < time.Duration(0) || maxInterval < time.Duration(0) || minInterval > maxInterval {
		panic(fmt.Sprint("invalid interval valus, minInterval:", minInterval, ", maxInterval:", maxInterval))
	}
	r := randx.New()
	return &Retry{MaxTimes: times, IntervalFunc: func(retryTime int) time.Duration {
		interval := minInterval + time.Duration(r.Int64Between(int64(minInterval), int64(maxInterval)+1))
		return interval
	}}
}

// NewExponentialBackOff return Binary Exponential Back off retry.
// The interval random chooses between [0, 2^n * initialInterval] for n-th retry.
// If initialInterval, minInterval or maxInterval is negative, or minInterval equals or larger than maxInterval, will panic.
func NewExponentialBackOff(times int, initialInterval time.Duration, minInterval time.Duration,
	maxInterval time.Duration) *Retry {
	if initialInterval < time.Duration(0) || minInterval < time.Duration(0) ||
		maxInterval < time.Duration(0) || minInterval > maxInterval {
		panic(fmt.Sprint("invalid interval valus, initialInterval:", initialInterval,
			", minInterval:", minInterval, ", maxInterval:", maxInterval))
	}
	r := randx.New()
	intervalLimit := initialInterval
	return &Retry{MaxTimes: times, IntervalFunc: func(retryTime int) time.Duration {
		if intervalLimit == time.Duration(0) {
			return minInterval
		}
		var interval time.Duration
		if intervalLimit > maxInterval {
			interval = time.Duration(r.Int64Between(int64(minInterval), int64(maxInterval)+1))
		} else {
			interval = time.Duration(r.Int64Between(int64(minInterval), int64(intervalLimit)+1))
			intervalLimit = intervalLimit * 2
		}
		return interval
	}}
}

// Run is Retry run func, until succeed or exceed max retry times.
// Return the latest error if retry final failed.
func (r *Retry) Run(f func() error) error {

	var err error
	for retryTime := 0; retryTime < r.MaxTimes; retryTime++ {
		err = f()
		if err == nil {
			break
		}
		if r.IntervalFunc != nil {
			time.Sleep(r.IntervalFunc(retryTime + 1))
		}
	}
	return err
}
