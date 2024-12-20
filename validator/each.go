package validator

import (
	"context"
	"strconv"

	"github.com/gopi-frame/contract/validation"
	error2 "github.com/gopi-frame/validation/errpack"
)

// Each returns a validator builder that validates the given value's each element using the given rules.
func Each[T any](rules ...validation.Rule[T]) RuleFunc[[]T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, values []T) validation.Error {
		var bag = error2.NewBag()
		for index, value := range values {
			if err := Group(rules...).Validate(ctx, builder, value); err != nil {
				bag.AddError(strconv.Itoa(index), err)
			}
		}
		if bag.Fails() {
			return bag
		}
		return nil
	}
}
