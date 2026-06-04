package carbon

import (
	"time"
)

// HasError reports whether it has error.
func (c *Carbon) HasError() bool { _ = "STUB: not implemented"; return false }

// IsNil reports whether it is nil pointer.
func (c *Carbon) IsNil() bool {
	_ = "STUB: not implemented"

	// IsEmpty reports whether it is empty value.
	return false
}

func (c *Carbon) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// IsZero reports whether it is a zero time(0001-01-01 00:00:00 +0000 UTC).
func (c *Carbon) IsZero() bool { _ = "STUB: not implemented"; return false }

// IsEpoch reports whether it is a unix epoch time(1970-01-01 00:00:00 +0000 UTC).
func (c *Carbon) IsEpoch() bool { _ = "STUB: not implemented"; return false }

// IsValid reports whether it is a valid time.
func (c *Carbon) IsValid() bool { _ = "STUB: not implemented"; return false }

// IsInvalid reports whether it is an invalid time.
func (c *Carbon) IsInvalid() bool { _ = "STUB: not implemented"; return false }

// IsDST reports whether it is a daylight saving time.
func (c *Carbon) IsDST() bool { _ = "STUB: not implemented"; return false }

// IsAM reports whether it is before noon.
func (c *Carbon) IsAM() bool { _ = "STUB: not implemented"; return false }

// IsPM reports whether it is after noon.
func (c *Carbon) IsPM() bool { _ = "STUB: not implemented"; return false }

// IsLeapYear reports whether it is a leap year.
func (c *Carbon) IsLeapYear() bool { _ = "STUB: not implemented"; return false }

// IsLongYear reports whether it is a long year,
//
// refer to https://en.wikipedia.org/wiki/ISO_8601#Week_dates.
func (c *Carbon) IsLongYear() bool { _ = "STUB: not implemented"; return false }

// IsJanuary reports whether it is January.
func (c *Carbon) IsJanuary() bool { _ = "STUB: not implemented"; return false }

// IsFebruary reports whether it is February.
func (c *Carbon) IsFebruary() bool { _ = "STUB: not implemented"; return false }

// IsMarch reports whether it is March.
func (c *Carbon) IsMarch() bool { _ = "STUB: not implemented"; return false }

// IsApril reports whether it is April.
func (c *Carbon) IsApril() bool { _ = "STUB: not implemented"; return false }

// IsMay reports whether it is May.
func (c *Carbon) IsMay() bool { _ = "STUB: not implemented"; return false }

// IsJune reports whether it is June.
func (c *Carbon) IsJune() bool { _ = "STUB: not implemented"; return false }

// IsJuly reports whether it is July.
func (c *Carbon) IsJuly() bool { _ = "STUB: not implemented"; return false }

// IsAugust reports whether it is August.
func (c *Carbon) IsAugust() bool { _ = "STUB: not implemented"; return false }

// IsSeptember reports whether it is September.
func (c *Carbon) IsSeptember() bool { _ = "STUB: not implemented"; return false }

// IsOctober reports whether it is October.
func (c *Carbon) IsOctober() bool { _ = "STUB: not implemented"; return false }

// IsNovember reports whether it is November.
func (c *Carbon) IsNovember() bool { _ = "STUB: not implemented"; return false }

// IsDecember reports whether it is December.
func (c *Carbon) IsDecember() bool { _ = "STUB: not implemented"; return false }

// IsMonday reports whether it is Monday.
func (c *Carbon) IsMonday() bool { _ = "STUB: not implemented"; return false }

// IsTuesday reports whether it is Tuesday.
func (c *Carbon) IsTuesday() bool { _ = "STUB: not implemented"; return false }

// IsWednesday reports whether it is Wednesday.
func (c *Carbon) IsWednesday() bool { _ = "STUB: not implemented"; return false }

// IsThursday reports whether it is Thursday.
func (c *Carbon) IsThursday() bool { _ = "STUB: not implemented"; return false }

// IsFriday reports whether it is Friday.
func (c *Carbon) IsFriday() bool { _ = "STUB: not implemented"; return false }

