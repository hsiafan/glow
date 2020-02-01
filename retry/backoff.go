package retry

import "time"

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
