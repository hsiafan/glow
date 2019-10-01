package glow

import (
	"github.com/hsiafan/sugar/randx"
	"time"
)

// Retry is a utils for retrying execution.
// This type does not mutable status, can be reused.
type Retry struct {
	MaxTimes     int                               // Retry times, including the first execution. Negative or zero value causes no retry, which has same effect as value 1.
	IntervalFunc func(retryTime int) time.Duration // A function provide interval between retries. A negative or zero duration causes next retry immediately. The retryTime param start with 1.
}

// NewFixIntervalRetry return a new Retry with fixed interval between Retry.
// If interval is negative, will be treat as 0.
func NewFixIntervalRetry(times int, interval time.Duration) *Retry {
	if interval < time.Duration(0) {
		interval = time.Duration(0)
	}
	return &Retry{MaxTimes: times, IntervalFunc: func(retryTime int) time.Duration {
		return interval
	}}
}

// NewFixIntervalRetry return a new Retry with random interval between minInterval(inclusive) and maxInterval(inclusive).
// If minInterval or maxInterval is negative, will be treat as 0.
// If minInterval equals or larger than maxInterval, will always use minInterval as interval
func NewRandomIntervalRetry(times int, minInterval time.Duration, maxInterval time.Duration) *Retry {
	if minInterval < time.Duration(0) {
		minInterval = time.Duration(0)
	}
	if maxInterval < time.Duration(0) {
		minInterval = time.Duration(0)
	}
	r := randx.New()
	return &Retry{MaxTimes: times, IntervalFunc: func(retryTime int) time.Duration {
		if minInterval >= maxInterval {
			return minInterval
		}
		interval := minInterval + time.Duration(r.Int64Between(int64(minInterval), int64(maxInterval)+1))
		return interval
	}}
}

// NewBinaryExponentialBackOff return Binary Exponential Back off retry.
// The interval random chooses between [0, 2^n * initialInterval] for n-th retry.
// If initialInterval is negative, will be treat as 0.
func NewBinaryExponentialBackOff(times int, initialInterval time.Duration) *Retry {
	if initialInterval < time.Duration(0) {
		initialInterval = time.Duration(0)
	}
	r := randx.New()
	intervalLimit := initialInterval
	return &Retry{MaxTimes: times, IntervalFunc: func(retryTime int) time.Duration {
		if intervalLimit == time.Duration(0) {
			return intervalLimit
		}
		interval := time.Duration(r.Int64Within(int64(intervalLimit) + 1))
		intervalLimit = intervalLimit * 2
		return interval
	}}
}

// Retry run func, until succeed or exceed max retry times.
// Return the latest error if retry final failed.
func (r *Retry) Run(f func() error) error {

	var err error = nil
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
