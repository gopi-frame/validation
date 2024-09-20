package validation

import "github.com/gopi-frame/contract/validation"

type Option func(v *Validator) error

func WithTranslator(translator validation.Translator) Option {
	return func(v *Validator) error {
		v.translator = translator
		return nil
	}
}

func WithDefaultLanguage(language string) Option {
	return func(v *Validator) error {
		v.defaultLanguage = language
		return nil
	}
}

func WithErrorBuilder(builder validation.ErrorBuilder) Option {
	return func(v *Validator) error {
		v.errorBuilder = builder
		return nil
	}
}

func WithMessages(messages map[string]string) Option {
	return func(v *Validator) error {
		v.messages = messages
		return nil
	}
}
