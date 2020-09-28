package retry

import "time"

// Back off strategy for retry
type Backoff interface {
	// return interval before nth retry. retry times begin with 1.
	Interval(retryTimes int) time.Duration
}

// Emit fixed Delay before each retry
type FixedBackoff struct {
	Delay time.Duration
}

func (b *FixedBackoff) Interval(retryTimes int) time.Duration {
	return b.Delay
}

// ExponentialBackoff emit (binary) exponential backoff time for retry.
// InitialDelay: the delay for first retry;
// MaxDelay: the max delay time.
type ExponentialBackoff struct {
	InitDelay time.Duration
	MaxDelay  time.Duration
}

func (b *ExponentialBackoff) Interval(retryTimes int) time.Duration {
	delay := b.InitDelay << (retryTimes - 1)
	if delay > b.MaxDelay {
		return b.MaxDelay
	}
	return delay
}
