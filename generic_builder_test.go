package validation

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlank(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Blank("name", ""))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Blank("name", "John Doe"))
		assert.True(t, validated.Fails())
	})
}

func TestNotBlank(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotBlank("name", "John Doe"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotBlank("name", ""))
		assert.True(t, validated.Fails())
	})
}

func TestIn(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), In("id", 1, 1, 2, 3))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), In("id", 1, 2, 3, 4))
		assert.True(t, validated.Fails())
	})
}

func TestNotIn(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotIn("id", 1, 2, 3, 4))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotIn("id", 1, 1, 2, 3))
		assert.True(t, validated.Fails())
	})
}

func TestEqualTo(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), EqualTo("id", 1, 1))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), EqualTo("id", 1, 2))
		assert.True(t, validated.Fails())
	})
}

func TestNotEqualTo(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotEqualTo("id", 1, 2))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), NotEqualTo("id", 1, 1))
		assert.True(t, validated.Fails())
	})
}

func TestLessThan(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), LessThan("id", 1, 2))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), LessThan("id", 1, 1))
		assert.True(t, validated.Fails())
	})
}

func TestLessThanOrEqualTo(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), LessThanOrEqualTo("id", 1, 1))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), LessThanOrEqualTo("id", 1, 0))
		assert.True(t, validated.Fails())
	})
}

func TestGreaterThan(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), GreaterThan("id", 1, 0))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), GreaterThan("id", 1, 1))
		assert.True(t, validated.Fails())
	})
}

func TestGreaterThanOrEqualTo(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), GreaterThanOrEqualTo("id", 1, 1))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), GreaterThanOrEqualTo("id", 1, 2))
		assert.True(t, validated.Fails())
	})
}
