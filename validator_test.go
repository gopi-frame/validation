package validation

import (
	"context"
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/code"
	"github.com/gopi-frame/validation/translator"
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

	t.Run("with translations and default language", func(t *testing.T) {
		v, err := NewValidator(WithDefaultLanguage("zh-CN"))
		if err != nil {
			assert.FailNow(t, err.Error())
		} else {
			translator.RegisterTranslation("zh-CN", map[string]string{
				code.IsNotBlank: "{{.attribute}}不能为空。",
			})
			validated := v.Validate(context.Background(), NotBlank("name", ""), GreaterThan("age", 16, 18))
			assert.Equal(t, true, validated.HasError("name"))
			errs := validated.GetErrors("name")
			assert.Len(t, errs, 1)
			assert.Equal(t, "name不能为空。", errs.Get(code.IsNotBlank).Error())
			errs = validated.GetErrors("age")
			assert.Len(t, errs, 1)
			assert.Equal(t, "age should be greater than 18.", errs.Get(code.IsGreaterThan).Error())
		}
	})

	t.Run("with translations and context language", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			assert.FailNow(t, err.Error())
		} else {
			translator.RegisterTranslation("zh-CN", map[string]string{
				code.IsNotBlank: "{{.attribute}}不能为空。",
			})
			validated := v.Validate(BindLanguage(context.Background(), "zh-CN"), NotBlank("name", ""), GreaterThan("age", 16, 18))
			assert.Equal(t, true, validated.HasError("name"))
			errs := validated.GetErrors("name")
			assert.Len(t, errs, 1)
			assert.Equal(t, "name不能为空。", errs.Get(code.IsNotBlank).Error())
			errs = validated.GetErrors("age")
			assert.Len(t, errs, 1)
			assert.Equal(t, "age should be greater than 18.", errs.Get(code.IsGreaterThan).Error())
		}
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

func TestValidator_GlobalCustomErrorMessage(t *testing.T) {
	var password = "1234"
	v, err := NewValidator(WithMessages(map[string]string{
		code.IsMinLength: "{{.attribute}}长度不能少于{{.min}}",
	}))
	if !assert.NoError(t, err) {
		assert.FailNow(t, err.Error())
	}
	validated := v.Validate(context.Background(), MinLength("密码", password, 6).SetKey("password"))
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
