package validation

import (
	"context"
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/code"
	"github.com/gopi-frame/validation/validator"
	"github.com/stretchr/testify/assert"
	"html/template"
	"strings"
	"testing"
)

var mockMessages = map[string]map[string]string{
	"en": {
		"attribute.name": "Name",
		code.IsNotBlank:  "{{.attribute}} should not be blank.",
	},
	"zh": {
		code.IsNotBlank: "{{.attribute}}不能为空。",
	},
}

type mockTranslator struct {
	defaultLanguage string
	language        string
	t               map[string]map[string]string
}

func (m *mockTranslator) T(key string, params map[string]any) string {
	sb := new(strings.Builder)
	var message string
	if str := m.t[m.language][key]; str != "" {
		message = str
	} else {
		message = m.t[m.defaultLanguage][key]
	}
	if err := template.Must(template.New("").Parse(message)).Execute(sb, params); err != nil {
		panic(err)
	} else {
		return sb.String()
	}
}

func (m *mockTranslator) P(key string, _ any, params map[string]any) string {
	return m.T(key, params)
}

func (m *mockTranslator) Locale(language string) validation.Translator {
	m2 := &mockTranslator{
		defaultLanguage: m.defaultLanguage,
		language:        language,
		t:               m.t,
	}
	return m2
}

func TestValidator(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			assert.FailNow(t, err.Error())
		} else {
			validated := v.Validate(context.Background(), Blank("name", ""))
			assert.Equal(t, false, validated.Fails())
		}
	})

	t.Run("block", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			assert.FailNow(t, err.Error())
		} else {
			validated := v.Validate(context.Background(), NotBlank("name", ""))
			assert.Equal(t, true, validated.HasError("name"))
			errs := validated.GetErrors("name")
			assert.Len(t, errs, 1)
			assert.Equal(t, "name should not be blank.", errs.Get(code.IsNotBlank).Error())
		}
	})

	t.Run("with translator", func(t *testing.T) {
		t.Run("not set default language", func(t *testing.T) {
			v, err := NewValidator(WithTranslator(&mockTranslator{t: mockMessages, defaultLanguage: "en"}))
			if err != nil {
				assert.FailNow(t, err.Error())
			} else {
				validated := v.Validate(context.Background(), NotBlank("name", ""))
				assert.Equal(t, true, validated.HasError("name"))
				errs := validated.GetErrors("name")
				assert.Len(t, errs, 1)
				assert.Equal(t, "Name should not be blank.", errs.Get(code.IsNotBlank).Error())
			}
		})

		t.Run("set default language", func(t *testing.T) {
			v, err := NewValidator(WithTranslator(&mockTranslator{t: mockMessages, defaultLanguage: "en"}), WithDefaultLanguage("zh"))
			if err != nil {
				assert.FailNow(t, err.Error())
			} else {
				validated := v.Validate(context.Background(), NotBlank("name", ""))
				assert.Equal(t, true, validated.HasError("name"))
				errs := validated.GetErrors("name")
				assert.Len(t, errs, 1)
				assert.Equal(t, "Name不能为空。", errs.Get(code.IsNotBlank).Error())
			}
		})
	})
}

type mockUser struct {
	Username string
	Password string
	Age      uint
	Tags     []string
}

func (m *mockUser) Validate(ctx context.Context, _ validation.ErrorBuilder) validation.Error {
	return Validate(ctx,
		Group("username", m.Username, validator.IsNotBlank[string](), validator.IsAsciiNumeric()),
		Group(
			"password",
			m.Password,
			validator.IsNotBlank[string](),
			validator.IsMinLength(6),
			validator.IsMaxLength(16),
			validator.IsMatch("^[a-zA-Z0-9!@#$%^&*()_+-=]{6,16}$"),
		),
		Group("age", m.Age, validator.IsGreaterThanOrEqualTo[uint](18)),
		Group("tags", m.Tags, validator.IsMaxCount[string](5), validator.IsUnique[string](), validator.Each[string](
			validator.IsNotBlank[string](),
			validator.IsMaxLength(16),
		)),
	)
}

func TestValidateIt(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var user = &mockUser{
			Username: "gopi",
			Password: "123456",
			Age:      20,
			Tags: []string{
				"golang",
				"web",
				"validation",
			},
		}
		validated := ValidateIt(context.Background(), user)
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		var user = &mockUser{
			Username: "gopi!!!",
			Password: "1234",
			Age:      10,
			Tags: []string{
				"golang",
				"golang",
				"validation!!!!!!!!!!!!",
				"web",
				"frame",
				"too-many",
			},
		}
		validated := ValidateIt(context.Background(), user)
		if assert.True(t, validated.Fails()) {
			assert.True(t, validated.HasError("username"))
			assert.True(t, validated.HasError("password"))
			assert.True(t, validated.HasError("age"))
			assert.True(t, validated.HasError("tags"))
			assert.True(t, validated.HasError("tags.2"))
		}
	})
}

func TestValidator_GlobalCustomErrorMessage(t *testing.T) {
	var password = "1234"
	validator, err := NewValidator(WithMessages(map[string]string{
		code.IsMinLength: "{{.attribute}}长度不能少于{{.min}}",
	}))
	if !assert.NoError(t, err) {
		assert.FailNow(t, err.Error())
	}
	validated := validator.Validate(context.Background(), MinLength("密码", password, 6).SetKey("password"))
	if assert.True(t, validated.Fails()) {
		assert.Equal(t, "密码长度不能少于6", validated.GetError("password", code.IsMinLength).Error())
	}
}

func TestValidate_PartCustomErrorMessage(t *testing.T) {
	var password = "1234"
	validated := Validate(context.Background(), MinLength("password", password, 6)).SetMessages(map[string]map[string]string{
		"password": {
			code.IsMinLength: "密码长度不能小于6",
		},
	})
	if assert.True(t, validated.Fails()) {
		assert.Equal(t, "密码长度不能小于6", validated.GetError("password", code.IsMinLength).Error())
	}
}
