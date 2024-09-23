package validation

import (
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
)

// If returns a validator builder that validates the given value using the given rules if the condition is true.
// if the value is an implementation of the [validation.Validatable] interface, it will be validated first before the rules.
func If[T any](condition bool, attribute string, value T, rules ...validation.Rule[T]) validation.ValidatorBuilder {
	return NewBuilder(validator.If(condition, rules...).SetValue(value)).SetAttribute(attribute)
}
