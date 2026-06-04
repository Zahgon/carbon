package carbon

var seasons = map[int]int{
	// month: index
	1:  3, // winter
	2:  3, // winter
	3:  0, // spring
	4:  0, // spring
	5:  0, // spring
	6:  1, // summer
	7:  1, // summer
	8:  1, // summer
	9:  2, // autumn
	10: 2, // autumn
	11: 2, // autumn
	12: 3, // winter
}

// Season gets season name according to the meteorological division method like "Spring", i18n is supported.
func (c *Carbon) Season() string { _ = "STUB: not implemented"; return "" }

// StartOfSeason returns a Carbon instance for start of the season.
func (c *Carbon) StartOfSeason() *Carbon { _ = "STUB: not implemented"; return nil }

// EndOfSeason returns a Carbon instance for end of the season.
func (c *Carbon) EndOfSeason() *Carbon { _ = "STUB: not implemented"; return nil }

// IsSpring reports whether is spring.
func (c *Carbon) IsSpring() bool { _ = "STUB: not implemented"; return false }

// IsSummer reports whether is summer.
func (c *Carbon) IsSummer() bool { _ = "STUB: not implemented"; return false }

// IsAutumn reports whether is autumn.
func (c *Carbon) IsAutumn() bool { _ = "STUB: not implemented"; return false }

// IsWinter reports whether is winter.
func (c *Carbon) IsWinter() bool { _ = "STUB: not implemented"; return false }
