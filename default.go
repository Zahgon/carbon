package carbon

var (
	// DefaultLayout default layout
	DefaultLayout = DateTimeLayout

	// DefaultTimezone default timezone
	DefaultTimezone = UTC

	// DefaultLocale default language locale
	DefaultLocale = "en"

	// DefaultWeekStartsAt default start date of the week
	DefaultWeekStartsAt = Monday

	// DefaultWeekendDays default weekend days of the week
	DefaultWeekendDays = []Weekday{
		Saturday, Sunday,
	}
)

type Default struct {
	Layout       string
	Timezone     string
	Locale       string
	WeekStartsAt Weekday
	WeekendDays  []Weekday
}

// SetDefault sets default.
func SetDefault(d Default) { _ = "STUB: not implemented"; return }

// ResetDefault resets default.
func ResetDefault() { _ = "STUB: not implemented"; return }
