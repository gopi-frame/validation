package validation

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/gopi-frame/validation/code"
	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var data = `{"name":"gopi","age":18}`
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), JSON("value", data))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var data = `{"name":"gopi","age:18}`
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), JSON("value", data))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "value should be a valid JSON.", validated.GetError("value", code.IsJSON).Error())
		}
	})
}

func TestJSONArray(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var data = `[{"name":"gopi","age":18}]`
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), JSONArray("value", data))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var data = `{"name":"gopi","age":18}`
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), JSONArray("value", data))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "value should be a valid JSON array.", validated.GetError("value", code.IsJSONArray).Error())
		}
	})
}

func TestJSONObject(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var data = `{"name":"gopi","age":18}`
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), JSONObject("value", data))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var data = `[{"name":"gopi","age":18}]`
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), JSONObject("value", data))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid JSON object.", validated.GetError("value", code.IsJSONObject).Error())
	})
}

func TestJSONString(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var data = `"gopi"`
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), JSONString("value", data))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var data = `{"name":"gopi","age":18}`
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), JSONString("value", data))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid JSON string.", validated.GetError("value", code.IsJSONString).Error())
	})
}

func TestUUID(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var data = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), UUID("value", data))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var data = "6ba7b810-9dad-11d1-80b4-00c04fd430c81"
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), UUID("value", data))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid UUID.", validated.GetError("value", code.IsUUID).Error())
	})
}

func TestUUIDV1(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var data = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), UUIDv1("value", data))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var data = uuid.NewString()
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), UUIDv1("value", data))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid version 1 UUID.", validated.GetError("value", code.IsUUIDV1).Error())
	})
}

func TestUUIDV2(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var data = uuid.Must(uuid.NewDCEPerson()).String()
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), UUIDv2("value", data))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var data = uuid.NewString()
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), UUIDv2("value", data))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid version 2 UUID.", validated.GetError("value", code.IsUUIDV2).Error())
	})
}

func TestUUIDV3(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var data = uuid.NewMD5(uuid.NameSpaceDNS, []byte("gopi")).String()
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), UUIDv3("value", data))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var data = uuid.NewString()
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), UUIDv3("value", data))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid version 3 UUID.", validated.GetError("value", code.IsUUIDV3).Error())
	})
}

func TestUUIDV4(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var data = uuid.New().String()
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), UUIDv4("value", data))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var data = uuid.NewMD5(uuid.NameSpaceDNS, []byte("gopi")).String()
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), UUIDv4("value", data))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid version 4 UUID.", validated.GetError("value", code.IsUUIDV4).Error())
	})
}

func TestUUIDV5(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var data = uuid.NewSHA1(uuid.NameSpaceDNS, []byte("gopi")).String()
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), UUIDv5("value", data))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var data = uuid.NewString()
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), UUIDv5("value", data))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid version 5 UUID.", validated.GetError("value", code.IsUUIDV5).Error())
	})
}

func TestULID(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var data = "01J6H6TW0W1DH96AT9MEJJ5M32"
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), ULID("value", data))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var data = "01J6H6TW0W1DH96AT9MEJJ5M321"
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), ULID("value", data))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid ULID.", validated.GetError("value", code.IsULID).Error())
	})
}

func TestBase64(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var data = "Z29waQ=="
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Base64("value", data))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var data = "Z29waQ"
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Base64("value", data))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid base64 string.", validated.GetError("value", code.IsBase64).Error())
	})
}

func TestBase32(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var data = "JBSWY3DPEHPK3PXP"
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Base32("value", data))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var data = "JBSWY3DPEHPK3PXP1"
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Base32("value", data))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid base32 string.", validated.GetError("value", code.IsBase32).Error())
	})
}
