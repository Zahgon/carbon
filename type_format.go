package carbon

import (
	"database/sql/driver"
)

// FormatType defines a FormatType generic struct.
type FormatType[T FormatTyper] struct {
	*Carbon
}

// NewFormatType returns a new FormatType generic instance.
func NewFormatType[T FormatTyper](c *Carbon) *FormatType[T] { _ = "STUB: not implemented"; return nil }

// Scan implements "driver.Scanner" interface for FormatType generic struct.
func (t *FormatType[T]) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// Value implements "driver.Valuer" interface for FormatType generic struct.
func (t FormatType[T]) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// MarshalJSON implements "json.Marshaler" interface for FormatType generic struct.
func (t FormatType[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements "json.Unmarshaler" interface for FormatType generic struct.
func (t *FormatType[T]) UnmarshalJSON(src []byte) error { _ = "STUB: not implemented"; return nil }

// String implements "Stringer" interface for FormatType generic struct.
func (t *FormatType[T]) String() string { _ = "STUB: not implemented"; return "" }

// getFormat returns the format of FormatType generic struct.
func (t *FormatType[T]) getFormat() string { _ = "STUB: not implemented"; return "" }
