package timex

import "time"

const SimpleTime = "2006-01-02 15:04:05"
const SimpleTimeMills = "2006-01-02 15:04:05.000"
const SimpleDate = "2006-01-02"

// EpochMills convert unix epoch mills to time
func EpochMills(millis int64) time.Time {
	return time.Unix(0, millis*int64(time.Millisecond))
}

// Epoch convert unix epoch seconds to time
func Epoch(secs int64) time.Time {
	return time.Unix(secs, 0)
}
