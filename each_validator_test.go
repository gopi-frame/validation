package validation

import (
	"context"
	"github.com/gopi-frame/validation/code"
	"github.com/gopi-frame/validation/validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEach(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var passwords = []string{
			"password123!",
			"password123!",
			"password123!",
		}
		validated := Validate(context.Background(), Each(
			"password",
			passwords,
			validator.IsNotBlank[string](),
			validator.IsMaxLength(16),
			validator.IsMatch("^[a-zA-Z0-9!@#$%^&*()_+-=]{6,16}$"),
		))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		var passwords = []string{
			"password123!",
			"password123!",
			"password123!*()123α",
		}
		validated := Validate(context.Background(), Each(
			"password",
			passwords,
			validator.IsNotBlank[string](),
			validator.IsMaxLength(16),
			validator.IsMatch("^[a-zA-Z0-9!@#$%^&*()_+-=]{6,16}$"),
		))
		if assert.True(t, validated.Fails()) {
			errs := validated.GetErrors("password.2")
			assert.Len(t, errs, 2)
			assert.True(t, errs.Has(code.IsMaxLength))
			assert.True(t, errs.Has(code.IsMatch))
		}
	})
}

func TestEach_Deep(t *testing.T) {
	var passwords = [][]string{
		{
			"password123!",
			"password123!",
			"password123!*()123α",
		},
	}
	validated := Validate(context.Background(), Each(
		"password",
		passwords,
		validator.Each[string](
			validator.IsNotBlank[string](),
			validator.IsMaxLength(16),
			validator.IsMatch("^[a-zA-Z0-9!@#$%^&*()_+-=]{6,16}$"),
		),
	))
	if assert.True(t, validated.Fails()) {
		errs := validated.GetErrors("password.0.2")
		assert.Len(t, errs, 2)
		assert.True(t, errs.Has(code.IsMaxLength))
		assert.True(t, errs.Has(code.IsMatch))
	}
}
