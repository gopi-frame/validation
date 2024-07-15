package contract

// Form form interface
type Form interface {
	SetLocale(locale string)
	Locale() string
	AddError(key string, message string)
	Fails() bool
	Errors() map[string][]string
	CustomRules() []Rule
}
