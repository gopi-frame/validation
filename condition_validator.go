package validation

import (
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
)

func If[T any](condition bool, attribute string, value T, rules ...validation.Rule[T]) validation.ValidatorBuilder {
	return NewBuilder(validator.If(condition, rules...).SetValue(value)).SetAttribute(attribute)
}
