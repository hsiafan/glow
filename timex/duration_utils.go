package timex

import "time"

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
