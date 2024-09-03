package validation

import (
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
)

// Length returns a builder function to check if a string has a specific length.
func Length(attribute string, value string, length int) validation.ValidatorBuilder {
	return NewBuilder(validator.IsLength(length).SetValue(value)).SetAttribute(attribute)
}

// MinLength returns a builder function to check if a string has a minimum length.
func MinLength(attribute string, value string, length int) validation.ValidatorBuilder {
	return NewBuilder(validator.IsMinLength(length).SetValue(value)).SetAttribute(attribute)
}

// MaxLength returns a builder function to check if a string has a maximum length.
func MaxLength(attribute string, value string, length int) validation.ValidatorBuilder {
	return NewBuilder(validator.IsMaxLength(length).SetValue(value)).SetAttribute(attribute)
}

// StartsWith returns a builder function to check if a string starts with a prefix.
func StartsWith(attribute string, value string, prefix string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsStartsWith(prefix).SetValue(value)).SetAttribute(attribute)
}

// StartsWithAny returns a builder function to check if a string starts with any of the given prefixes.
func StartsWithAny(attribute string, value string, prefixes ...string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsStartsWithAny(prefixes...).SetValue(value)).SetAttribute(attribute)
}

// EndsWith returns a builder function to check if a string ends with a suffix.
func EndsWith(attribute string, value string, suffix string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsEndsWith(suffix).SetValue(value)).SetAttribute(attribute)
}

// EndsWithAny returns a builder function to check if a string ends with any of the given suffixes.
func EndsWithAny(attribute string, value string, suffixes ...string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsEndsWithAny(suffixes...).SetValue(value)).SetAttribute(attribute)
}

// NotStartsWith returns a builder function to check if a string does not start with a prefix.
func NotStartsWith(attribute string, value string, prefix string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsNotStartsWith(prefix).SetValue(value)).SetAttribute(attribute)
}

// NotStartsWithAny returns a builder function to check if a string does not start with any of the given prefixes.
func NotStartsWithAny(attribute string, value string, prefixes ...string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsNotStartsWithAny(prefixes...).SetValue(value)).SetAttribute(attribute)
}

// NotEndsWith returns a builder function to check if a string does not end with a suffix.
func NotEndsWith(attribute string, value string, suffix string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsNotEndsWith(suffix).SetValue(value)).SetAttribute(attribute)
}

// NotEndsWithAny returns a builder function to check if a string does not end with any of the given suffixes.
func NotEndsWithAny(attribute string, value string, suffixes ...string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsNotEndsWithAny(suffixes...).SetValue(value)).SetAttribute(attribute)
}

// Match returns a builder function to check if a string matches a regular expression pattern.
func Match(attribute string, value string, pattern string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsMatch(pattern).SetValue(value)).SetAttribute(attribute)
}

// NotMatch returns a builder function to check if a string does not match a regular expression pattern.
func NotMatch(attribute string, value string, pattern string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsNotMatch(pattern).SetValue(value)).SetAttribute(attribute)
}

// Contains returns a builder function to check if a string contains a substring.
func Contains(attribute string, value string, substring string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsContains(substring).SetValue(value)).SetAttribute(attribute)
}

// NotContains returns a builder function to check if a string does not contain a substring.
func NotContains(attribute string, value string, substring string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsNotContains(substring).SetValue(value)).SetAttribute(attribute)
}

// Upper returns a builder function to check if a string is uppercase.
func Upper(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsUpper().SetValue(value)).SetAttribute(attribute)
}

// Lower returns a builder function to check if a string is lowercase.
func Lower(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsLower().SetValue(value)).SetAttribute(attribute)
}

// Alpha returns a builder function to check if a string contains only alphabetic characters.
func Alpha(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsAlpha().SetValue(value)).SetAttribute(attribute)
}

// AlphaNumeric returns a builder function to check if a string contains only alphanumeric characters.
func AlphaNumeric(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsAlphaNumeric().SetValue(value)).SetAttribute(attribute)
}

// AlphaDash returns a builder function to check if a string contains only alphanumeric characters and dashes.
func AlphaDash(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsAlphaDash().SetValue(value)).SetAttribute(attribute)
}

// Ascii returns a builder function to check if a string contains only ASCII characters.
func Ascii(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsAscii().SetValue(value)).SetAttribute(attribute)
}

// AsciiNumeric returns a builder function to check if a string contains only ASCII characters and numbers.
func AsciiNumeric(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsAsciiNumeric().SetValue(value)).SetAttribute(attribute)
}

// AsciiDash returns a builder function to check if a string contains only ASCII characters and dashes.
func AsciiDash(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsAsciiDash().SetValue(value)).SetAttribute(attribute)
}

// Number checks if the string is a number.
func Number(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsNumber().SetValue(value)).SetAttribute(attribute)
}

// PositiveNumber checks if the string is a positive number.
func PositiveNumber(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsPositiveNumber().SetValue(value)).SetAttribute(attribute)
}

// NegativeNumber checks if the string is a negative number.
func NegativeNumber(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsNegativeNumber().SetValue(value)).SetAttribute(attribute)
}

// Integer checks if the string is an integer.
func Integer(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsInteger().SetValue(value)).SetAttribute(attribute)
}

// PositiveInteger checks if the string is a positive integer.
func PositiveInteger(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsPositiveInteger().SetValue(value)).SetAttribute(attribute)
}

// NegativeInteger checks if the string is a negative integer.
func NegativeInteger(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsNegativeInteger().SetValue(value)).SetAttribute(attribute)
}

// Decimal if the string is a decimal.
func Decimal(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsDecimal().SetValue(value)).SetAttribute(attribute)
}

// Binary checks if the string is a binary.
func Binary(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsBinary().SetValue(value)).SetAttribute(attribute)
}

// Octal checks if the string is an octal.
func Octal(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsOctal().SetValue(value)).SetAttribute(attribute)
}

// Hexadecimal checks if the string is a hexadecimal.
func Hexadecimal(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsHexadecimal().SetValue(value)).SetAttribute(attribute)
}
