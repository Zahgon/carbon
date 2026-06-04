package carbon

// Parse parses a time string as a Carbon instance by default layouts.
//
// Note: it doesn't support parsing timestamp string.
func Parse(value string, timezone ...string) *Carbon { _ = "STUB: not implemented"; return nil }

// ParseByLayout parses a time string as a Carbon instance by a confirmed layout.
//
// Note: it doesn't support parsing timestamp string.
func ParseByLayout(value, layout string, timezone ...string) *Carbon {
	_ = "STUB: not implemented"
	return nil
}

// ParseByFormat parses a time string as a Carbon instance by a confirmed format.
//
// Note: If the letter used conflicts with the format sign, please use the escape character "\" to escape the letter
func ParseByFormat(value, format string, timezone ...string) *Carbon {
	_ = "STUB: not implemented"
	return nil
}

// ParseByLayouts parses a time string as a Carbon instance by multiple fuzzy layouts.
//
// Note: it doesn't support parsing timestamp string.
func ParseByLayouts(value string, layouts []string, timezone ...string) *Carbon {
	_ = "STUB: not implemented"
	return nil
}

// ParseByFormats parses a time string as a Carbon instance by multiple fuzzy formats.
//
// Note: it doesn't support parsing timestamp string.
func ParseByFormats(value string, formats []string, timezone ...string) *Carbon {
	_ = "STUB: not implemented"
	return nil
}
