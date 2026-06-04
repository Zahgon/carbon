package carbon

import (
	"database/sql/driver"
)

// Scan implements "driver.Scanner" interface for Carbon struct.
func (c *Carbon) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Value implements "driver.Valuer" interface for Carbon struct.
func (c Carbon) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// MarshalJSON implements "json.Marshaler" interface for Carbon struct.
func (c Carbon) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements "json.Unmarshaler" interface for Carbon struct.
func (c *Carbon) UnmarshalJSON(src []byte) error { _ = "STUB: not implemented"; return nil }

// String implements "Stringer" interface for Carbon struct.
func (c *Carbon) String() string { _ = "STUB: not implemented"; return "" }
