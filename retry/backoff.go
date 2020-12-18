package retry

import "time"

// Back off strategy for retry
type Backoff interface {
	// return interval before nth retry. retry times begin with 1.
	Interval(retryTimes int) time.Duration

	// Some implement may need to create one new instance for catch retry call.
	// if so, override this method to create a copy; return self otherwise.
	Copy() Backoff
}

var _ Backoff = (*FixedBackoff)(nil)

// FixedBackoff emit fixed Delay for each retry. This backoff can be reused.
type FixedBackoff struct {
	Delay time.Duration
}

func (b *FixedBackoff) Copy() Backoff {
	return b
}

// NewFixedBackOff create and return new FixedBackoff
func NewFixedBackOff(delay time.Duration) *FixedBackoff {
	return &FixedBackoff{
		Delay: delay,
	}
}

func (b *FixedBackoff) Interval(retryTimes int) time.Duration {
	return b.Delay
}

var _ Backoff = (*ExponentialBackoff)(nil)

// ExponentialBackoff emit (binary) exponential backoff time for retry. This backoff can be reused.
// InitialDelay: the delay for first retry;
// MaxDelay: the max delay time.
type ExponentialBackoff struct {
	InitDelay time.Duration
	MaxDelay  time.Duration
}

func (b *ExponentialBackoff) Copy() Backoff {
	return b
}

// NewExponentialBackoff create and return new ExponentialBackoff
func NewExponentialBackoff(initDelay, maxDelay time.Duration) *ExponentialBackoff {
	return &ExponentialBackoff{
		InitDelay: initDelay,
		MaxDelay:  maxDelay,
	}
}

func (b *ExponentialBackoff) Interval(retryTimes int) time.Duration {
	delay := b.InitDelay << (retryTimes - 1)
	if delay > b.MaxDelay {
		return b.MaxDelay
	}
	return delay
}

var _ Backoff = (*FibonacciBackoff)(nil)

// FibonacciBackoff emit backoff time using fibonacci sequence. The sequence begin with 1, such as:
// 1, 1, 2, 3, 5, 8, 13, 21
//
// This backoff can not reuse, should create new one for each retry call.
//
// InitialDelay: the delay for first retry;
// MaxDelay: the max delay time.
type FibonacciBackoff struct {
	InitDelay time.Duration
	MaxDelay  time.Duration
	n_1       time.Duration // n-1th fibonacci sequence value
	n_2       time.Duration // n-2th fibonacci sequence value
}

func (b *FibonacciBackoff) Copy() Backoff {
	newB := *b
	return &newB
}

// NewFibonacciBackoff create and return new FibonacciBackoff
func NewFibonacciBackoff(initDelay, maxDelay time.Duration) *FibonacciBackoff {
	return &FibonacciBackoff{
		InitDelay: initDelay,
		MaxDelay:  maxDelay,
		n_1:       initDelay,
		n_2:       initDelay,
	}
}

func (b *FibonacciBackoff) Interval(retryTimes int) time.Duration {
	if retryTimes == 1 || retryTimes == 2 {
		return b.InitDelay
	}
	delay := b.n_2 + b.n_1
	if delay > b.MaxDelay {
		return b.MaxDelay
	}
	b.n_2 = b.n_1
	b.n_1 = delay
	return delay
}
