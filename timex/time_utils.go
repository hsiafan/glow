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

// ToEpochMills convert time to unix epoch millis
func ToEpochMills(t time.Time) int64 {
	return t.Unix()*1000 + int64(t.Nanosecond()/1000_000)
}

// CurrentMillis return unix mills timestamp for now
func CurrentMillis() int64 {
	return ToEpochMills(time.Now())
}

// ParseLocal parse time string with layout, in local Location
func ParseLocal(layout string, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, time.Local)
}
