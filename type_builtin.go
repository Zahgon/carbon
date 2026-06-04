package carbon

type (
	timestampType      int64
	timestampMicroType int64
	timestampMilliType int64
	timestampNanoType  int64

	datetimeType      string
	datetimeMicroType string
	datetimeMilliType string
	datetimeNanoType  string

	dateType      string
	dateMilliType string
	dateMicroType string
	dateNanoType  string

	timeType      string
	timeMilliType string
	timeMicroType string
	timeNanoType  string
)

type (
	Timestamp      = TimestampType[timestampType]
	TimestampMilli = TimestampType[timestampMilliType]
	TimestampMicro = TimestampType[timestampMicroType]
	TimestampNano  = TimestampType[timestampNanoType]

	DateTime      = LayoutType[datetimeType]
	DateTimeMicro = LayoutType[datetimeMicroType]
	DateTimeMilli = LayoutType[datetimeMilliType]
	DateTimeNano  = LayoutType[datetimeNanoType]

	Date      = LayoutType[dateType]
	DateMilli = LayoutType[dateMilliType]
	DateMicro = LayoutType[dateMicroType]
	DateNano  = LayoutType[dateNanoType]

	Time      = LayoutType[timeType]
	TimeMilli = LayoutType[timeMilliType]
	TimeMicro = LayoutType[timeMicroType]
	TimeNano  = LayoutType[timeNanoType]
)

func NewTimestamp(c *Carbon) *Timestamp { _ = "STUB: not implemented"; return nil }

func NewTimestampMilli(c *Carbon) *TimestampMilli { _ = "STUB: not implemented"; return nil }

func NewTimestampMicro(c *Carbon) *TimestampMicro { _ = "STUB: not implemented"; return nil }

func NewTimestampNano(c *Carbon) *TimestampNano { _ = "STUB: not implemented"; return nil }

func NewDateTime(c *Carbon) *DateTime { _ = "STUB: not implemented"; return nil }

func NewDateTimeMilli(c *Carbon) *DateTimeMilli { _ = "STUB: not implemented"; return nil }

func NewDateTimeMicro(c *Carbon) *DateTimeMicro { _ = "STUB: not implemented"; return nil }

func NewDateTimeNano(c *Carbon) *DateTimeNano { _ = "STUB: not implemented"; return nil }

func NewDate(c *Carbon) *Date { _ = "STUB: not implemented"; return nil }

func NewDateMilli(c *Carbon) *DateMilli { _ = "STUB: not implemented"; return nil }

func NewDateMicro(c *Carbon) *DateMicro { _ = "STUB: not implemented"; return nil }

func NewDateNano(c *Carbon) *DateNano { _ = "STUB: not implemented"; return nil }

func NewTime(c *Carbon) *Time { _ = "STUB: not implemented"; return nil }

func NewTimeMilli(c *Carbon) *TimeMilli { _ = "STUB: not implemented"; return nil }

func NewTimeMicro(c *Carbon) *TimeMicro { _ = "STUB: not implemented"; return nil }

func NewTimeNano(c *Carbon) *TimeNano { _ = "STUB: not implemented"; return nil }

func (t timestampType) Precision() string { _ = "STUB: not implemented"; return "" }

func (t timestampMilliType) Precision() string { _ = "STUB: not implemented"; return "" }

func (t timestampMicroType) Precision() string { _ = "STUB: not implemented"; return "" }

func (t timestampNanoType) Precision() string { _ = "STUB: not implemented"; return "" }

func (t datetimeType) Layout() string { _ = "STUB: not implemented"; return "" }

func (t datetimeMilliType) Layout() string { _ = "STUB: not implemented"; return "" }

func (t datetimeMicroType) Layout() string { _ = "STUB: not implemented"; return "" }

func (t datetimeNanoType) Layout() string { _ = "STUB: not implemented"; return "" }

func (t dateType) Layout() string { _ = "STUB: not implemented"; return "" }

func (t dateMilliType) Layout() string { _ = "STUB: not implemented"; return "" }

func (t dateMicroType) Layout() string { _ = "STUB: not implemented"; return "" }

func (t dateNanoType) Layout() string { _ = "STUB: not implemented"; return "" }

func (t timeType) Layout() string { _ = "STUB: not implemented"; return "" }

func (t timeMilliType) Layout() string { _ = "STUB: not implemented"; return "" }

func (t timeMicroType) Layout() string { _ = "STUB: not implemented"; return "" }

func (t timeNanoType) Layout() string { _ = "STUB: not implemented"; return "" }
