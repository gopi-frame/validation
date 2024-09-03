package validation

import (
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
)

func Group[T any](attribute string, value T, rules ...validation.Rule[T]) validation.ValidatorBuilder {
	return NewBuilder(validator.Group(rules...).SetValue(value)).SetAttribute(attribute)
}
