package validator

import (
	"context"

	"github.com/gopi-frame/validation/message"

	"github.com/gopi-frame/contract/enum"
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/code"
)

func IsEnum[T enum.Enum]() RuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value T) validation.Error {
		var dummy T
		if !dummy.Contains(value) {
			return builder.BuildError(
				code.IsEnum,
				message.IsEnum,
			)
		}
		return nil
	}
}

func IsEnumString[T enum.Enum]() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		var dummy T
		if _, err := dummy.Parse(value); err != nil {
			return builder.BuildError(
				code.IsEnumString,
				message.IsEnumString,
			)
		}
		return nil
	}
}

func IsEnumValue[T enum.Enum]() RuleFunc[any] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value any) validation.Error {
		if e, ok := value.(T); ok {
			var dummy T
			if !dummy.Contains(e) {
				return builder.BuildError(
					code.IsEnumValue,
					message.IsEnumValue,
				)
			}
		}
		return nil
	}
}
