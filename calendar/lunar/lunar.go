// Package lunar is part of the carbon package.
package lunar

import (
	"time"

	"github.com/dromara/carbon/v2/calendar"
)

var (
	numbers = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	months  = []string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "腊"}
	weeks   = []string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}
	animals = []string{"猴", "鸡", "狗", "猪", "鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊"}

	festivals = map[string]string{
		// "month-day": "name"
		"1-1":   "春节",
		"1-15":  "元宵节",
		"2-2":   "龙抬头",
		"3-3":   "上巳节",
		"5-5":   "端午节",
		"7-7":   "七夕节",
		"7-15":  "中元节",
		"8-15":  "中秋节",
		"9-9":   "重阳节",
		"10-1":  "寒衣节",
		"10-15": "下元节",
		"12-8":  "腊八节",
	}

	years = []int{
		0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2, // 1900-1909
		0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977, // 1910-1919
		0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970, // 1920-1929
		0x06566, 0x0d4a0, 0x0ea50, 0x16a95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950, // 1930-1939
		0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557, // 1940-1949
		0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5d0, 0x14573, 0x052d0, 0x0a9a8, 0x0e950, 0x06aa0, // 1950-1959
		0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0, // 1960-1969
		0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b5a0, 0x195a6, // 1970-1979
		0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570, // 1980-1989
		0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x05ac0, 0x0ab60, 0x096d5, 0x092e0, // 1990-1999
		0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5, // 2000-2009
		0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930, // 2010-2019
		0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530, // 2020-2029
		0x05aa0, 0x076a3, 0x096d0, 0x04bd7, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45, // 2030-2039
		0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0, // 2040-2049
		0x14b63, 0x09370, 0x049f8, 0x04970, 0x064b0, 0x168a6, 0x0ea50, 0x06b20, 0x1a6c4, 0x0aae0, // 2050-2059
		0x0a2e0, 0x0d2e3, 0x0c960, 0x0d557, 0x0d4a0, 0x0da50, 0x05d55, 0x056a0, 0x0a6d0, 0x055d4, // 2060-2069
		0x052d0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x0b6a6, 0x0ad50, 0x055a0, 0x0aba4, 0x0a5b0, 0x052b0, // 2070-2079
		0x0b273, 0x06930, 0x07337, 0x06aa0, 0x0ad50, 0x14b55, 0x04b60, 0x0a570, 0x054e4, 0x0d160, // 2080-2089
		0x0e968, 0x0d520, 0x0daa0, 0x16aa6, 0x056d0, 0x04ae0, 0x0a9d4, 0x0a2d0, 0x0d150, 0x0f252, // 2090-2099
		0x0d520, // 2100
	}

	maxYear = 2100
	minYear = 1900
)

// Lunar defines a Lunar struct.
type Lunar struct {
	year, month, day int
	isLeapMonth      bool
	Error            error
}

// NewLunar returns a new Lunar instance.
func NewLunar(year, month, day int, isLeapMonth bool) *Lunar { _ = "STUB: not implemented"; return nil }

// FromStdTime creates a Lunar instance from standard time.Time.
func FromStdTime(t time.Time) *Lunar { _ = "STUB: not implemented"; return nil }

// ToGregorian converts Lunar instance to Gregorian instance.
func (l *Lunar) ToGregorian(timezone ...string) *calendar.Gregorian {
	_ = "STUB: not implemented"
	return nil
}

// add the time difference of the month before the leap month

// https://github.com/dromara/carbon/issues/219

// Animal gets lunar animal name like "猴".
func (l *Lunar) Animal() string { _ = "STUB: not implemented"; return "" }

// Festival gets lunar festival name like "春节".
func (l *Lunar) Festival() string { _ = "STUB: not implemented"; return "" }

// Year gets lunar year like 2020.
func (l *Lunar) Year() int { _ = "STUB: not implemented"; return 0 }

// Month gets lunar month like 8.
func (l *Lunar) Month() int { _ = "STUB: not implemented"; return 0 }

// Day gets lunar day like 5.
func (l *Lunar) Day() int { _ = "STUB: not implemented"; return 0 }

// LeapMonth gets lunar leap month like 2.
func (l *Lunar) LeapMonth() int { _ = "STUB: not implemented"; return 0 }

// String implements "Stringer" interface for Lunar.
func (l *Lunar) String() string { _ = "STUB: not implemented"; return "" }

// ToYearString outputs a string in lunar year format like "二零二零".
func (l *Lunar) ToYearString() (year string) { _ = "STUB: not implemented"; return "" }

// ToMonthString outputs a string in lunar month format like "正月".
func (l *Lunar) ToMonthString() (month string) { _ = "STUB: not implemented"; return "" }

