package validator

import (
	"context"
	"github.com/gopi-frame/contract/validation"
	"strconv"
)

func Each[T any](rules ...validation.Rule[T]) RuleFunc[[]T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, values []T) validation.Error {
		var bag = NewErrorBag(nil)
		for index, value := range values {
			for _, rule := range rules {
				if err := rule.Validate(ctx, builder, value); err != nil {
					bag.AddError(strconv.Itoa(index), err)
				}
			}
		}
		if bag.Fails() {
			return bag
		}
		return nil
	}
}
