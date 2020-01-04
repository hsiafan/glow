package timex

import "time"

const SimpleTime = "2006-01-02 15:04:05"
const SimpleTimeMills = "2006-01-02 15:04:05.000"
const SimpleDate = "2006-01-02"

// OfEpochMills convert unix epoch mills to time
func OfEpochMills(millis int64) time.Time {
	return time.Unix(0, millis*int64(time.Millisecond))
}

// OfEpoch convert unix epoch seconds to time
func OfEpoch(secs int64) time.Time {
	return time.Unix(secs, 0)
}

// time to unix epoch millis
func ToEpochMills(t time.Time) int64 {
	return t.Unix()*1000 + int64(t.Nanosecond()/1000_000)
}
