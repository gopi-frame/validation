package validator

import (
	"context"
	"fmt"
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/code"
	"github.com/gopi-frame/validation/message"
	"strconv"
)

func IsContainsKey[K comparable, V any](key K) MapRuleFunc[K, V] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value map[K]V) validation.Error {
		if _, ok := value[key]; !ok {
			return builder.BuildError(
				code.IsContainsKey,
				message.IsContainsKey,
				NewParam("key", strconv.Quote(fmt.Sprintf("%v", key))),
			)
		}
		return nil
	}
}

func IsNotContainsKey[K comparable, V any](key K) MapRuleFunc[K, V] {
	return func(ctx context.Context, builder validation.ErrorBuilder, value map[K]V) validation.Error {
		if _, ok := value[key]; ok {
			return builder.BuildError(
				code.IsNotContainsKey,
				message.IsNotContainsKey,
				NewParam("key", strconv.Quote(fmt.Sprintf("%v", key))),
			)
		}
		return nil
	}
}
