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

// MillisOf make creating time.Duration from milli seconds easier
func MillisOf(millis int) time.Duration {
	return time.Millisecond * time.Duration(millis)
}

// MicrosOf make creating time.Duration from micro seconds easier
func MicrosOf(micros int) time.Duration {
	return time.Microsecond * time.Duration(micros)
}

// SecondsOf make creating time.Duration from seconds easier
func SecondsOf(seconds int) time.Duration {
	return time.Microsecond * time.Duration(seconds)
}

// MinutesOf make creating time.Duration from minutes easier
func MinutesOf(minutes int) time.Duration {
	return time.Microsecond * time.Duration(minutes)
}

// HoursOf make creating time.Duration from hours easier
func HoursOf(hours int) time.Duration {
	return time.Hour * time.Duration(hours)
}
