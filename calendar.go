package carbon

import (
	"github.com/dromara/carbon/v2/calendar/hebrew"
	"github.com/dromara/carbon/v2/calendar/julian"
	"github.com/dromara/carbon/v2/calendar/lunar"
	"github.com/dromara/carbon/v2/calendar/persian"
)

// Lunar converts Carbon instance to Lunar instance.
func (c *Carbon) Lunar() *lunar.Lunar { _ = "STUB: not implemented"; return nil }

// CreateFromLunar creates a Carbon instance from Lunar date.
func CreateFromLunar(year, month, day int, isLeapMonth bool) *Carbon {
	_ = "STUB: not implemented"
	return nil
}

// Julian converts Carbon instance to Julian instance.
func (c *Carbon) Julian() *julian.Julian { _ = "STUB: not implemented"; return nil }

// CreateFromJulian creates a Carbon instance from Julian Day or Modified Julian Day.
func CreateFromJulian(f float64) *Carbon { _ = "STUB: not implemented"; return nil }

// Persian converts Carbon instance to Persian instance.
func (c *Carbon) Persian() *persian.Persian { _ = "STUB: not implemented"; return nil }

// CreateFromPersian creates a Carbon instance from Persian date.
func CreateFromPersian(year, month, day int) *Carbon { _ = "STUB: not implemented"; return nil }

// Hebrew converts Carbon instance to Hebrew instance.
func (c *Carbon) Hebrew() *hebrew.Hebrew { _ = "STUB: not implemented"; return nil }

// CreateFromHebrew creates a Carbon instance from Hebrew date.
func CreateFromHebrew(year, month, day int) *Carbon { _ = "STUB: not implemented"; return nil }
