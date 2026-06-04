package carbon

// StdTime gets standard time.Time.
func (c *Carbon) StdTime() StdTime { _ = "STUB: not implemented"; return *new(StdTime) }

// DaysInYear gets total days in year like 365.
func (c *Carbon) DaysInYear() int { _ = "STUB: not implemented"; return 0 }

// DaysInMonth gets total days in month like 30.
func (c *Carbon) DaysInMonth() int { _ = "STUB: not implemented"; return 0 }

// MonthOfYear gets month of year like 12.
func (c *Carbon) MonthOfYear() int { _ = "STUB: not implemented"; return 0 }

// DayOfYear gets day of year like 365.
func (c *Carbon) DayOfYear() int { _ = "STUB: not implemented"; return 0 }

// DayOfMonth gets day of month like 30.
func (c *Carbon) DayOfMonth() int { _ = "STUB: not implemented"; return 0 }

// DayOfWeek gets day of week like 6, start from 1.
func (c *Carbon) DayOfWeek() int { _ = "STUB: not implemented"; return 0 }

// WeekOfYear gets week of year like 1.
//
// refer to https://en.wikipedia.org/wiki/ISO_8601#Week_dates.
func (c *Carbon) WeekOfYear() int { _ = "STUB: not implemented"; return 0 }

// WeekOfMonth gets week of month like 1.
func (c *Carbon) WeekOfMonth() int { _ = "STUB: not implemented"; return 0 }

// DateTime gets current year, month, day, hour, minute, and second like 2020, 8, 5, 13, 14, 15.
func (c *Carbon) DateTime() (year, month, day, hour, minute, second int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, 0, 0
}

// DateTimeMilli gets current year, month, day, hour, minute, second and millisecond like 2020, 8, 5, 13, 14, 15, 999.
func (c *Carbon) DateTimeMilli() (year, month, day, hour, minute, second, millisecond int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, 0, 0, 0
}

// DateTimeMicro gets current year, month, day, hour, minute, second and microsecond like 2020, 8, 5, 13, 14, 15, 999999.
func (c *Carbon) DateTimeMicro() (year, month, day, hour, minute, second, microsecond int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, 0, 0, 0
}

// DateTimeNano gets current year, month, day, hour, minute, second and nanosecond like 2020, 8, 5, 13, 14, 15, 999999999.
func (c *Carbon) DateTimeNano() (year, month, day, hour, minute, second, nanosecond int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, 0, 0, 0
}

// Date gets current year, month, and day like 2020, 8, 5.
func (c *Carbon) Date() (year, month, day int) { _ = "STUB: not implemented"; return 0, 0, 0 }

// DateMilli gets current year, month, day and millisecond like 2020, 8, 5, 999.
func (c *Carbon) DateMilli() (year, month, day, millisecond int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// DateMicro gets current year, month, day and microsecond like 2020, 8, 5, 999999.
func (c *Carbon) DateMicro() (year, month, day, microsecond int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// DateNano gets current year, month, day and nanosecond like 2020, 8, 5, 999999999.
func (c *Carbon) DateNano() (year, month, day, nanosecond int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// Time gets current hour, minute, and second like 13, 14, 15.
func (c *Carbon) Time() (hour, minute, second int) { _ = "STUB: not implemented"; return 0, 0, 0 }

// TimeMilli gets current hour, minute, second and millisecond like 13, 14, 15, 999.
func (c *Carbon) TimeMilli() (hour, minute, second, millisecond int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// TimeMicro gets current hour, minute, second and microsecond like 13, 14, 15, 999999.
func (c *Carbon) TimeMicro() (hour, minute, second, microsecond int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// TimeNano gets current hour, minute, second and nanosecond like 13, 14, 15, 999999999.
func (c *Carbon) TimeNano() (hour, minute, second, nanosecond int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// Century gets current century like 21.
func (c *Carbon) Century() int { _ = "STUB: not implemented"; return 0 }

// Decade gets current decade like 20.
func (c *Carbon) Decade() int { _ = "STUB: not implemented"; return 0 }

// Year gets current year like 2020.
func (c *Carbon) Year() int { _ = "STUB: not implemented"; return 0 }

// Quarter gets current quarter like 3.
func (c *Carbon) Quarter() (quarter int) { _ = "STUB: not implemented"; return 0 }

// Month gets current month like 8.
func (c *Carbon) Month() int { _ = "STUB: not implemented"; return 0 }

// Week gets current week like 6, start from 0.
func (c *Carbon) Week() int { _ = "STUB: not implemented"; return 0 }

// Day gets current day like 5.
func (c *Carbon) Day() int { _ = "STUB: not implemented"; return 0 }

// Hour gets current hour like 13.
func (c *Carbon) Hour() int { _ = "STUB: not implemented"; return 0 }

// Minute gets current minute like 14.
func (c *Carbon) Minute() int { _ = "STUB: not implemented"; return 0 }

// Second gets current second like 9.
func (c *Carbon) Second() int { _ = "STUB: not implemented"; return 0 }

// Millisecond gets current millisecond like 999.
func (c *Carbon) Millisecond() int { _ = "STUB: not implemented"; return 0 }

// Microsecond gets current microsecond like 999999.
func (c *Carbon) Microsecond() int { _ = "STUB: not implemented"; return 0 }

// Nanosecond gets current nanosecond like 999999999.
func (c *Carbon) Nanosecond() int { _ = "STUB: not implemented"; return 0 }

// Timestamp gets timestamp with second precision like 1596604455.
func (c *Carbon) Timestamp() int64 { _ = "STUB: not implemented"; return 0 }

// TimestampMilli gets timestamp with millisecond precision like 1596604455000.
func (c *Carbon) TimestampMilli() int64 { _ = "STUB: not implemented"; return 0 }

// TimestampMicro gets timestamp with microsecond precision like 1596604455000000.
func (c *Carbon) TimestampMicro() int64 { _ = "STUB: not implemented"; return 0 }

// TimestampNano gets timestamp with nanosecond precision like 1596604455000000000.
func (c *Carbon) TimestampNano() int64 { _ = "STUB: not implemented"; return 0 }

// Timezone gets timezone location like "Asia/Shanghai".
func (c *Carbon) Timezone() string { _ = "STUB: not implemented"; return "" }

// ZoneName gets timezone name like "CST".
func (c *Carbon) ZoneName() string { _ = "STUB: not implemented"; return "" }

// ZoneOffset gets timezone offset seconds from the UTC timezone like 28800.
func (c *Carbon) ZoneOffset() int { _ = "STUB: not implemented"; return 0 }

// Locale gets locale name like "zh-CN".
func (c *Carbon) Locale() string { _ = "STUB: not implemented"; return "" }

// WeekStartsAt returns start day of the week.
func (c *Carbon) WeekStartsAt() Weekday { _ = "STUB: not implemented"; return *new(Weekday) }

// WeekEndsAt returns end day of the week.
func (c *Carbon) WeekEndsAt() Weekday { _ = "STUB: not implemented"; return *new(Weekday) }

// CurrentLayout returns the layout used for parsing the time string.
func (c *Carbon) CurrentLayout() string { _ = "STUB: not implemented"; return "" }

// Age gets age like 18.
func (c *Carbon) Age() int { _ = "STUB: not implemented"; return 0 }
