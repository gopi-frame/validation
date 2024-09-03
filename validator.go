package validation

import (
	"context"
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/translator"
	"github.com/gopi-frame/validation/validator"
)

var _instance *Validator

func init() {
	instance, err := NewValidator()
	if err != nil {
		panic(err)
	}
	_instance = instance
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

type Option func(v validation.Validator) error

type Validator struct {
	translator      validation.Translator
	defaultLanguage string
	errorBuilder    validation.ErrorBuilder
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

func WithTranslator(translator validation.Translator) Option {
	return func(v validation.Validator) error {
		v.(*Validator).translator = translator
		return nil
	}
}

func WithDefaultLanguage(language string) Option {
	return func(v validation.Validator) error {
		v.(*Validator).defaultLanguage = language
		return nil
	}
}

func WithErrorBuilder(builder validation.ErrorBuilder) Option {
	return func(v validation.Validator) error {
		v.(*Validator).errorBuilder = builder
		return nil
	}
}

func (v *Validator) clone() *Validator {
	return &Validator{
		translator:      v.translator,
		defaultLanguage: v.defaultLanguage,
		errorBuilder:    v.errorBuilder,
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
	bag := validator.NewErrorBag(v2.translator)
	for key, validators := range validatorCtx.validators {
		for _, v := range validators {
			if err := v.Validate(ctx, v2); err != nil {
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
	return validator.NewError(code, message, params...).SetTranslator(v.translator)
}

func Validate(ctx context.Context, builders ...validation.ValidatorBuilder) *validator.ErrorBag {
	return _instance.Validate(ctx, builders...).(*validator.ErrorBag)
}

func ValidateIt(ctx context.Context, validatable validation.Validatable) validation.ErrorBag {
	return _instance.Validate(ctx, NewBuilder(validatable))
}

func Value[T any](ctx context.Context, value T, rules ...validation.Rule[T]) validation.ErrorBag {
	return Attribute(ctx, "value", value, rules...)
}

func Attribute[T any](ctx context.Context, attribute string, value T, rules ...validation.Rule[T]) validation.ErrorBag {
	var builders []validation.ValidatorBuilder
	for _, rule := range rules {
		builders = append(builders, NewBuilder(validator.ValidatableFunc(func(ctx context.Context, builder validation.ErrorBuilder) validation.Error {
			return rule.Validate(ctx, builder, value)
		})).SetAttribute(attribute))
	}
	return _instance.Validate(ctx, builders...)
}
