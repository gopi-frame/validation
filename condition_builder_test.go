package validation

import (
	"context"
	"testing"

	"github.com/gopi-frame/validation/code"
	"github.com/gopi-frame/validation/validator"
	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var value = "password123456"
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), If(
			true,
			"password",
			value,
			validator.IsNotBlank[string](),
			validator.IsMinLength(6),
			validator.IsMaxLength(16),
			validator.IsMatch("^[a-zA-Z0-9!@#$%^&*()_+-=]{6,16}$"),
		))
		assert.False(t, validated.Fails())
	})

	t.Run("skip", func(t *testing.T) {
		var value = "toolooooooooooooooooooooooooooog"
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), If(
			false,
			"password",
			value,
			validator.IsNotBlank[string](),
			validator.IsMinLength(6),
			validator.IsMaxLength(16),
			validator.IsMatch("^[a-zA-Z0-9!@#$%^&*()_+-=]{6,16}$"),
		))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var value = "toolooooooooooooooooooooooooooog"
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), If(
			true,
			"password",
			value,
			validator.IsNotBlank[string](),
			validator.IsMinLength(6),
			validator.IsMaxLength(16),
			validator.IsMatch("^[a-zA-Z0-9!@#$%^&*()_+-=]{6,16}$"),
		))
		if assert.True(t, validated.Fails()) {
			assert.True(t, validated.HasError("password"))
			assert.True(t, validated.FailedAt("password", code.IsMaxLength))
			assert.True(t, validated.FailedAt("password", code.IsMatch))
		}
	})
}
