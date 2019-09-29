package glow

import (
	"math/rand"
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

// NewFixIntervalRetry return a new Retry with random interval between minInterval(included) and maxInterval(excluded).
// If minInterval or maxInterval is negative, will be treat as 0.
// If minInterval equals or larger than maxInterval, will always use minInterval as interval
func NewRandomIntervalRetry(times int, minInterval time.Duration, maxInterval time.Duration) *Retry {
	if minInterval < time.Duration(0) {
		minInterval = time.Duration(0)
	}
	if maxInterval < time.Duration(0) {
		minInterval = time.Duration(0)
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Retry{MaxTimes: times, IntervalFunc: func(retryTime int) time.Duration {
		if minInterval >= maxInterval {
			return minInterval
		}
		interval := minInterval + time.Duration(random.Int63()%(int64(maxInterval)-int64(minInterval)))
		return interval
	}}
}

// NewBinaryExponentialBackOff return Binary Exponential Back off retry.
// The interval random chooses between (2^n-1 * initialInterval, 2^n * initialInterval) for n-th retry.
// If initialInterval is negative, will be treat as 0.
func NewBinaryExponentialBackOff(times int, initialInterval time.Duration) *Retry {
	if initialInterval < time.Duration(0) {
		initialInterval = time.Duration(0)
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	intervalLimit := initialInterval
	return &Retry{MaxTimes: times, IntervalFunc: func(retryTime int) time.Duration {
		if intervalLimit == time.Duration(0) {
			return intervalLimit
		}
		interval := time.Duration(random.Int63() % int64(intervalLimit))
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
