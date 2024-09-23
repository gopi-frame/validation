package validation

import (
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
)

// Group returns a validator builder that validates the given value using the given rules.
// If the value is an implementation of the [validation.Validatable] interface, it will be validated first before the rules.
func Group[T any](attribute string, value T, rules ...validation.Rule[T]) validation.ValidatorBuilder {
	return NewBuilder(validator.Group(rules...).SetValue(value)).SetAttribute(attribute)
}
