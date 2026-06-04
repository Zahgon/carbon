package carbon

import (
	"embed"
	"sync"
)

//go:embed lang
var fs embed.FS

// localeCache caches parsed locale resources to avoid repeated file loading and JSON parsing
var localeCache sync.Map

// cachedResources holds the cached resources for each language.
type cachedResources struct {
	once      sync.Once
	resources map[string]string
	err       error
}

// Language defines a Language struct.
type Language struct {
	dir       string
	locale    string
	resources map[string]string
	Error     error
	rw        *sync.RWMutex
}

// NewLanguage returns a new Language instance.
func NewLanguage() *Language { _ = "STUB: not implemented"; return nil }

// Copy returns a new copy of the current Language instance
func (lang *Language) Copy() *Language { _ = "STUB: not implemented"; return nil }

// SetLocale sets language locale.
func (lang *Language) SetLocale(locale string) *Language { _ = "STUB: not implemented"; return nil }

// Early return if locale hasn't changed and resources are already loaded

// Create a copy of the cached resources to avoid modifying the cache
// Pre-allocate with exact capacity for better memory efficiency

// SetResources sets language resources.
func (lang *Language) SetResources(resources map[string]string) *Language {
	_ = "STUB: not implemented"
	return nil
}

// returns a translated string.
func (lang *Language) translate(unit string, value int64) string {
	_ = "STUB: not implemented"
	return ""
}

// If resources is empty, set default locale and retry
