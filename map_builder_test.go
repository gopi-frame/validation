package validation

import (
	"context"
	"testing"

	"github.com/gopi-frame/validation/code"
	"github.com/stretchr/testify/assert"
)

func TestContainsKey(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		m := map[string]string{
			"username": "gopi",
		}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), ContainsKey("user", m, "username"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		m := map[string]string{
			"username": "gopi",
		}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), ContainsKey("user", m, "password"))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "user should contain key \"password\".", validated.GetError("user", code.IsContainsKey).Error())
		}
	})
}

func TestNotContainsKey(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		m := map[string]string{
			"username": "gopi",
		}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotContainsKey("user", m, "password"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		m := map[string]string{
			"username": "gopi",
		}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotContainsKey("user", m, "username"))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "user should not contain key \"username\".", validated.GetError("user", code.IsNotContainsKey).Error())
		}
	})
}
