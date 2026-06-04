package hebrew

import (
	"time"

	"github.com/dromara/carbon/v2/calendar"
)

type Locale string

const (
	EnLocale      Locale = "en"
	HeLocale      Locale = "he"
	defaultLocale        = EnLocale
	hebrewEpoch          = 347995.5
)

var (
	EnMonths = []string{"Nisan", "Iyyar", "Sivan", "Tammuz", "Av", "Elul", "Tishri", "Heshvan", "Kislev", "Teveth", "Shevat", "Adar", "Adar Bet"}
	HeMonths = []string{"ניסן", "אייר", "סיוון", "תמוז", "אב", "אלול", "תשרי", "חשוון", "כסלו", "טבת", "שבט", "אדר", "אדר ב"}
	EnWeeks  = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	HeWeeks  = []string{"ראשון", "שני", "שלישי", "רביעי", "חמישי", "שישי", "שבת"}
)

type Hebrew struct {
	year, month, day int
	Error            error
}

// NewHebrew creates a new Hebrew calendar instance with specified year, month, and day
func NewHebrew(year, month, day int) *Hebrew { _ = "STUB: not implemented"; return nil }

// FromStdTime converts standard time to Hebrew calendar date
func FromStdTime(t time.Time) *Hebrew { _ = "STUB: not implemented"; return nil }

// Special handling for January 1, 1 CE

// Authoritative implementation: directly use Julian Day Number

// ToGregorian converts Hebrew date to Gregorian date
func (h *Hebrew) ToGregorian(timezone ...string) *calendar.Gregorian {
	_ = "STUB: not implemented"
	return nil
}

// IsValid checks if the Hebrew date is valid
func (h *Hebrew) IsValid() bool { _ = "STUB: not implemented"; return false }

// Hebrew year range: 1-9999, including 3761 (corresponding to 1 CE)

// Check if month is within valid range for the year

// Check if day is within valid range for the month

// IsLeapYear checks if the Hebrew year is a leap year
func (h *Hebrew) IsLeapYear() bool { _ = "STUB: not implemented"; return false }

// Year returns the Hebrew year
func (h *Hebrew) Year() int { _ = "STUB: not implemented"; return 0 }

// Month returns the Hebrew month (1-13, where 13 is Adar Bet in leap years)
func (h *Hebrew) Month() int { _ = "STUB: not implemented"; return 0 }

// Day returns the day of the Hebrew month
func (h *Hebrew) Day() int { _ = "STUB: not implemented"; return 0 }

// String returns the Hebrew date in "YYYY-MM-DD" format
func (h *Hebrew) String() string { _ = "STUB: not implemented"; return "" }

// ToMonthString returns the Hebrew month name in the specified locale
func (h *Hebrew) ToMonthString(locale ...Locale) string { _ = "STUB: not implemented"; return "" }

// ToWeekString returns the weekday name in the specified locale
func (h *Hebrew) ToWeekString(locale ...Locale) string { _ = "STUB: not implemented"; return "" }

// gregorian2jdn converts Gregorian date to Julian Day Number
func gregorian2jdn(year, month, day int) float64 { _ = "STUB: not implemented"; return 0 }

// jdn2gregorian converts Julian Day Number to Gregorian date
func jdn2gregorian(jdn int) (year, month, day int) { _ = "STUB: not implemented"; return 0, 0, 0 }

// jdn2hebrew converts Julian Day Number to Hebrew date
func jdn2hebrew(jdn float64) (year, month, day int) {
	_ = "STUB: not implemented"
	// Estimate year
	return 0, 0, 0
}

// Precisely locate year

// Determine month

// hebrew2jdn converts Hebrew date to Julian Day Number using authoritative algorithm
func hebrew2jdn(year, month, day int) float64 { _ = "STUB: not implemented"; return 0 }

// isLeapYear checks if the Hebrew year is a leap year
func isLeapYear(year int) bool { _ = "STUB: not implemented"; return false }

// getMonthsFromEpoch calculates the number of months elapsed since the Hebrew epoch
func getMonthsFromEpoch(year int) int { _ = "STUB: not implemented"; return 0 }

// getJDNInYear calculates the Julian Day Number of Hebrew New Year (Tishri 1)
func getJDNInYear(year int) float64 { _ = "STUB: not implemented"; return 0 }

// getMonthsInYear calculates the number of months in a year
func getMonthsInYear(year int) int { _ = "STUB: not implemented"; return 0 }

// getDaysInMonth calculates the number of days in a month
func getDaysInMonth(year, month int) int {
	_ = "STUB: not implemented"
	// Fixed 29-day months
	return 0
}

// Adar in non-leap years is 29 days

// Calculate total days in the year

// Heshvan (month 8)

// Kislev (month 9)

// Other months are 30 days
