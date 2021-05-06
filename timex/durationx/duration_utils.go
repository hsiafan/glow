package durationx

import "time"

// NanoSeconds create time.Duration from nano seconds
func NanoSeconds(nanos int) time.Duration {
	return time.Nanosecond * time.Duration(nanos)
}

// MicroSeconds create time.Duration from micro seconds
func MicroSeconds(micros int) time.Duration {
	return time.Microsecond * time.Duration(micros)
}

// MilliSeconds create time.Duration from milli seconds
func MilliSeconds(millis int) time.Duration {
	return time.Millisecond * time.Duration(millis)
}

// Seconds create time.Duration from seconds
func Seconds(seconds int) time.Duration {
	return time.Second * time.Duration(seconds)
}

// Minutes create time.Duration from minutes
func Minutes(minutes int) time.Duration {
	return time.Minute * time.Duration(minutes)
}

// Hours create time.Duration from hours
func Hours(hours int) time.Duration {
	return time.Hour * time.Duration(hours)
}

// HoursDuration create time.Duration from hours
func Days(days int) time.Duration {
	return time.Hour * 24 * time.Duration(days)
}
