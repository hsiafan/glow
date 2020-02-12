package retry

import (
	"time"
)

// For retry run code
type Retry struct {
	backoff     Backoff
	retryTimes  int
	shouldRetry func(err error) bool
}

// Create new Retry with options
// The retryTimes do not count the first execution. If retryTimes less than 0, will not do retry.
func New(retryTimes int, options ...func(*Retry)) *Retry {
	retry := &Retry{
		backoff:    &fixedDelayBackoff{0},
		retryTimes: retryTimes,
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
	return BackoffBy(&fixedDelayBackoff{delay: delay})
}

// ExponentialBackoff delay for retry.
// initialDelay: the delay for first retry;
// maxDelay: the max delay time.
func ExponentialBackoff(initialDelay, maxDelay time.Duration) func(*Retry) {
	return BackoffBy(&binaryExponentialBackoff{initialDelay, maxDelay})
}

// Set back off strategy for retry. Default wait no delay
func BackoffBy(backOff Backoff) func(retry *Retry) {
	return func(retry *Retry) {
		retry.backoff = backOff
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
		time.Sleep(r.backoff.intervalBefore(i + 1))
		err = f()
		if err == nil || !r.shouldRetry(err) {
			return nil
		}
	}
	return err
}
