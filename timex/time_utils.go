package timex

import (
	"time"
)

const SimpleTime = "2006-01-02 15:04:05"
const SimpleTimeMills = "2006-01-02 15:04:05.000"
const SimpleDate = "2006-01-02"

// EpochMillsTime convert unix epoch mills to time
func EpochMillsTime(millis int64) time.Time {
	return time.Unix(0, millis*int64(time.Millisecond))
}

// EpochTime convert unix epoch seconds to time
func EpochTime(secs int64) time.Time {
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

// Date create a time, at the beginning of day.
func Date(year int, month time.Month, day int, loc *time.Location) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, loc)
}

// LocalDate create a time, at the beginning of day, at local time zone.
func LocalDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

// TruncateMinute return the beginning time of the minute.
func TruncateMinute(t time.Time) time.Time {
	return t.Truncate(time.Minute)
}

// TruncateMinute return the beginning time of the hour.
func TruncateHour(t time.Time) time.Time {
	return t.Truncate(time.Hour)
}

// TruncateDay return the beginning time of the day.
func TruncateDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// TruncateMonth return the beginning time of the month.
func TruncateMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

// TruncateYear return the beginning time of the month.
func TruncateYear(t time.Time) time.Time {
	year, _, _ := t.Date()
	return time.Date(year, time.January, 1, 0, 0, 0, 0, t.Location())
}

// TruncateWeek return the beginning time of the week(Sunday).
func TruncateWeek(t time.Time) time.Time {
	weekday := t.Weekday()
	t = t.AddDate(0, 0, int(weekday-time.Sunday))
	return TruncateDay(t)
}
