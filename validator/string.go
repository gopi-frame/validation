package validator

import (
	"context"
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/code"
	error2 "github.com/gopi-frame/validation/error"
	"github.com/gopi-frame/validation/is"
	"github.com/gopi-frame/validation/message"
	"regexp"
	"strconv"
	"strings"
)

func IsLength(length int) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if len(value) != length {
			return errorBuilder.BuildError(code.IsLength, message.IsLength, error2.NewParam("length", strconv.Itoa(length)))
		}
		return nil
	}
}

func IsMinLength(length int) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if len(value) < length {
			return errorBuilder.BuildError(code.IsMinLength, message.IsMinLength, error2.NewParam("min", strconv.Itoa(length)))
		}
		return nil
	}
}

func IsMaxLength(length int) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if len(value) > length {
			return errorBuilder.BuildError(code.IsMaxLength, message.IsMaxLength, error2.NewParam("max", strconv.Itoa(length)))
		}
		return nil
	}
}

func IsStartsWith(prefix string) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !strings.HasPrefix(value, prefix) {
			return errorBuilder.BuildError(code.IsStartsWith, message.IsStartsWith, error2.NewParam("prefix", strconv.Quote(prefix)))
		}
		return nil
	}
}

func IsStartsWithAny(prefixes ...string) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		for _, prefix := range prefixes {
			if strings.HasPrefix(value, prefix) {
				return nil
			}
		}
		var ps []string
		for _, prefix := range prefixes {
			ps = append(ps, strconv.Quote(prefix))
		}
		return errorBuilder.BuildError(code.IsStartsWithAny, message.IsStartsWithAny, error2.NewParam("prefixes", strings.Join(ps, ", ")))
	}
}

func IsEndsWith(suffix string) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !strings.HasSuffix(value, suffix) {
			return errorBuilder.BuildError(code.IsEndsWith, message.IsEndsWith, error2.NewParam("suffix", strconv.Quote(suffix)))
		}
		return nil
	}
}

func IsEndsWithAny(suffixes ...string) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		for _, suffix := range suffixes {
			if strings.HasSuffix(value, suffix) {
				return nil
			}
		}
		var ss []string
		for _, suffix := range suffixes {
			ss = append(ss, strconv.Quote(suffix))
		}
		return errorBuilder.BuildError(code.IsEndsWithAny, message.IsEndsWithAny, error2.NewParam("suffixes", strings.Join(ss, ", ")))
	}
}

func IsNotStartsWith(prefix string) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if strings.HasPrefix(value, prefix) {
			return errorBuilder.BuildError(code.IsNotStartsWith, message.IsNotStartsWith, error2.NewParam("prefix", strconv.Quote(prefix)))
		}
		return nil
	}
}

func IsNotStartsWithAny(prefixes ...string) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		for _, prefix := range prefixes {
			if strings.HasPrefix(value, prefix) {
				var ps []string
				for _, prefix := range prefixes {
					ps = append(ps, strconv.Quote(prefix))
				}
				return errorBuilder.BuildError(code.IsNotStartsWithAny, message.IsNotStartsWithAny, error2.NewParam("prefixes", strings.Join(ps, ", ")))
			}
		}
		return nil
	}
}

func IsNotEndsWith(suffix string) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if strings.HasSuffix(value, suffix) {
			return errorBuilder.BuildError(code.IsNotEndsWith, message.IsNotEndsWith, error2.NewParam("suffix", strconv.Quote(suffix)))
		}
		return nil
	}
}

func IsNotEndsWithAny(suffixes ...string) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		for _, suffix := range suffixes {
			if strings.HasSuffix(value, suffix) {
				var ss []string
				for _, suffix := range suffixes {
					ss = append(ss, strconv.Quote(suffix))
				}
				return errorBuilder.BuildError(code.IsNotEndsWithAny, message.IsNotEndsWithAny, error2.NewParam("suffixes", strings.Join(ss, ", ")))
			}
		}
		return nil
	}
}

func IsMatch(pattern string) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !regexp.MustCompile(pattern).MatchString(value) {
			return errorBuilder.BuildError(code.IsMatch, message.IsMatch, error2.NewParam("pattern", strconv.Quote(pattern)))
		}
		return nil
	}
}

