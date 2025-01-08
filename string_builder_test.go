package validation

import (
	"context"
	"testing"

	"github.com/gopi-frame/validation/code"
	"github.com/stretchr/testify/assert"
)

func TestStartsWith(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), StartsWith("value", "hello", "h"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), StartsWith("value", "hello", "a"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should start with \"a\".", validated.GetError("value", code.IsStartsWith).Error())
	})
}

func TestStartsWithAny(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), StartsWithAny("value", "hello", "a", "h"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), StartsWithAny("value", "hello", "a", "b"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should start with one of \"a\", \"b\".", validated.GetError("value", code.IsStartsWithAny).Error())
	})
}

func TestEndsWith(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), EndsWith("value", "hello", "o"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), EndsWith("value", "hello", "a"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should end with \"a\".", validated.GetError("value", code.IsEndsWith).Error())
	})
}

func TestEndsWithAny(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), EndsWithAny("value", "hello", "o", "a"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), EndsWithAny("value", "hello", "a", "b"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should end with one of \"a\", \"b\".", validated.GetError("value", code.IsEndsWithAny).Error())
	})
}

func TestNotStartsWith(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotStartsWith("value", "hello", "a"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotStartsWith("value", "hello", "h"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should not start with \"h\".", validated.GetError("value", code.IsNotStartsWith).Error())
	})
}

func TestNotStartsWithAny(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotStartsWithAny("value", "hello", "a", "b"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotStartsWithAny("value", "hello", "h", "a"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should not start with any of \"h\", \"a\".", validated.GetError("value", code.IsNotStartsWithAny).Error())
	})
}

func TestNotEndsWith(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotEndsWith("value", "hello", "a"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotEndsWith("value", "hello", "o"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should not end with \"o\".", validated.GetError("value", code.IsNotEndsWith).Error())
	})
}

func TestNotEndsWithAny(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotEndsWithAny("value", "hello", "a", "b"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotEndsWithAny("value", "hello", "o", "a"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should not end with any of \"o\", \"a\".", validated.GetError("value", code.IsNotEndsWithAny).Error())
	})
}

func TestMatch(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Match("value", "hello", "^h.*"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Match("value", "hello", "^a.*"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should match \"^a.*\".", validated.GetError("value", code.IsMatch).Error())
	})
}

func TestNotMatch(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotMatch("value", "hello", "^a.*"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotMatch("value", "hello", "^h.*"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should not match \"^h.*\".", validated.GetError("value", code.IsNotMatch).Error())
	})
}

func TestContains(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Contains("value", "hello", "ll"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Contains("value", "hello", "a"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should contain \"a\".", validated.GetError("value", code.IsContains).Error())
	})
}

func TestNotContains(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotContains("value", "hello", "a"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotContains("value", "hello", "ll"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should not contain \"ll\".", validated.GetError("value", code.IsNotContains).Error())
	})
}

func TestUpper(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Upper("value", "HELLO."))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Upper("value", "hello"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be uppercase.", validated.GetError("value", code.IsUpper).Error())
	})
}

func TestLower(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Lower("value", "hello"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Lower("value", "HELLO"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be lowercase.", validated.GetError("value", code.IsLower).Error())
	})
}

func TestAlpha(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Alpha("value", "hello"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Alpha("value", "hello1"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should only contain letter.", validated.GetError("value", code.IsAlpha).Error())
	})
}

func TestAlphaNumeric(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), AlphaNumeric("value", "hello1"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), AlphaNumeric("value", "hello1!"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should only contain letter and number.", validated.GetError("value", code.IsAlphaNumeric).Error())
	})
}

func TestAlphaDash(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), AlphaDash("value", "hello-1"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), AlphaDash("value", "hello1!"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should only contain letter, number and dash (-, _).", validated.GetError("value", code.IsAlphaDash).Error())
	})
}

func TestAscii(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Ascii("value", "hello"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Ascii("value", "hello\u0080"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should only contain ascii letter (a-z, A-Z).", validated.GetError("value", code.IsAscii).Error())
	})
}

func TestAsciiNumeric(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), AsciiNumeric("value", "123"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), AsciiNumeric("value", "123!"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should only contain ascii letter (a-z, A-Z) and number.", validated.GetError("value", code.IsAsciiNumeric).Error())
	})
}

func TestAsciiDash(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), AsciiDash("value", "hello-1"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), AsciiDash("value", "hello1!"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should only contain ascii letter (a-z, A-Z), number and dash (-, _).", validated.GetError("value", code.IsAsciiDash).Error())
	})
}

func TestNumber(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Number("value", "123"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Number("value", "123!"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a number.", validated.GetError("value", code.IsNumber).Error())
	})
}

func TestPositiveNumber(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), PositiveNumber("value", "123"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), PositiveNumber("value", "-123"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a positive number.", validated.GetError("value", code.IsPositiveNumber).Error())
	})
}

func TestNegativeNumber(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NegativeNumber("value", "-123"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NegativeNumber("value", "123"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a negative number.", validated.GetError("value", code.IsNegativeNumber).Error())
	})
}

func TestInteger(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Integer("value", "123"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Integer("value", "123.123"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be an integer.", validated.GetError("value", code.IsInteger).Error())
	})
}

func TestPositiveInteger(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), PositiveInteger("value", "123"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), PositiveInteger("value", "-123"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a positive integer.", validated.GetError("value", code.IsPositiveInteger).Error())
	})
}

func TestNegativeInteger(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NegativeInteger("value", "-123"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NegativeInteger("value", "123"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a negative integer.", validated.GetError("value", code.IsNegativeInteger).Error())
	})
}

func TestDecimal(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Decimal("value", "123.123"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Decimal("value", "0x123"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a decimal number.", validated.GetError("value", code.IsDecimal).Error())
	})
}

func TestBinary(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Binary("value", "0b0011"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Binary("value", "0x123!"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a binary number.", validated.GetError("value", code.IsBinary).Error())
	})
}

func TestOctal(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Octal("value", "0o777"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Octal("value", "0x123!"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be an octal number.", validated.GetError("value", code.IsOctal).Error())
	})
}

func TestHexadecimal(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Hexadecimal("value", "0x123"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Hexadecimal("value", "0o123!"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a hexadecimal number.", validated.GetError("value", code.IsHexadecimal).Error())
	})
}

func TestLength(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Length("value", "123456", 6))
		assert.False(t, validated.Fails())
	})

	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Length("value", "1234567", 6))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "value should have length 6.", validated.GetError("value", code.IsLength).Error())
		}
	})
}

func TestMinLength(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), MinLength("value", "123456", 6))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), MinLength("value", "1234", 6))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "value should have length greater than or equal to 6.", validated.GetError("value", code.IsMinLength).Error())
		}
	})
}

func TestMaxLength(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), MaxLength("value", "123456", 6))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), MaxLength("value", "1234567", 6))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "value should have length less than or equal to 6.", validated.GetError("value", code.IsMaxLength).Error())
		}
	})
}
