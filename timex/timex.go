package timex

import "time"

// StartOfDay ...
func StartOfDay(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}

// StartOfNextDay ...
func StartOfNextDay(now time.Time) time.Time {
	return StartOfDay(now).AddDate(0, 0, 1)
}

// StartOfLastDay ...
func StartOfLastDay(now time.Time) time.Time {
	return StartOfDay(now).AddDate(0, 0, -1)
}

// StartOfWeek ...
func StartOfWeek(now time.Time) time.Time {
	weekIndex := (int)(now.Weekday())
	if weekIndex == 0 {
		weekIndex = 7
	}
	return time.Date(now.Year(), now.Month(), now.Day()-weekIndex+1, 0, 0, 0, 0, time.Local)
}

// StartOfNextWeek ...
func StartOfNextWeek(now time.Time) time.Time {
	return StartOfWeek(now).AddDate(0, 0, 7)
}

// StartOfLastWeek ...
func StartOfLastWeek(now time.Time) time.Time {
	return StartOfWeek(now).AddDate(0, 0, -7)
}

// StartOfMonth ...
func StartOfMonth(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
}

// StartOfNextMonth ...
func StartOfNextMonth(now time.Time) time.Time {
	return StartOfMonth(now).AddDate(0, 1, 0)
}

// StartOfLastMonth ...
func StartOfLastMonth(now time.Time) time.Time {
	return StartOfMonth(now).AddDate(0, -1, 0)
}

// StartOfYear ...
func StartOfYear(now time.Time) time.Time {
	return time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
}

// StartOfNextYear ...
func StartOfNextYear(now time.Time) time.Time {
	return StartOfYear(now).AddDate(1, 0, 0)
}

// StartOfLastYear ...
func StartOfLastYear(now time.Time) time.Time {
	return StartOfYear(now).AddDate(-1, 0, 0)
}

// StartOfToday ...
func StartOfToday() time.Time {
	return StartOfDay(time.Now())
}

// StartOfTomorrow ...
func StartOfTomorrow() time.Time {
	return StartOfNextDay(time.Now())
}

// StartOfYesterday ...
func StartOfYesterday() time.Time {
	return StartOfLastDay(time.Now())
}

// StartOfWeekNow ...
func StartOfWeekNow() time.Time {
	return StartOfWeek(time.Now())
}

// StartOfNextWeekNow ...
func StartOfNextWeekNow() time.Time {
	return StartOfNextWeek(time.Now())
}

// StartOfLastWeekNow ...
func StartOfLastWeekNow() time.Time {
	return StartOfLastWeek(time.Now())
}

// StartOfMonthNow ...
func StartOfMonthNow() time.Time {
	return StartOfMonth(time.Now())
}

// StartOfNextMonthNow ...
func StartOfNextMonthNow() time.Time {
	return StartOfNextMonth(time.Now())
}

// StartOfLastMonthNow ...
func StartOfLastMonthNow() time.Time {
	return StartOfLastMonth(time.Now())
}

// StartOfYearNow ...
func StartOfYearNow() time.Time {
	return StartOfYear(time.Now())
}

// StartOfNextYearNow ...
func StartOfNextYearNow() time.Time {
	return StartOfNextYear(time.Now())
}

// StartOfLastYearNow ...
func StartOfLastYearNow() time.Time {
	return StartOfLastYear(time.Now())
}

// EndOfDay ...
func EndOfDay(now time.Time) time.Time {
	return StartOfNextDay(now).Add(-time.Nanosecond)
}

// EndOfNextDay ...
func EndOfNextDay(now time.Time) time.Time {
	return EndOfDay(now).AddDate(0, 0, 1)
}

// EndOfLastDay ...
func EndOfLastDay(now time.Time) time.Time {
	return StartOfDay(now).Add(-time.Nanosecond)
}

// EndOfWeek ...
func EndOfWeek(now time.Time) time.Time {
	return StartOfNextWeek(now).Add(-time.Nanosecond)
}

// EndOfNextWeek ...
func EndOfNextWeek(now time.Time) time.Time {
	return EndOfWeek(now).AddDate(0, 0, 7)
}

// EndOfLastWeek ...
func EndOfLastWeek(now time.Time) time.Time {
	return StartOfWeek(now).Add(-time.Nanosecond)
}

// EndOfMonth ...
func EndOfMonth(now time.Time) time.Time {
	return StartOfNextMonth(now).Add(-time.Nanosecond)
}

// EndOfNextMonth ...
func EndOfNextMonth(now time.Time) time.Time {
	return StartOfMonth(now).AddDate(0, 2, 0).Add(-time.Nanosecond)
}

// EndOfLastMonth ...
func EndOfLastMonth(now time.Time) time.Time {
	return StartOfMonth(now).Add(-time.Nanosecond)
}

// EndOfYear ...
func EndOfYear(now time.Time) time.Time {
	return StartOfNextYear(now).Add(-time.Nanosecond)
}

// EndOfNextYear ...
func EndOfNextYear(now time.Time) time.Time {
	return StartOfYear(now).AddDate(2, 0, 0).Add(-time.Nanosecond)
}

// EndOfLastYear ...
func EndOfLastYear(now time.Time) time.Time {
	return StartOfYear(now).Add(-time.Nanosecond)
}

// EndOfToday ...
func EndOfToday() time.Time {
	return EndOfDay(time.Now())
}

// EndOfTomorrow ...
func EndOfTomorrow() time.Time {
	return EndOfNextDay(time.Now())
}

// EndOfYesterday ...
func EndOfYesterday() time.Time {
	return EndOfLastDay(time.Now())
}

// EndOfWeekNow ...
func EndOfWeekNow() time.Time {
	return EndOfWeek(time.Now())
}

// EndOfNextWeekNow ...
func EndOfNextWeekNow() time.Time {
	return EndOfNextWeek(time.Now())
}

// EndOfLastWeekNow ...
func EndOfLastWeekNow() time.Time {
	return EndOfLastWeek(time.Now())
}

// EndOfMonthNow ...
func EndOfMonthNow() time.Time {
	return EndOfMonth(time.Now())
}

// EndOfNextMonthNow ...
func EndOfNextMonthNow() time.Time {
	return EndOfNextMonth(time.Now())
}

// EndOfLastMonthNow ...
func EndOfLastMonthNow() time.Time {
	return EndOfLastMonth(time.Now())
}

// EndOfYearNow ...
func EndOfYearNow() time.Time {
	return EndOfYear(time.Now())
}

// EndOfNextYearNow ...
func EndOfNextYearNow() time.Time {
	return EndOfNextYear(time.Now())
}

// EndOfLastYearNow ...
func EndOfLastYearNow() time.Time {
	return EndOfLastYear(time.Now())
}
