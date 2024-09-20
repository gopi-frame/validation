package validator

import (
	"context"
	"github.com/gopi-frame/contract/validation"
	error2 "github.com/gopi-frame/validation/error"
)

func Group[T any](rules ...validation.Rule[T]) RuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		var bag = error2.NewErrorBag(nil)
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
