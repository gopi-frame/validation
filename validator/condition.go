package validator

import (
	"context"
	"github.com/gopi-frame/contract/validation"
)

// If returns a validator builder that validates the given value using the given rules when the condition is true.
// if the value is an implementation of the [validation.Validatable] interface, it will be validated first before the rules.
func If[T any](condition bool, rules ...validation.Rule[T]) RuleFunc[T] {
	if condition {
		return Group(rules...)
	}
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		return nil
	}
}
