package carbon

import (
	"database/sql/driver"
)

// LayoutType defines a LayoutType generic struct
type LayoutType[T LayoutTyper] struct {
	*Carbon
}

// NewLayoutType returns a new LayoutType generic instance.
func NewLayoutType[T LayoutTyper](c *Carbon) *LayoutType[T] { _ = "STUB: not implemented"; return nil }

// Scan implements "driver.Scanner" interface for LayoutType generic struct.
func (t *LayoutType[T]) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Value implements "driver.Valuer" interface for LayoutType generic struct.
func (t LayoutType[T]) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// MarshalJSON implements "json.Marshaler" interface for LayoutType generic struct.
func (t LayoutType[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements "json.Unmarshaler" interface for LayoutType generic struct.
func (t *LayoutType[T]) UnmarshalJSON(src []byte) error { _ = "STUB: not implemented"; return nil }

// String implements "Stringer" interface for LayoutType generic struct.
func (t *LayoutType[T]) String() string { _ = "STUB: not implemented"; return "" }

// getLayout returns the layout of LayoutType generic struct.
func (t *LayoutType[T]) getLayout() string { _ = "STUB: not implemented"; return "" }
