package retry

import (
	"github.com/hsiafan/glow/intx"
	"time"
)

// For retry run code
type Retry struct {
	backOff     BackOff
	retryTimes  int
	shouldRetry func(err error) bool
}

// Create new Retry with options
func New(options ...func(*Retry)) *Retry {
	retry := &Retry{
		backOff:    &fixedDelayBackOff{0},
		retryTimes: intx.MaxInt,
		shouldRetry: func(err error) bool {
			return true
		},
	}
	for _, opt := range options {
		opt(retry)
	}
	return retry
}

// Wait fixed delay before each retry. Default is 0.
func FixedDelay(delay time.Duration) func(*Retry) {
	return BackOffBy(&fixedDelayBackOff{delay: delay})
}

// ExponentialBackOff delay for retry.
// initialDelay: the delay for first retry;
// maxDelay: the max delay time.
func ExponentialBackOff(initialDelay, maxDelay time.Duration) func(*Retry) {
	return BackOffBy(&binaryExponentialBackOff{initialDelay, maxDelay})
}

// Set back off strategy for retry. Default wait no delay
func BackOffBy(backOff BackOff) func(retry *Retry) {
	return func(retry *Retry) {
		retry.backOff = backOff
	}
}

// Set retry times. Default is max Int
func MaxTimes(times int) func(*Retry) {
	return func(retry *Retry) {
		retry.retryTimes = times
	}
}

// Retry if err is accepted by f. Default retry for all errors
func If(f func(err error) bool) func(*Retry) {
	return func(retry *Retry) {
		retry.shouldRetry = f
	}
}

//retry run code until no error, or max retry times exceeded
func (r *Retry) Run(f func() error) error {
	err := f()
	if err == nil {
		return nil
	}
	for i := 1; i <= r.retryTimes; i++ {
		time.Sleep(r.backOff.intervalBefore(i + 1))
		err = f()
		if err == nil || !r.shouldRetry(err) {
			return nil
		}
	}
	return err
}

// Back off strategy for retry
type BackOff interface {
	// return interval before nth retry. retry times begin with 1.
	intervalBefore(retryTimes int) time.Duration
}

type fixedDelayBackOff struct {
	delay time.Duration
}

func (b *fixedDelayBackOff) intervalBefore(retryTimes int) time.Duration {
	return b.delay
}

type binaryExponentialBackOff struct {
	initDelay time.Duration
	maxDelay  time.Duration
}

func (b *binaryExponentialBackOff) intervalBefore(retryTimes int) time.Duration {
	delay := b.initDelay << (retryTimes - 1)
	if delay > b.maxDelay {
		return b.maxDelay
	}
	return delay
}