// ToWeekString outputs a string in week layout like "周一".
func (l *Lunar) ToWeekString() (month string) { _ = "STUB: not implemented"; return "" }

// ToDayString outputs a string in lunar day format like "廿一".
func (l *Lunar) ToDayString() (day string) { _ = "STUB: not implemented"; return "" }

// ToDateString outputs a string in lunar date format like "二零二零年腊月初五".
// 获取农历日期字符串，如 "二零二零年腊月初五"
func (l *Lunar) ToDateString() string { _ = "STUB: not implemented"; return "" }

// IsValid reports whether is a valid lunar date.
func (l *Lunar) IsValid() bool { _ = "STUB: not implemented"; return false }

// IsLeapYear reports whether is a lunar leap year.
func (l *Lunar) IsLeapYear() bool { _ = "STUB: not implemented"; return false }

// IsLeapMonth reports whether is a lunar leap month.
func (l *Lunar) IsLeapMonth() bool { _ = "STUB: not implemented"; return false }

// IsRatYear reports whether is lunar year of Rat.
func (l *Lunar) IsRatYear() bool { _ = "STUB: not implemented"; return false }

// IsOxYear reports whether is lunar year of Ox.
func (l *Lunar) IsOxYear() bool { _ = "STUB: not implemented"; return false }

// IsTigerYear reports whether is lunar year of Tiger.
func (l *Lunar) IsTigerYear() bool { _ = "STUB: not implemented"; return false }

// IsRabbitYear reports whether is lunar year of Rabbit.
func (l *Lunar) IsRabbitYear() bool { _ = "STUB: not implemented"; return false }

// IsDragonYear reports whether is lunar year of Dragon.
func (l *Lunar) IsDragonYear() bool { _ = "STUB: not implemented"; return false }

// IsSnakeYear reports whether is lunar year of Snake.
func (l *Lunar) IsSnakeYear() bool { _ = "STUB: not implemented"; return false }

// IsHorseYear reports whether is lunar year of Horse.
func (l *Lunar) IsHorseYear() bool { _ = "STUB: not implemented"; return false }

// IsGoatYear reports whether is lunar year of Goat.
func (l *Lunar) IsGoatYear() bool { _ = "STUB: not implemented"; return false }

// IsMonkeyYear reports whether is lunar year of Monkey.
func (l *Lunar) IsMonkeyYear() bool { _ = "STUB: not implemented"; return false }

// IsRoosterYear reports whether is lunar year of Rooster.
func (l *Lunar) IsRoosterYear() bool { _ = "STUB: not implemented"; return false }

// IsDogYear reports whether is lunar year of Dog.
func (l *Lunar) IsDogYear() bool { _ = "STUB: not implemented"; return false }

// IsPigYear reports whether is lunar year of Pig.
func (l *Lunar) IsPigYear() bool { _ = "STUB: not implemented"; return false }

// getOffsetInYear calculates the total number of days from the beginning of the year to the specified month.
// It handles leap months by adding the leap month days when encountered.
// Returns the offset in days.
func getOffsetInYear(year, month int) int { _ = "STUB: not implemented"; return 0 }

// getOffsetInMonth calculates the total number of days from the minimum year (1900) to the specified year.
// This represents the cumulative days across all years up to but not including the target year.
// Returns the offset in days.
func getOffsetInMonth(year int) int { _ = "STUB: not implemented"; return 0 }

// getDaysInYear calculates the total number of days in a lunar year.
// It uses the lunar calendar data array to determine which months have 30 days vs 29 days.
// The base is 348 days (12 months × 29 days), then adds days for months with 30 days.
// Finally adds the leap month days if the year has a leap month.
// Returns the total number of days in the year.
func getDaysInYear(year int) int { _ = "STUB: not implemented"; return 0 }

// getDaysInMonth calculates the number of days in a specific lunar month.
// It uses the lunar calendar data array to determine if the month has 30 or 29 days.
// The bit pattern in the data array indicates which months are long (30 days).
// Returns 30 for long months, 29 for short months.
func getDaysInMonth(year, month int) int { _ = "STUB: not implemented"; return 0 }

// getDaysInLeapMonth calculates the number of days in the leap month of a lunar year.
// If the year has no leap month, returns 0.
// If the year has a leap month, determines if it's a long (30 days) or short (29 days) month.
// Returns the number of days in the leap month, or 0 if no leap month exists.
func getDaysInLeapMonth(year int) int { _ = "STUB: not implemented"; return 0 }

// getLeapMonth determines which month is the leap month in a lunar year.
// Returns 0 if the year has no leap month, or the month number (1-12) if a leap month exists.
// The leap month information is stored in the lower 4 bits of the lunar calendar data.
// Returns 0 for years outside the supported range (1900-2100).
func getLeapMonth(year int) int { _ = "STUB: not implemented"; return 0 }
