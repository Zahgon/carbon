package carbon

const (
	minDuration Duration = -1 << 63
	maxDuration Duration = 1<<63 - 1
)

// ZeroValue returns the zero value of Carbon instance.
func ZeroValue() *Carbon {
	_ = "STUB: not implemented"

	// EpochValue returns the unix epoch value of Carbon instance.
	return nil
}

func EpochValue() *Carbon { _ = "STUB: not implemented"; return nil }

// MaxValue returns the maximum value of Carbon instance.
func MaxValue() *Carbon { _ = "STUB: not implemented"; return nil }

// MinValue returns the minimum value of Carbon instance.
func MinValue() *Carbon { _ = "STUB: not implemented"; return nil }

// MaxDuration returns the maximum value of duration instance.
func MaxDuration() Duration {
	_ = "STUB: not implemented"

	// MinDuration returns the minimum value of duration instance.
	return *new(Duration)
}

func MinDuration() Duration {
	_ = "STUB: not implemented"

	// Max returns the maximum Carbon instance from some given Carbon instances.
	return *new(Duration)
}

func Max(c1 *Carbon, c2 ...*Carbon) (c *Carbon) {
	_ = "STUB: not implemented"
	// If first carbon is invalid, return it immediately
	return nil
}

// If no additional arguments, return the first one

// Check all additional arguments

// If any carbon is invalid, return it immediately

// Update maximum if current carbon is greater or equal

// Min returns the minimum Carbon instance from some given Carbon instances.
func Min(c1 *Carbon, c2 ...*Carbon) (c *Carbon) {
	_ = "STUB: not implemented"
	// If first carbon is invalid, return it immediately
	return nil
}

// If no additional arguments, return the first one

// Check all additional arguments

// If any carbon is invalid, return it immediately

// Update minimum if current carbon is less or equal

// Closest returns the closest Carbon instance from some given Carbon instances.
func (c *Carbon) Closest(c1 *Carbon, c2 ...*Carbon) *Carbon {
	_ = "STUB: not implemented"
	// Validate the base carbon instance
	return nil
}

// Validate the first comparison instance

// If no additional arguments, return the first one

// Find the closest among all instances

// Check all additional arguments

// Validate each argument

// Calculate difference and update if closer

// Farthest returns the farthest Carbon instance from some given Carbon instances.
func (c *Carbon) Farthest(c1 *Carbon, c2 ...*Carbon) *Carbon {
	_ = "STUB: not implemented"
	// Validate the base carbon instance
	return nil
}

// Validate the first comparison instance

// If no additional arguments, return the first one

// Find the farthest among all instances

// Check all additional arguments

// Validate each argument

// Calculate difference and update if farther
