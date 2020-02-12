package retry

import "time"

// Back off strategy for retry
type Backoff interface {
	// return interval before nth retry. retry times begin with 1.
	intervalBefore(retryTimes int) time.Duration
}

type fixedDelayBackoff struct {
	delay time.Duration
}

func (b *fixedDelayBackoff) intervalBefore(retryTimes int) time.Duration {
	return b.delay
}

type binaryExponentialBackoff struct {
	initDelay time.Duration
	maxDelay  time.Duration
}

func (b *binaryExponentialBackoff) intervalBefore(retryTimes int) time.Duration {
	delay := b.initDelay << (retryTimes - 1)
	if delay > b.maxDelay {
		return b.maxDelay
	}
	return delay
}
