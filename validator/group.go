package validator

import (
	"context"

	"github.com/gopi-frame/contract/validation"
	error2 "github.com/gopi-frame/validation/errpack"
)

// Group returns a validator builder that validates the given value using the given rules.
// if the value is an implementation of Validatable, it will be validated first before the rules.
func Group[T any](rules ...validation.Rule[T]) RuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		var bag = error2.NewBag()
		if v, ok := any(value).(validation.Validatable); ok {
			if err := v.Validate(ctx, builder); err != nil {
				bag.AddError("", err)
			}
		}
		for _, rule := range rules {
			if err := rule.Validate(ctx, builder, value); err != nil {
				bag.AddError("", err)
			}
		}
		if bag.Fails() {
			return bag
		}
		return nil
	}
}
