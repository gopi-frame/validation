package validation

import (
	"strings"

	"github.com/gopi-frame/validation/contract"
)

// Form abstract form
type Form struct {
	contract.Form
	locale   string
	messages map[string][]string
}

// SetLocale set locale
func (f *Form) SetLocale(locale string) {
	f.locale = locale
}

// Locale get locale
func (f *Form) Locale() string {
	return f.locale
}

// Empty empty
func (f *Form) Empty() bool {
	if len(f.messages) == 0 {
		return true
	}
	for _, msgs := range f.messages {
		if len(msgs) > 0 {
			return false
		}
	}
	return true
}

// Fails fails
func (f *Form) Fails() bool {
	return !f.Empty()
}

// Errors errors
func (f *Form) Errors() map[string][]string {
	return f.messages
}

// AddError add error
func (f *Form) AddError(key, message string) {
	if f.messages == nil {
		f.messages = make(map[string][]string)
	}
	msgs := f.messages[key]
	if len(msgs) == 0 {
		f.messages[key] = append(f.messages[key], strings.TrimSpace(message))
	} else {
		for _, msg := range msgs {
			if msg == message {
				return
			}
		}
		f.messages[key] = append(f.messages[key], strings.TrimSpace(message))
	}
}

// CustomRules custom rules
func (f *Form) CustomRules() []contract.Rule {
	return nil
}
