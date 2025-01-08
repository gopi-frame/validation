package validation

import (
	"context"
	"testing"

	"github.com/gopi-frame/validation/code"
	"github.com/stretchr/testify/assert"
)

func TestIncludes(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var elements = []string{"a", "b", "c"}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Includes("elements", elements, "a", "b"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var elements = []string{"a", "b", "c"}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Includes("elements", elements, "a", "b", "d"))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "elements should include \"a\", \"b\", \"d\".", validated.GetError("elements", code.IsIncludes).Error())
		}
	})
}

func TestExcludes(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var elements = []string{"a", "b", "c"}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Excludes("elements", elements, "d"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		var elements = []string{"a", "b", "c"}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Excludes("elements", elements, "a"))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "elements should exclude \"a\".", validated.GetError("elements", code.IsExcludes).Error())
		}
	})
}

func TestUnique(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var elements = []string{"a", "b", "c"}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Unique("elements", elements))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		var elements = []string{"a", "b", "c", "a"}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Unique("elements", elements))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "elements should not contain duplicate elements.", validated.GetError("elements", code.IsUnique).Error())
		}
	})
}

func TestCount(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var elements = []string{"a", "b", "c"}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Count("elements", elements, 3))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var elements = []string{"a", "b", "c"}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Count("elements", elements, 2))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "elements should contain 2 element(s).", validated.GetError("elements", code.IsCount).Error())
		}
	})
}

func TestMinCount(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var elements = []string{"a", "b", "c"}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), MinCount("elements", elements, 2))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		var elements = []string{"a", "b", "c"}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), MinCount("elements", elements, 4))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "elements should contain at least 4 element(s).", validated.GetError("elements", code.IsMinCount).Error())
		}
	})
}

func TestMaxCount(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var elements = []string{"a", "b", "c"}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), MaxCount("elements", elements, 4))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		var elements = []string{"a", "b", "c"}
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), MaxCount("elements", elements, 2))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "elements should contain at most 2 element(s).", validated.GetError("elements", code.IsMaxCount).Error())
		}
	})
}
