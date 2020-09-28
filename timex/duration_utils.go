package timex

import "time"

// NanosDuration create time.Duration from nano seconds
func NanosDuration(nanos int) time.Duration {
	return time.Nanosecond * time.Duration(nanos)
}

// MicrosDuration create time.Duration from micro seconds
func MicrosDuration(micros int) time.Duration {
	return time.Microsecond * time.Duration(micros)
}

// MillisDuration create time.Duration from milli seconds
func MillisDuration(millis int) time.Duration {
	return time.Millisecond * time.Duration(millis)
}

// SecondsDuration create time.Duration from seconds
func SecondsDuration(seconds int) time.Duration {
	return time.Second * time.Duration(seconds)
}

// MinutesDuration create time.Duration from minutes
func MinutesDuration(minutes int) time.Duration {
	return time.Minute * time.Duration(minutes)
}

// HoursDuration create time.Duration from hours
func HoursDuration(hours int) time.Duration {
	return time.Hour * time.Duration(hours)
}

// HoursDuration create time.Duration from hours
func DaysDuration(days int) time.Duration {
	return time.Hour * 24 * time.Duration(days)
}
