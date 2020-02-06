package durationx

import "time"

// Millis make creating time.Duration from milli seconds easier
func Millis(millis int) time.Duration {
	return time.Millisecond * time.Duration(millis)
}

// Micros make creating time.Duration from micro seconds easier
func Micros(micros int) time.Duration {
	return time.Microsecond * time.Duration(micros)
}

// Seconds make creating time.Duration from seconds easier
func Seconds(seconds int) time.Duration {
	return time.Microsecond * time.Duration(seconds)
}

// Minutes make creating time.Duration from minutes easier
func Minutes(minutes int) time.Duration {
	return time.Microsecond * time.Duration(minutes)
}

// Hours make creating time.Duration from hours easier
func Hours(hours int) time.Duration {
	return time.Hour * time.Duration(hours)
}
