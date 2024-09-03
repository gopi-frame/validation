package validator

import (
	"cmp"
	"context"
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/code"
	"github.com/gopi-frame/validation/message"
)

func IsBlank[T comparable]() RuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		if value != *new(T) {
			return builder.BuildError(code.IsBlank, message.IsBlank)
		}
		return nil
	}
}

func IsNotBlank[T comparable]() RuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		if value == *new(T) {
			return builder.BuildError(code.IsNotBlank, message.IsNotBlank)
		}
		return nil
	}
}

func IsIn[T comparable](values ...T) RuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		for _, v := range values {
			if value == v {
				return nil
			}
		}
		return builder.BuildError(code.IsIn, message.IsIn, NewParam("values", values))
	}
}

func IsNotIn[T comparable](values ...T) RuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		for _, v := range values {
			if value == v {
				return builder.BuildError(code.IsNotIn, message.IsNotIn, NewParam("values", values))
			}
		}
		return nil
	}
}

func IsEqualTo[T comparable](other T) RuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		if value != other {
			return builder.BuildError(code.IsEqualTo, message.IsEqualTo, NewParam("value", other))
		}
		return nil
	}
}

func IsNotEqualTo[T comparable](other T) RuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		if value == other {
			return builder.BuildError(code.IsNotEqualTo, message.IsNotEqualTo, NewParam("value", other))
		}
		return nil
	}
}

func IsLessThan[T cmp.Ordered](other T) RuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		if value >= other {
			return builder.BuildError(code.IsLessThan, message.IsLessThan, NewParam("value", other))
		}
		return nil
	}
}

func IsLessThanOrEqualTo[T cmp.Ordered](other T) RuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		if value > other {
			return builder.BuildError(code.IsLessThanOrEqualTo, message.IsLessThanOrEqualTo, NewParam("value", other))
		}
		return nil
	}
}

func IsGreaterThan[T cmp.Ordered](other T) RuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		if value <= other {
			return builder.BuildError(code.IsGreaterThan, message.IsGreaterThan, NewParam("value", other))
		}
		return nil
	}
}

func IsGreaterThanOrEqualTo[T cmp.Ordered](other T) RuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		if value < other {
			return builder.BuildError(code.IsGreaterThanOrEqualTo, message.IsGreaterThanOrEqualTo, NewParam("value", other))
		}
		return nil
	}
}
