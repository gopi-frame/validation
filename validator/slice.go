package validator

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/code"
	error2 "github.com/gopi-frame/validation/errpack"
	"github.com/gopi-frame/validation/message"
)

func IsIncludes[T comparable](elements ...T) SliceRuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, s []T) validation.Error {
		var includes []string
		for _, element := range elements {
			includes = append(includes, fmt.Sprintf("\"%v\"", element))
		}
		for _, e := range elements {
			if !slices.Contains(s, e) {
				return builder.BuildError(
					code.IsIncludes,
					message.IsIncludes,
					error2.NewParam("values", strings.Join(includes, ", ")),
				)
			}
		}
		return nil
	}
}

func IsExcludes[T comparable](elements ...T) SliceRuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, s []T) validation.Error {
		var excludes []string
		for _, element := range elements {
			excludes = append(excludes, fmt.Sprintf("\"%v\"", element))
		}
		for _, e := range elements {
			if slices.Contains(s, e) {
				return builder.BuildError(
					code.IsExcludes,
					message.IsExcludes,
					error2.NewParam("values", strings.Join(excludes, ", ")),
				)
			}
		}
		return nil
	}
}

func IsUnique[T comparable]() SliceRuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, s []T) validation.Error {
		var values = make(map[T]struct{})
		for _, e := range s {
			if _, ok := values[e]; ok {
				return builder.BuildError(code.IsUnique, message.IsUnique)
			}
			values[e] = struct{}{}
		}
		return nil
	}
}

func IsCount[T comparable](count int) SliceRuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, s []T) validation.Error {
		if len(s) != count {
			return builder.BuildError(code.IsCount, message.IsCount, error2.NewParam("count", count))
		}
		return nil
	}
}

func IsMinCount[T comparable](count int) SliceRuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, s []T) validation.Error {
		if len(s) < count {
			return builder.BuildError(code.IsMinCount, message.IsMinCount, error2.NewParam("count", count))
		}
		return nil
	}
}

func IsMaxCount[T comparable](count int) SliceRuleFunc[T] {
	return func(ctx context.Context, builder validation.ErrorBuilder, s []T) validation.Error {
		if len(s) > count {
			return builder.BuildError(code.IsMaxCount, message.IsMaxCount, error2.NewParam("count", count))
		}
		return nil
	}
}
