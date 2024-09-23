package validation

import (
	"context"
	"github.com/gopi-frame/validation/code"
	"github.com/gopi-frame/validation/validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroup(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		validated := Validate(context.Background(), Group(
			"password",
			"123456",
			validator.IsNotBlank[string](),
			validator.IsMinLength(6),
			validator.IsMaxLength(16),
			validator.IsMatch("^[a-zA-Z0-9!@#$%^&*()_+-=]{6,16}$"),
		))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		validated := Validate(context.Background(), Group(
			"password",
			"123456123456|||||||",
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

func TestGroup_ValidatableImpl(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var user = &mockUser{
			Username: "gopi",
			Password: "123456",
			Age:      25,
			Tags:     nil,
		}
		validated := Validate(context.Background(), Group(
			"User",
			user,
		))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var user = &mockUser{
			Username: "gopi",
			Password: "123456123456|||||||",
			Age:      25,
			Tags:     nil,
		}
		validated := Validate(context.Background(), Group(
			"User",
			user,
		))
		if assert.True(t, validated.Fails()) {
			assert.True(t, validated.HasError("User.password"))
			assert.True(t, validated.FailedAt("User.password", code.IsMaxLength))
			assert.True(t, validated.FailedAt("User.password", code.IsMatch))
		}
	})
}