// IsSaturday reports whether it is Saturday.
func (c *Carbon) IsSaturday() bool { _ = "STUB: not implemented"; return false }

// IsSunday reports whether it is Sunday.
func (c *Carbon) IsSunday() bool { _ = "STUB: not implemented"; return false }

// IsWeekday reports whether it is weekday.
func (c *Carbon) IsWeekday() bool { _ = "STUB: not implemented"; return false }

// IsWeekend reports whether it is weekend.
func (c *Carbon) IsWeekend() bool { _ = "STUB: not implemented"; return false }

// IsNow reports whether it is now time.
func (c *Carbon) IsNow() bool { _ = "STUB: not implemented"; return false }

// IsFuture reports whether it is future time.
func (c *Carbon) IsFuture() bool { _ = "STUB: not implemented"; return false }

// IsPast reports whether it is past time.
func (c *Carbon) IsPast() bool { _ = "STUB: not implemented"; return false }

// IsYesterday reports whether it is yesterday.
func (c *Carbon) IsYesterday() bool { _ = "STUB: not implemented"; return false }

// IsToday reports whether it is today.
func (c *Carbon) IsToday() bool { _ = "STUB: not implemented"; return false }

// IsTomorrow reports whether it is tomorrow.
func (c *Carbon) IsTomorrow() bool { _ = "STUB: not implemented"; return false }

// IsSameCentury reports whether it is same century.
func (c *Carbon) IsSameCentury(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// IsSameDecade reports whether it is same decade.
func (c *Carbon) IsSameDecade(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// IsSameYear reports whether it is same year.
func (c *Carbon) IsSameYear(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// IsSameQuarter reports whether it is same quarter.
func (c *Carbon) IsSameQuarter(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// IsSameMonth reports whether it is same month.
func (c *Carbon) IsSameMonth(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// IsSameDay reports whether it is same day.
func (c *Carbon) IsSameDay(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// IsSameHour reports whether it is same hour.
func (c *Carbon) IsSameHour(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// IsSameMinute reports whether it is same minute.
func (c *Carbon) IsSameMinute(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// IsSameSecond reports whether it is same second.
func (c *Carbon) IsSameSecond(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// Compare compares by an operator.
func (c *Carbon) Compare(operator string, t *Carbon) bool { _ = "STUB: not implemented"; return false }

// Gt reports whether greater than.
func (c *Carbon) Gt(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// Lt reports whether less than.
func (c *Carbon) Lt(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// Eq reports whether equal.
func (c *Carbon) Eq(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// Ne reports whether not equal.
func (c *Carbon) Ne(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// Gte reports whether greater than or equal.
func (c *Carbon) Gte(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// Lte reports whether less than or equal.
func (c *Carbon) Lte(t *Carbon) bool { _ = "STUB: not implemented"; return false }

// Between reports whether between two times, including the start and end time.
func (c *Carbon) Between(start *Carbon, end *Carbon) bool { _ = "STUB: not implemented"; return false }

// BetweenIncludedStart reports whether between two times, including the start time.
func (c *Carbon) BetweenIncludedStart(start *Carbon, end *Carbon) bool {
	_ = "STUB: not implemented"
	return false
}

// BetweenIncludedEnd reports whether between two times, including the end time.
func (c *Carbon) BetweenIncludedEnd(start *Carbon, end *Carbon) bool {
	_ = "STUB: not implemented"
	return false
}

// BetweenIncludedBoth reports whether between two times, including the start and end time.
func (c *Carbon) BetweenIncludedBoth(start *Carbon, end *Carbon) bool {
	_ = "STUB: not implemented"
	return false
}

// isMonth reports whether the current month matches the given month.
// It returns false if the Carbon instance is invalid.
func (c *Carbon) isMonth(month time.Month) bool { _ = "STUB: not implemented"; return false }

// isWeekday reports whether the current weekday matches the given weekday.
// It returns false if the Carbon instance is invalid.
func (c *Carbon) isWeekday(weekday time.Weekday) bool { _ = "STUB: not implemented"; return false }
