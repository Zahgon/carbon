// Package calendar is part of the carbon package.
package calendar

import (
	"time"
)

// Gregorian defines a Gregorian struct.
type Gregorian struct {
	Time  time.Time
	Error error
}

// String implements "Stringer" interface.
func (g *Gregorian) String() string { _ = "STUB: not implemented"; return "" }

func (g *Gregorian) IsLeapYear() bool { _ = "STUB: not implemented"; return false }
