package validator

import (
	"context"

	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/code"
	error2 "github.com/gopi-frame/validation/errpack"
	"github.com/gopi-frame/validation/is"
	"github.com/gopi-frame/validation/message"
)

func IsIP() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if !is.IP(s) {
			return builder.BuildError(code.IsIP, message.IsIP)
		}
		return nil
	}
}

func IsIP4() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if !is.IP4(s) {
			return builder.BuildError(code.IsIPv4, message.IsIPv4)
		}
		return nil
	}
}

func IsIP6() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if !is.IP6(s) {
			return builder.BuildError(code.IsIPv6, message.IsIPv6)
		}
		return nil
	}
}

func IsURL() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if !is.URL(s) {
			return builder.BuildError(code.IsURL, message.IsURL)
		}
		return nil
	}
}

func IsURLWithScheme(scheme string) StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if !is.URLWithScheme(s, scheme) {
			return builder.BuildError(code.IsURLWithScheme, message.IsURLWithScheme, error2.NewParam("scheme", scheme))
		}
		return nil
	}
}

func IsRequestURI() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if !is.RequestURI(s) {
			return builder.BuildError(code.IsRequestURI, message.IsRequestURI)
		}
		return nil
	}
}

func IsURLQuery() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if !is.URLQuery(s) {
			return builder.BuildError(code.IsURLQuery, message.IsURLQuery)
		}
		return nil
	}
}
