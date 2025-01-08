package validation

import (
	"context"
	"testing"

	"github.com/gopi-frame/validation/code"
	"github.com/stretchr/testify/assert"
)

func TestIP(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), IP("client.ip", "127.0.0.1"))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), IP("client.ip", "127.0.0.1.1"))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "client.ip should be a valid IP address.", validated.GetError("client.ip", code.IsIP).Error())
		}
	})
}

func TestIPv4(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), IPv4("client.ip", "127.0.0.1"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), IPv4("client.ip", "127.0.0.1.1"))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "client.ip should be a valid IPv4 address.", validated.GetError("client.ip", code.IsIPv4).Error())
		}
	})
}

func TestIPv6(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), IPv6("client.ip", "::1"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), IPv6("client.ip", "::1.1"))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "client.ip should be a valid IPv6 address.", validated.GetError("client.ip", code.IsIPv6).Error())
		}
	})
}

func TestURL(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), URL("client.url", "http://localhost"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), URL("client.url", "htt.:.//@p://// /inval\\//id?1%aax%25/xα?"+string([]byte{0x7f})))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "client.url should be a valid URL.", validated.GetError("client.url", code.IsURL).Error())
		}
	})
}

func TestURLWithScheme(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), URLWithScheme("client.url", "http://localhost", "http"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), URLWithScheme("client.url", "http://localhost", "https"))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "client.url should be a valid URL with scheme https.", validated.GetError("client.url", code.IsURLWithScheme).Error())
		}
	})
}

func TestRequestURI(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RequestURI("client.url", "http://localhost"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RequestURI("client.url", "htt.:.//@p://// /inval\\//id?1%aax%25/xα?"+string([]byte{0x7f})))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "client.url should be a valid request URI.", validated.GetError("client.url", code.IsRequestURI).Error())
		}
	})
}

func TestURLQuery(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), URLQuery("client.url", "http://localhost"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), URLQuery("Query", "a=1;").SetKey("client.query"))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "Query should be a valid URL query string.", validated.GetError("client.query", code.IsURLQuery).Error())
		}
	})
}
