package validation

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gopi-frame/validation/contract"
)

// NewValidator new validator
func NewValidator() *Validator {
	validate := validator.New(validator.WithRequiredStructEnabled())
	en := en.New()
	zh := zh.New()
	translator := ut.New(en, en, zh)
	entranslator, _ := translator.GetTranslator(en.Locale())
	zhtranslator, _ := translator.GetTranslator(zh.Locale())
	entranslations.RegisterDefaultTranslations(validate, entranslator)
	zhtranslations.RegisterDefaultTranslations(validate, zhtranslator)
	return &Validator{
		Validate:   validate,
		Translator: translator,
	}
}

// Validator validator
type Validator struct {
	*validator.Validate
	Translator *ut.UniversalTranslator
}

// ValidateForm validate form
func (v *Validator) ValidateForm(form contract.Form) {
	err := v.Struct(form)
	if err == nil {
		for _, rule := range form.CustomRules() {
			if !rule.Validate(form) {
				break
			}
		}
		return
	}
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		panic(err)
	}
	if v.Translator != nil {
		translator, _ := v.Translator.GetTranslator(form.Locale())
		for _, err := range errs {
			form.AddError(err.Field(), err.Translate(translator))
		}
	} else {
		for _, err := range errs {
			form.AddError(err.Field(), err.Error())
		}
	}
}
