package carbon

// DiffInYears gets the difference in years.
func (c *Carbon) DiffInYears(carbon ...*Carbon) int64 { _ = "STUB: not implemented"; return 0 }

// DiffAbsInYears gets the difference in years with absolute value.
func (c *Carbon) DiffAbsInYears(carbon ...*Carbon) int64 { _ = "STUB: not implemented"; return 0 }

// DiffInMonths gets the difference in months.
func (c *Carbon) DiffInMonths(carbon ...*Carbon) int64 { _ = "STUB: not implemented"; return 0 }

// DiffAbsInMonths gets the difference in months with absolute value.
func (c *Carbon) DiffAbsInMonths(carbon ...*Carbon) int64 { _ = "STUB: not implemented"; return 0 }

// DiffInWeeks gets the difference in weeks.
func (c *Carbon) DiffInWeeks(carbon ...*Carbon) int64 { _ = "STUB: not implemented"; return 0 }

// DiffAbsInWeeks gets the difference in weeks with absolute value.
func (c *Carbon) DiffAbsInWeeks(carbon ...*Carbon) int64 { _ = "STUB: not implemented"; return 0 }

// DiffInDays gets the difference in days.
func (c *Carbon) DiffInDays(carbon ...*Carbon) int64 { _ = "STUB: not implemented"; return 0 }

// DiffAbsInDays gets the difference in days with absolute value.
func (c *Carbon) DiffAbsInDays(carbon ...*Carbon) int64 { _ = "STUB: not implemented"; return 0 }

// DiffInHours gets the difference in hours.
func (c *Carbon) DiffInHours(carbon ...*Carbon) int64 { _ = "STUB: not implemented"; return 0 }

// DiffAbsInHours gets the difference in hours with absolute value.
func (c *Carbon) DiffAbsInHours(carbon ...*Carbon) int64 { _ = "STUB: not implemented"; return 0 }

// DiffInMinutes gets the difference in minutes.
func (c *Carbon) DiffInMinutes(carbon ...*Carbon) int64 { _ = "STUB: not implemented"; return 0 }

// DiffAbsInMinutes gets the difference in minutes with absolute value.
func (c *Carbon) DiffAbsInMinutes(carbon ...*Carbon) int64 { _ = "STUB: not implemented"; return 0 }

// DiffInSeconds gets the difference in seconds.
func (c *Carbon) DiffInSeconds(carbon ...*Carbon) int64 { _ = "STUB: not implemented"; return 0 }

// DiffAbsInSeconds gets the difference in seconds with absolute value.
func (c *Carbon) DiffAbsInSeconds(carbon ...*Carbon) int64 { _ = "STUB: not implemented"; return 0 }

// DiffInString gets the difference in string, i18n is supported.
func (c *Carbon) DiffInString(carbon ...*Carbon) string { _ = "STUB: not implemented"; return "" }

// DiffAbsInString gets the difference in string with absolute value, i18n is supported.
func (c *Carbon) DiffAbsInString(carbon ...*Carbon) string { _ = "STUB: not implemented"; return "" }

// DiffInDuration gets the difference in duration.
func (c *Carbon) DiffInDuration(carbon ...*Carbon) Duration {
	_ = "STUB: not implemented"
	return *new(Duration)
}

// DiffAbsInDuration gets the difference in duration with absolute value.
func (c *Carbon) DiffAbsInDuration(carbon ...*Carbon) Duration {
	_ = "STUB: not implemented"
	return *new(Duration)
}

// DiffForHumans gets the difference in a human-readable format, i18n is supported.
func (c *Carbon) DiffForHumans(carbon ...*Carbon) string { _ = "STUB: not implemented"; return "" }

// Concurrent-safe access to language resources

// gets the difference for unit and value.
func (c *Carbon) diff(end *Carbon) (unit string, value int64) {
	_ = "STUB: not implemented"
	// Years
	return "", 0
}

// Months

// Weeks

// Days

// Hours

// Minutes

// Seconds

// gets the difference in months.
// Nil and invalid inputs return 0 to match existing tests.
func getDiffInMonths(start, end *Carbon) int64 { _ = "STUB: not implemented"; return 0 }
