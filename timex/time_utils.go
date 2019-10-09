package timex

import "time"

// EpochMills convert unix epoch mills to time
func EpochMills(millis int64) time.Time {
	return time.Unix(0, millis*int64(time.Millisecond))
}

// Epoch convert unix epoch seconds to time
func Epoch(secs int64) time.Time {
	return time.Unix(secs, 0)
}
