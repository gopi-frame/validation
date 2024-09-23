package validation

import (
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
)

func Each[T any](attribute string, values []T, rules ...validation.Rule[T]) validation.ValidatorBuilder {
	return NewBuilder(validator.Each[T](rules...).SetValue(values)).SetAttribute(attribute)
}
