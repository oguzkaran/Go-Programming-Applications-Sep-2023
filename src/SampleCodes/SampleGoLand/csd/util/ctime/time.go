package ctime

import "time"

func EndOfMonthYear(month time.Month, year int, location *time.Location) time.Time {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, location)
}

func EndOfMonth(dt *time.Time) time.Time {
	return time.Date(dt.Year(), dt.Month()+1, 0, 0, 0, 0, 0, dt.Location())
}

func EndOfMonthDayValue(dt *time.Time) int {
	t := EndOfMonth(dt)

	return t.Day()
}
