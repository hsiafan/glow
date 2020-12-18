package retry

import (
	"time"
)

// Retry is for retry run code
type Retry struct {
	backoff     Backoff
	retryTimes  int
	shouldRetry func(err error) bool
}

// Option is retry options.
type Option func(*Retry)

// Create new Retry with options
// The retryTimes do not count the first execution. If retryTimes less than 0, will not do retry.
func New(retryTimes int, options ...Option) *Retry {
	retry := &Retry{
		backoff:    &FixedBackoff{0},
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

// Set back off strategy for retry. Default wait no Delay
func BackoffBy(backOff Backoff) Option {
	return func(retry *Retry) {
		retry.backoff = backOff
	}
}

// WithExponentialBackoff use new ExponentialBackoff instance. This is a wrap for BackoffBy
func WithExponentialBackoff(initialDelay, maxDelay time.Duration) Option {
	return BackoffBy(&ExponentialBackoff{
		InitDelay: initialDelay,
		MaxDelay:  maxDelay,
	})
}

// WithFixedDelay use FixedBackoff instance. This is a wrap for BackoffBy
func WithFixedDelay(delay time.Duration) Option {
	return BackoffBy(&FixedBackoff{Delay: delay})
}

// If add a retry option, only retry when err is accepted by f. Default will retry for all errors
func If(f func(err error) bool) Option {
	return func(retry *Retry) {
		retry.shouldRetry = f
	}
}

// Run retry run code until no error, or max retry times exceeded
func (r *Retry) Run(f func() error) error {
	backoff := r.backoff.Copy()
	err := f()
	if err == nil || !r.shouldRetry(err) {
		return err
	}
	for i := 1; i <= r.retryTimes; i++ {
		time.Sleep(backoff.Interval(i + 1))
		err = f()
		if err == nil || !r.shouldRetry(err) {
			return err
		}
	}
	return err
}
