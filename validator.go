package validation

import (
	"context"

	"github.com/gopi-frame/contract/validation"
	error2 "github.com/gopi-frame/validation/errpack"
	"github.com/gopi-frame/validation/translator"
)

type contextKey string

var languageKey contextKey = "language"

// BindLanguage binds language to context.
// It is useful when you want to change the language temporarily.
func BindLanguage(ctx context.Context, language string) context.Context {
	return context.WithValue(ctx, languageKey, language)
}

// LanguageFromContext returns language from context.
func LanguageFromContext(ctx context.Context) string {
	l, _ := ctx.Value(languageKey).(string)
	return l
}

type validateContext struct {
	validators map[string][]validation.Validatable
}

func (v *validateContext) AddValidate(key string, validator validation.Validatable) {
	if v.validators == nil {
		v.validators = make(map[string][]validation.Validatable)
	}
	v.validators[key] = append(v.validators[key], validator)
}

type Validator struct {
	translator      validation.Translator
	defaultLanguage string
	errorBuilder    validation.ErrorBuilder
	messages        map[string]string
}

func NewValidator(options ...Option) (*Validator, error) {
	v := new(Validator)
	v.translator = translator.New()
	for _, option := range options {
		if err := option(v); err != nil {
			return nil, err
		}
	}
	return v, nil
}

func (v *Validator) clone() *Validator {
	return &Validator{
		translator:      v.translator,
		defaultLanguage: v.defaultLanguage,
		errorBuilder:    v.errorBuilder,
		messages:        v.messages,
	}
}

func (v *Validator) Validate(ctx context.Context, builders ...validation.ValidatorBuilder) validation.ErrorBag {
	validatorCtx := new(validateContext)
	for _, builder := range builders {
		builder.Build(validatorCtx)
	}
	v2 := v.clone()
	if language := LanguageFromContext(ctx); language != "" {
		v2.translator = v2.translator.Locale(language)
	} else if v2.defaultLanguage != "" {
		v2.translator = v2.translator.Locale(v.defaultLanguage)
	}
	bag := error2.NewBag()
	for key, validators := range validatorCtx.validators {
		for _, v := range validators {
			if err := v.Validate(ctx, v2); err != nil {
				if message, ok := v2.messages[err.Code()]; ok {
					err = err.SetMessage(message)
				}
				bag.AddError(key, err)
			}
		}
	}
	return bag
}

func (v *Validator) BuildError(code string, message string, params ...validation.Param) validation.Error {
	if v.errorBuilder != nil {
		return v.errorBuilder.BuildError(code, message, params...)
	}
	return error2.NewError(code, message, params...).SetTranslator(v.translator)
}