func IsNotMatch(pattern string) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if regexp.MustCompile(pattern).MatchString(value) {
			return errorBuilder.BuildError(code.IsNotMatch, message.IsNotMatch, error2.NewParam("pattern", strconv.Quote(pattern)))
		}
		return nil
	}
}

func IsContains(substring string) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !strings.Contains(value, substring) {
			return errorBuilder.BuildError(code.IsContains, message.IsContains, error2.NewParam("substring", strconv.Quote(substring)))
		}
		return nil
	}
}

func IsNotContains(substring string) StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if strings.Contains(value, substring) {
			return errorBuilder.BuildError(code.IsNotContains, message.IsNotContains, error2.NewParam("substring", strconv.Quote(substring)))
		}
		return nil
	}
}

func IsUpper() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.Upper(value) {
			return errorBuilder.BuildError(code.IsUpper, message.IsUpper)
		}
		return nil
	}
}

func IsLower() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.Lower(value) {
			return errorBuilder.BuildError(code.IsLower, message.IsLower)
		}
		return nil
	}
}

func IsAlpha() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.Alpha(value) {
			return errorBuilder.BuildError(code.IsAlpha, message.IsAlpha)
		}
		return nil
	}
}

func IsAlphaNumeric() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.AlphaNumeric(value) {
			return errorBuilder.BuildError(code.IsAlphaNumeric, message.IsAlphaNumeric)
		}
		return nil
	}
}

func IsAlphaDash() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.AlphaDash(value) {
			return errorBuilder.BuildError(code.IsAlphaDash, message.IsAlphaDash)
		}
		return nil
	}
}

func IsAscii() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.Ascii(value) {
			return errorBuilder.BuildError(code.IsAscii, message.IsAscii)
		}
		return nil
	}
}

func IsAsciiNumeric() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.AsciiNumeric(value) {
			return errorBuilder.BuildError(code.IsAsciiNumeric, message.IsAsciiNumeric)
		}
		return nil
	}
}

func IsAsciiDash() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.AsciiDash(value) {
			return errorBuilder.BuildError(code.IsAsciiDash, message.IsAsciiDash)
		}
		return nil
	}
}

func IsNumber() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.Number(value) {
			return errorBuilder.BuildError(code.IsNumber, message.IsNumber)
		}
		return nil
	}
}

func IsPositiveNumber() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.PositiveNumber(value) {
			return errorBuilder.BuildError(code.IsPositiveNumber, message.IsPositiveNumber)
		}
		return nil
	}
}

func IsNegativeNumber() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.NegativeNumber(value) {
			return errorBuilder.BuildError(code.IsNegativeNumber, message.IsNegativeNumber)
		}
		return nil
	}
}

func IsInteger() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.Integer(value) {
			return errorBuilder.BuildError(code.IsInteger, message.IsInteger)
		}
		return nil
	}
}

func IsPositiveInteger() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.PositiveInteger(value) {
			return errorBuilder.BuildError(code.IsPositiveInteger, message.IsPositiveInteger)
		}
		return nil
	}
}

func IsNegativeInteger() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.NegativeInteger(value) {
			return errorBuilder.BuildError(code.IsNegativeInteger, message.IsNegativeInteger)
		}
		return nil
	}
}

func IsDecimal() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.Decimal(value) {
			return errorBuilder.BuildError(code.IsDecimal, message.IsDecimal)
		}
		return nil
	}
}

func IsBinary() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.Binary(value) {
			return errorBuilder.BuildError(code.IsBinary, message.IsBinary)
		}
		return nil
	}
}

func IsOctal() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.Octal(value) {
			return errorBuilder.BuildError(code.IsOctal, message.IsOctal)
		}
		return nil
	}
}

func IsHexadecimal() StringRuleFunc {
	return func(ctx context.Context, errorBuilder validation.ErrorBuilder, value string) validation.Error {
		if !is.Hexadecimal(value) {
			return errorBuilder.BuildError(code.IsHexadecimal, message.IsHexadecimal)
		}
		return nil
	}
}
