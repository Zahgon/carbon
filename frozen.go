package carbon

import (
	"sync"
)

// FrozenNow defines a FrozenNow struct.
type FrozenNow struct {
	isFrozen int32
	testNow  *Carbon
	rw       sync.RWMutex
}

var frozenNow = &FrozenNow{}

// SetTestNow sets a test Carbon instance for now.
func SetTestNow(c *Carbon) { _ = "STUB: not implemented"; return }

// ClearTestNow clears the test Carbon instance for now.
func ClearTestNow() { _ = "STUB: not implemented"; return }

// IsTestNow reports whether is testing time.
func IsTestNow() bool { _ = "STUB: not implemented"; return false }
