package carbon

var constellations = []struct {
	startMonth, startDay int
	endMonth, endDay     int
}{
	{3, 21, 4, 19},   // Aries
	{4, 20, 5, 20},   // Taurus
	{5, 21, 6, 21},   // Gemini
	{6, 22, 7, 22},   // Cancer
	{7, 23, 8, 22},   // Leo
	{8, 23, 9, 22},   // Virgo
	{9, 23, 10, 23},  // Libra
	{10, 24, 11, 22}, // Scorpio
	{11, 23, 12, 21}, // Sagittarius
	{12, 22, 1, 19},  // Capricorn
	{1, 20, 2, 18},   // Aquarius
	{2, 19, 3, 20},   // Pisces
}

// Constellation gets constellation name like "Aries", i18n is supported.
func (c *Carbon) Constellation() string { _ = "STUB: not implemented"; return "" }

// IsAries reports whether is Aries.
func (c *Carbon) IsAries() bool { _ = "STUB: not implemented"; return false }

// IsTaurus reports whether is Taurus.
func (c *Carbon) IsTaurus() bool { _ = "STUB: not implemented"; return false }

// IsGemini reports whether is Gemini.
func (c *Carbon) IsGemini() bool { _ = "STUB: not implemented"; return false }

// IsCancer reports whether is Cancer.
func (c *Carbon) IsCancer() bool { _ = "STUB: not implemented"; return false }

// IsLeo reports whether is Leo.
func (c *Carbon) IsLeo() bool { _ = "STUB: not implemented"; return false }

// IsVirgo reports whether is Virgo.
func (c *Carbon) IsVirgo() bool { _ = "STUB: not implemented"; return false }

// IsLibra reports whether is Libra.
func (c *Carbon) IsLibra() bool { _ = "STUB: not implemented"; return false }

// IsScorpio reports whether is Scorpio.
func (c *Carbon) IsScorpio() bool { _ = "STUB: not implemented"; return false }

// IsSagittarius reports whether is Sagittarius.
func (c *Carbon) IsSagittarius() bool { _ = "STUB: not implemented"; return false }

// IsCapricorn reports whether is Capricorn.
func (c *Carbon) IsCapricorn() bool { _ = "STUB: not implemented"; return false }

// IsAquarius reports whether is Aquarius.
func (c *Carbon) IsAquarius() bool { _ = "STUB: not implemented"; return false }

// IsPisces reports whether is Pisces.
func (c *Carbon) IsPisces() bool { _ = "STUB: not implemented"; return false }
