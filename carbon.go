// @Package carbon
// @Description a simple, semantic and developer-friendly time package for golang
// @Source github.com/dromara/carbon
// @Document carbon.go-pkg.com
// @Developer gouguoyin
// @Email 245629560@qq.com

// Package carbon is a simple, semantic and developer-friendly time package for golang.
package carbon

import (
	"time"
)

type StdTime = time.Time
type Weekday = time.Weekday
type Location = time.Location
type Duration = time.Duration

// Carbon defines a Carbon struct.
type Carbon struct {
	time          StdTime
	weekStartsAt  Weekday
	weekendDays   []Weekday
	loc           *Location
	lang          *Language
	currentLayout string
	isEmpty       bool
	Error         error
}

// NewCarbon returns a new Carbon instance.
func NewCarbon(stdTime ...StdTime) *Carbon { _ = "STUB: not implemented"; return nil }

// Copy returns a copy of the Carbon instance.
func (c *Carbon) Copy() *Carbon { _ = "STUB: not implemented"; return nil }

// Create a deep copy of weekendDays slice to avoid shared reference

// Sleep sleeps for the specified duration like time.Sleep.
func Sleep(d time.Duration) { _ = "STUB: not implemented"; return }
