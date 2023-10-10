package timex

import "time"

// IsSameDay ...
func IsSameDay(a, b time.Time) bool {
	return StartOfDay(a).Equal(StartOfDay(b))
}

// IsSameWeek ...
func IsSameWeek(a, b time.Time) bool {
	return StartOfWeek(a).Equal(StartOfWeek(b))
}

// IsSameMonth ...
func IsSameMonth(a, b time.Time) bool {
	return StartOfMonth(a).Equal(StartOfMonth(b))
}

// IsSameYear ...
func IsSameYear(a, b time.Time) bool {
	return StartOfYear(a).Equal(StartOfYear(b))
}
