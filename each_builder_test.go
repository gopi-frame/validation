package validation

import (
	"context"
	"testing"

	"github.com/gopi-frame/validation/code"
	"github.com/gopi-frame/validation/validator"
	"github.com/stretchr/testify/assert"
)

func TestEach(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var passwords = []string{
			"password123!",
			"password123!",
			"password123!",
		}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Each(
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
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Each(
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
	v, err := NewValidator()
	if err != nil {
		t.Fatal(err)
	}
	validated := v.Validate(context.Background(), Each(
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

func TestEach_ValidatableImpl(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var users = []*mockUser{
			{
				Username: "gopi",
				Password: "password123!",
				Age:      25,
				Tags:     nil,
			},
			{
				Username: "gopi",
				Password: "password123!",
				Age:      25,
				Tags:     nil,
			},
			{
				Username: "gopi",
				Password: "password123!",
				Age:      25,
				Tags:     nil,
			},
		}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Each("Users", users))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var users = []*mockUser{
			{
				Username: "gopi",
				Password: "password123!",
				Age:      25,
				Tags:     nil,
			},
			{
				Username: "gopi",
				Password: "password123!",
				Age:      25,
				Tags:     nil,
			},
			{
				Username: "gopi",
				Password: "password123!*()123α",
				Age:      25,
				Tags:     nil,
			},
		}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Each("Users", users))
		if assert.True(t, validated.Fails()) {
			assert.True(t, validated.HasError("Users.2.password"))
			assert.True(t, validated.FailedAt("Users.2.password", code.IsMaxLength))
			assert.True(t, validated.FailedAt("Users.2.password", code.IsMatch))
		}
	})
}
