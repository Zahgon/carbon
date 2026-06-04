package carbon

import (
	"database/sql/driver"
)

// timestamp precision constants
const (
	PrecisionSecond      = "second"
	PrecisionMillisecond = "millisecond"
	PrecisionMicrosecond = "microsecond"
	PrecisionNanosecond  = "nanosecond"
)

// TimestampType defines a TimestampType generic struct.
type TimestampType[T TimestampTyper] struct {
	*Carbon
}

// NewTimestampType returns a new TimestampType generic instance.
func NewTimestampType[T TimestampTyper](c *Carbon) *TimestampType[T] {
	_ = "STUB: not implemented"
	return nil
}

// Scan implements "driver.Scanner" interface for TimestampType generic struct.
func (t *TimestampType[T]) Scan(src any) (err error) { _ = "STUB: not implemented"; return nil }

// Value implements "driver.Valuer" interface for TimestampType generic struct.
func (t TimestampType[T]) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// MarshalJSON implements "json.Marshaler" interface for TimestampType generic struct.
func (t TimestampType[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements "json.Unmarshaler" interface for TimestampType generic struct.
func (t *TimestampType[T]) UnmarshalJSON(src []byte) error { _ = "STUB: not implemented"; return nil }

// String implements "Stringer" interface for TimestampType generic struct.
func (t *TimestampType[T]) String() string { _ = "STUB: not implemented"; return "" }

// Int64 returns the timestamp value.
func (t *TimestampType[T]) Int64() (ts int64) { _ = "STUB: not implemented"; return 0 }

// getPrecision returns precision of TimestampType generic struct.
func (t *TimestampType[T]) getPrecision() string { _ = "STUB: not implemented"; return "" }
