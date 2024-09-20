package validation

import (
	"context"
	"github.com/gopi-frame/contract/validation"
	error2 "github.com/gopi-frame/validation/error"
	"github.com/gopi-frame/validation/translator"
)

var _instance *Validator

func init() {
	instance, err := NewValidator()
	if err != nil {
		panic(err)
	}
	_instance = instance
}

// SetDefaultValidator replaces the default validator instance with the given one.
func SetDefaultValidator(v *Validator) {
	_instance = v
}

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
	bag := error2.NewErrorBag(v2.translator)
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

func Validate(ctx context.Context, builders ...validation.ValidatorBuilder) *error2.ErrorBag {
	return _instance.Validate(ctx, builders...).(*error2.ErrorBag)
}

func ValidateIt(ctx context.Context, validatable validation.Validatable) validation.ErrorBag {
	return _instance.Validate(ctx, NewBuilder(validatable))
}

func Value[T any](ctx context.Context, value T, rules ...validation.Rule[T]) validation.ErrorBag {
	return Attribute(ctx, "value", value, rules...)
}

func Attribute[T any](ctx context.Context, attribute string, value T, rules ...validation.Rule[T]) validation.ErrorBag {
	return _instance.Validate(ctx, Group[T](attribute, value, rules...))
}
