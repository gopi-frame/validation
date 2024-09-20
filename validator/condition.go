package validator

import (
	"context"
	"github.com/gopi-frame/contract/validation"
)

func If[T any](condition bool, rules ...validation.Rule[T]) RuleFunc[T] {
	if condition {
		return Group(rules...)
	}
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		return nil
	}
}
