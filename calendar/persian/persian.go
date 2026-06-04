// Package persian is part of the carbon package.
package persian

import (
	"time"

	"github.com/dromara/carbon/v2/calendar"
)

type Locale string

const (
	EnLocale      Locale = "en"
	FaLocale      Locale = "fa"
	defaultLocale        = EnLocale
	persianEpoch         = 1948320
)

var (
	EnMonths = []string{"Farvardin", "Ordibehesht", "Khordad", "Tir", "Mordad", "Shahrivar", "Mehr", "Aban", "Azar", "Dey", "Bahman", "Esfand"}
	FaMonths = []string{"فروردین", "اردیبهشت", "خرداد", "تیر", "مرداد", "شهریور", "مهر", "آبان", "آذر", "دی", "بهمن", "اسفند"}

	EnWeeks = []string{"Yekshanbeh", "Doshanbeh", "Seshanbeh", "Chaharshanbeh", "Panjshanbeh", "Jomeh", "Shanbeh"}
	FaWeeks = []string{"نجشنبه", "دوشنبه", "سه شنبه", "چهارشنبه", "پنجشنبه", "جمعه", "شنبه"}
)

// Persian defines a Persian struct.
type Persian struct {
	year, month, day int
	Error            error
}

// NewPersian returns a new Persian instance.
func NewPersian(year, month, day int) *Persian { _ = "STUB: not implemented"; return nil }

// FromStdTime creates a Persian instance from standard time.Time.
func FromStdTime(t time.Time) (p *Persian) { _ = "STUB: not implemented"; return nil }

// ToGregorian converts Persian instance to Gregorian instance.
func (p *Persian) ToGregorian(timezone ...string) *calendar.Gregorian {
	_ = "STUB: not implemented"
	return nil
}

// Year gets the Persian year like 2020.
func (p *Persian) Year() int { _ = "STUB: not implemented"; return 0 }

// Month gets the Persian month like 8.
func (p *Persian) Month() int { _ = "STUB: not implemented"; return 0 }

// Day gets the Persian day like 5.
func (p *Persian) Day() int { _ = "STUB: not implemented"; return 0 }

// String implements the "Stringer" interface for Persian.
func (p *Persian) String() string { _ = "STUB: not implemented"; return "" }

// ToMonthString outputs a string in Persian month format like "فروردین".
func (p *Persian) ToMonthString(locale ...Locale) (month string) {
	_ = "STUB: not implemented"
	return ""
}

// ToWeekString outputs a string in week layout like "چهارشنبه".
func (p *Persian) ToWeekString(locale ...Locale) (month string) {
	_ = "STUB: not implemented"
	return ""
}

// IsValid reports whether the Persian date is valid.
func (p *Persian) IsValid() bool { _ = "STUB: not implemented"; return false }

// Check year range validation (Persian calendar starts from 622 CE)

// Check month-specific day validation

// Use IsLeapYear method

// IsLeapYear reports whether the Persian year is a leap year.
func (p *Persian) IsLeapYear() bool { _ = "STUB: not implemented"; return false }

// getPersianYear gets the Persian year from Julian Day Number.
func getPersianYear(jdn int) int { _ = "STUB: not implemented"; return 0 }

// getPersianJdn gets the Julian day number in the Persian calendar.
func getPersianJdn(year, month, day int) int { _ = "STUB: not implemented"; return 0 }

// jdn2persian converts Julian Day Number to Persian date (year, month, day).
func jdn2persian(jdn int) (year, month, day int) { _ = "STUB: not implemented"; return 0, 0, 0 }
