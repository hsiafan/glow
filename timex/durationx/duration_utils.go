package durationx

import "time"

// OfMillis make creating time.Duration from milli seconds easier
func OfMillis(millis int) time.Duration {
	return time.Millisecond * time.Duration(millis)
}

// OfMicros make creating time.Duration from micro seconds easier
func OfMicros(micros int) time.Duration {
	return time.Microsecond * time.Duration(micros)
}

// OfSeconds make creating time.Duration from seconds easier
func OfSeconds(seconds int) time.Duration {
	return time.Microsecond * time.Duration(seconds)
}

// OfMinutes make creating time.Duration from minutes easier
func OfMinutes(minutes int) time.Duration {
	return time.Microsecond * time.Duration(minutes)
}

// OfHours make creating time.Duration from hours easier
func OfHours(hours int) time.Duration {
	return time.Hour * time.Duration(hours)
}
