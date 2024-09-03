package validator

import (
	"context"
	"github.com/gopi-frame/contract/validation"
)

type ValidatableFunc func(ctx context.Context, builder validation.ErrorBuilder) validation.Error

func (f ValidatableFunc) Validate(ctx context.Context, builder validation.ErrorBuilder) validation.Error {
	return f(ctx, builder)
}

type RuleFunc[T any] func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error

func (f RuleFunc[T]) Validate(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
	return f(ctx, builder, value)
}

func (f RuleFunc[T]) SetValue(value T) validation.Validatable {
	return ValidatableFunc(func(ctx context.Context, builder validation.ErrorBuilder) validation.Error {
		return f(ctx, builder, value)
	})
}

type StringRuleFunc = RuleFunc[string]

type SliceRuleFunc[T any] RuleFunc[[]T]

func (f SliceRuleFunc[T]) SetValue(value []T) validation.Validatable {
	return ValidatableFunc(func(ctx context.Context, builder validation.ErrorBuilder) validation.Error {
		return f(ctx, builder, value)
	})
}

func (f SliceRuleFunc[T]) Validate(ctx context.Context, builder validation.ErrorBuilder, value []T) validation.Error {
	return f(ctx, builder, value)
}

type MapRuleFunc[K comparable, V any] RuleFunc[map[K]V]

func (f MapRuleFunc[K, V]) SetValue(value map[K]V) validation.Validatable {
	return ValidatableFunc(func(ctx context.Context, builder validation.ErrorBuilder) validation.Error {
		return f(ctx, builder, value)
	})
}

func (f MapRuleFunc[K, V]) Validate(ctx context.Context, builder validation.ErrorBuilder, value map[K]V) validation.Error {
	return f(ctx, builder, value)
}
