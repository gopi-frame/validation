package validator

import (
	"context"
	"github.com/gopi-frame/contract/validation"
)

func Group[T any](rules ...validation.Rule[T]) RuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		var bag = NewErrorBag(nil)
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
