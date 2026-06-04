// Package julian is part of the carbon package.
package julian

import (
	"time"

	"github.com/dromara/carbon/v2/calendar"
)

var (
	// julian day or modified julian day decimal precision
	decimalPrecision = 6

	// difference between Julian Day and Modified Julian Day
	diffJdFromMjd = 2400000.5
)

// Julian defines a Julian struct.
type Julian struct {
	jd, mjd float64
}

// NewJulian returns a new Lunar instance.
func NewJulian(f float64) (j *Julian) {
	_ = "STUB: not implemented"

	// get length of the integer part
	return nil
}

// modified julian day

// julian day

// FromStdTime creates a Julian instance from standard time.Time.
func FromStdTime(t time.Time) *Julian { _ = "STUB: not implemented"; return nil }

// Check if date is on or after Gregorian reform (October 15, 1582)

// ToGregorian converts Julian instance to Gregorian instance.
func (j *Julian) ToGregorian(timezone ...string) *calendar.Gregorian {
	_ = "STUB: not implemented"
	return nil
}

// JD gets julian day like 2460332.5
func (j *Julian) JD(precision ...int) float64 { _ = "STUB: not implemented"; return 0 }

// MJD gets modified julian day like 60332
func (j *Julian) MJD(precision ...int) float64 { _ = "STUB: not implemented"; return 0 }

// parseFloat64 round to n decimal places
func parseFloat64(f float64, n int) float64 { _ = "STUB: not implemented"; return 0 }
