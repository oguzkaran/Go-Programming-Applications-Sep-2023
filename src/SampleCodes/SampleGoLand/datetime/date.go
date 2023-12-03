package datetime

import (
	"fmt"
)

type CSDDate struct {
	day, month, year int
	dayOfWeek        int
}

var dayOfWeeksEN = [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
var daysOfMonths = [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func NewDate(day, month, year int) (*CSDDate, error) {
	if isValidDate(day, month, year) {
		return &CSDDate{day, month, year, dayOfWeek(day, month, year)}, nil
	}

	return nil, fmt.Errorf("invalid date values -> d:%d, m:%d, y:%d", day, month, year)
}

func (d *CSDDate) IsLeap() bool {
	return isLeapYear(d.year)
}

func (d *CSDDate) GetDayOfYear() int {
	return getDayOfYear(d.day, d.month, d.year)
}

func (d *CSDDate) GetDayOfWeek() int {
	return d.dayOfWeek
}

func (d *CSDDate) GetDayOfWeekEN() string {
	return dayOfWeeksEN[d.dayOfWeek]
}

func (d *CSDDate) Set(day, month, year int) error {
	if !isValidDate(day, month, year) {
		return fmt.Errorf("invalid date values -> d:%d, m:%d, y:%d", day, month, year)
	}

	d.day = day
	d.month = month
	d.year = year
	d.dayOfWeek = dayOfWeek(d.day, d.month, d.year)

	return nil
}

func (d *CSDDate) SetDay(day int) error {
	if isValidDate(day, d.month, d.year) {
		d.day = day
		d.dayOfWeek = dayOfWeek(d.day, d.month, d.year)
		return fmt.Errorf("invalid day value -> %d", day)
	}

	return nil
}

func (d *CSDDate) GetDay(day int) int {
	return d.day
}

func (d *CSDDate) SetMonth(month int) error {
	if isValidDate(d.day, month, d.year) {
		d.month = month
		d.dayOfWeek = dayOfWeek(d.day, d.month, d.year)
		return fmt.Errorf("invalid month value -> %d", month)
	}

	return nil
}
func (d *CSDDate) GetMonth(day int) int {
	return d.month
}

func (d *CSDDate) SetYear(year int) error {
	if isValidDate(d.day, d.month, year) {
		d.year = year
		d.dayOfWeek = dayOfWeek(d.day, d.month, d.year)
		return fmt.Errorf("invalid year value -> %d", year)
	}

	return nil
}

func (d *CSDDate) GetYear(day int) int {
	return d.year
}

func dayOfWeek(day, month, year int) int {
	return totalDays(getDayOfYear(day, month, year), year) % 7
}

func totalDays(dof, year int) int {
	result := dof

	for y := 1900; y < year; y++ {
		result += 365
		if isLeapYear(y) {
			result++
		}
	}

	return result
}

func getDayOfYear(day, month, year int) int {
	result := day

	for m := month - 1; m >= 1; m-- {
		result += daysOfMonths[m]
	}

	if month > 2 && isLeapYear(year) {
		result++
	}
	return result
}

func isValidDate(day, month, year int) bool {
	return (1 <= day && day <= 31) && (1 <= month && month <= 12) && (1900 <= year) && (day <= getDays(month, year))
}

func getDays(month, year int) int {

	if month == 2 && isLeapYear(year) {
		return 29
	}

	return daysOfMonths[month]
}

func isLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
