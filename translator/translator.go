package translator

import (
	"fmt"
	"strings"
	"sync"
	"text/template"

	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/code"
	"github.com/gopi-frame/validation/message"
)

var fallback = new(sync.Map)

var fallbackLanguage = "en"
var translations = new(sync.Map)

func RegisterTranslation(language string, messages map[string]string) {
	t := new(sync.Map)
	for c, m := range messages {
		fallback.Store(c, template.Must(template.New(c).Parse(m)))
	}
	translations.Store(language, t)
}

type Translator struct {
	messages *sync.Map
}

func New() *Translator {
	return &Translator{
		messages: fallback,
	}
}

func (t *Translator) T(key string, params map[string]any) string {
	if strings.HasPrefix(key, "attribute.") {
		key = key[len("attribute."):]
		return key
	}
	sb := new(strings.Builder)
	if v, ok := t.messages.Load(key); ok {
		if err := v.(*template.Template).Execute(sb, params); err != nil {
			panic(err)
		}
	} else if v, ok := fallback.Load(key); ok {
		if err := v.(*template.Template).Execute(sb, params); err != nil {
			panic(err)
		}
	} else {
		sb.WriteString(key)
		if len(params) > 0 {
			sb.WriteString(": ")
			sb.WriteString(fmt.Sprintf("%v", params))
		}
	}
	return sb.String()
}

func (t *Translator) P(key string, _ any, params map[string]any) string {
	return t.T(key, params)
}

func (t *Translator) Locale(language string) validation.Translator {
	if messages, ok := translations.Load(language); ok {
		return &Translator{messages: messages.(*sync.Map)}
	}
	return t
}

func init() {
	fallback.Store(code.IsBlank, template.Must(template.New(code.IsBlank).Parse(message.IsBlank)))
	fallback.Store(code.IsNotBlank, template.Must(template.New(code.IsNotBlank).Parse(message.IsNotBlank)))
	fallback.Store(code.IsIn, template.Must(template.New(code.IsIn).Parse(message.IsIn)))
	fallback.Store(code.IsNotIn, template.Must(template.New(code.IsNotIn).Parse(message.IsNotIn)))
	fallback.Store(code.IsEqualTo, template.Must(template.New(code.IsEqualTo).Parse(message.IsEqualTo)))
	fallback.Store(code.IsNotEqualTo, template.Must(template.New(code.IsNotEqualTo).Parse(message.IsNotEqualTo)))
	fallback.Store(code.IsLessThan, template.Must(template.New(code.IsLessThan).Parse(message.IsLessThan)))
	fallback.Store(code.IsLessThanOrEqualTo, template.Must(template.New(code.IsLessThanOrEqualTo).Parse(message.IsLessThanOrEqualTo)))
	fallback.Store(code.IsGreaterThan, template.Must(template.New(code.IsGreaterThan).Parse(message.IsGreaterThan)))
	fallback.Store(code.IsGreaterThanOrEqualTo, template.Must(template.New(code.IsGreaterThanOrEqualTo).Parse(message.IsGreaterThanOrEqualTo)))

	fallback.Store(code.IsLength, template.Must(template.New(code.IsLength).Parse(message.IsLength)))
	fallback.Store(code.IsMinLength, template.Must(template.New(code.IsMinLength).Parse(message.IsMinLength)))
	fallback.Store(code.IsMaxLength, template.Must(template.New(code.IsMaxLength).Parse(message.IsMaxLength)))
	fallback.Store(code.IsStartsWith, template.Must(template.New(code.IsStartsWith).Parse(message.IsStartsWith)))
	fallback.Store(code.IsStartsWithAny, template.Must(template.New(code.IsStartsWithAny).Parse(message.IsStartsWithAny)))
	fallback.Store(code.IsNotStartsWith, template.Must(template.New(code.IsNotStartsWith).Parse(message.IsNotStartsWith)))
	fallback.Store(code.IsNotStartsWithAny, template.Must(template.New(code.IsNotStartsWithAny).Parse(message.IsNotStartsWithAny)))
	fallback.Store(code.IsEndsWith, template.Must(template.New(code.IsEndsWith).Parse(message.IsEndsWith)))
	fallback.Store(code.IsEndsWithAny, template.Must(template.New(code.IsEndsWithAny).Parse(message.IsEndsWithAny)))
	fallback.Store(code.IsNotEndsWith, template.Must(template.New(code.IsNotEndsWith).Parse(message.IsNotEndsWith)))
	fallback.Store(code.IsNotEndsWithAny, template.Must(template.New(code.IsNotEndsWithAny).Parse(message.IsNotEndsWithAny)))
	fallback.Store(code.IsMatch, template.Must(template.New(code.IsMatch).Parse(message.IsMatch)))
	fallback.Store(code.IsNotMatch, template.Must(template.New(code.IsNotMatch).Parse(message.IsNotMatch)))
	fallback.Store(code.IsContains, template.Must(template.New(code.IsContains).Parse(message.IsContains)))
	fallback.Store(code.IsNotContains, template.Must(template.New(code.IsNotContains).Parse(message.IsNotContains)))
	fallback.Store(code.IsUpper, template.Must(template.New(code.IsUpper).Parse(message.IsUpper)))
	fallback.Store(code.IsLower, template.Must(template.New(code.IsLower).Parse(message.IsLower)))
	fallback.Store(code.IsAlpha, template.Must(template.New(code.IsAlpha).Parse(message.IsAlpha)))
	fallback.Store(code.IsAlphaNumeric, template.Must(template.New(code.IsAlphaNumeric).Parse(message.IsAlphaNumeric)))
	fallback.Store(code.IsAlphaDash, template.Must(template.New(code.IsAlphaDash).Parse(message.IsAlphaDash)))
	fallback.Store(code.IsAscii, template.Must(template.New(code.IsAscii).Parse(message.IsAscii)))
	fallback.Store(code.IsAsciiNumeric, template.Must(template.New(code.IsAsciiNumeric).Parse(message.IsAsciiNumeric)))
	fallback.Store(code.IsAsciiDash, template.Must(template.New(code.IsAsciiDash).Parse(message.IsAsciiDash)))
	fallback.Store(code.IsNumber, template.Must(template.New(code.IsNumber).Parse(message.IsNumber)))
	fallback.Store(code.IsPositiveNumber, template.Must(template.New(code.IsPositiveNumber).Parse(message.IsPositiveNumber)))
	fallback.Store(code.IsNegativeNumber, template.Must(template.New(code.IsNegativeNumber).Parse(message.IsNegativeNumber)))
	fallback.Store(code.IsInteger, template.Must(template.New(code.IsInteger).Parse(message.IsInteger)))
	fallback.Store(code.IsPositiveInteger, template.Must(template.New(code.IsPositiveInteger).Parse(message.IsPositiveInteger)))
	fallback.Store(code.IsNegativeInteger, template.Must(template.New(code.IsNegativeInteger).Parse(message.IsNegativeInteger)))
	fallback.Store(code.IsBinary, template.Must(template.New(code.IsBinary).Parse(message.IsBinary)))
	fallback.Store(code.IsOctal, template.Must(template.New(code.IsOctal).Parse(message.IsOctal)))
	fallback.Store(code.IsHexadecimal, template.Must(template.New(code.IsHexadecimal).Parse(message.IsHexadecimal)))
	fallback.Store(code.IsDecimal, template.Must(template.New(code.IsDecimal).Parse(message.IsDecimal)))

	fallback.Store(code.IsIncludes, template.Must(template.New(code.IsIncludes).Parse(message.IsIncludes)))
	fallback.Store(code.IsExcludes, template.Must(template.New(code.IsExcludes).Parse(message.IsExcludes)))
	fallback.Store(code.IsUnique, template.Must(template.New(code.IsUnique).Parse(message.IsUnique)))
	fallback.Store(code.IsCount, template.Must(template.New(code.IsCount).Parse(message.IsCount)))
	fallback.Store(code.IsMinCount, template.Must(template.New(code.IsMinCount).Parse(message.IsMinCount)))
	fallback.Store(code.IsMaxCount, template.Must(template.New(code.IsMaxCount).Parse(message.IsMaxCount)))

	fallback.Store(code.IsContainsKey, template.Must(template.New(code.IsContainsKey).Parse(message.IsContainsKey)))
	fallback.Store(code.IsNotContainsKey, template.Must(template.New(code.IsNotContainsKey).Parse(message.IsNotContainsKey)))

	fallback.Store(code.IsTime, template.Must(template.New(code.IsTime).Parse(message.IsTime)))
	fallback.Store(code.IsDuration, template.Must(template.New(code.IsDuration).Parse(message.IsDuration)))
	fallback.Store(code.IsTimezone, template.Must(template.New(code.IsTimezone).Parse(message.IsTimezone)))
	fallback.Store(code.IsBefore, template.Must(template.New(code.IsBefore).Parse(message.IsBefore)))
	fallback.Store(code.IsBeforeOrEqualTo, template.Must(template.New(code.IsBeforeOrEqualTo).Parse(message.IsBeforeOrEqualTo)))
	fallback.Store(code.IsAfter, template.Must(template.New(code.IsAfter).Parse(message.IsAfter)))
	fallback.Store(code.IsAfterOrEqualTo, template.Must(template.New(code.IsAfterOrEqualTo).Parse(message.IsAfterOrEqualTo)))
	fallback.Store(code.IsBeforeTZ, template.Must(template.New(code.IsBeforeTZ).Parse(message.IsBeforeTZ)))
	fallback.Store(code.IsBeforeOrEqualToTZ, template.Must(template.New(code.IsBeforeOrEqualToTZ).Parse(message.IsBeforeOrEqualToTZ)))
	fallback.Store(code.IsAfterTZ, template.Must(template.New(code.IsAfterTZ).Parse(message.IsAfterTZ)))
	fallback.Store(code.IsAfterOrEqualToTZ, template.Must(template.New(code.IsAfterOrEqualToTZ).Parse(message.IsAfterOrEqualToTZ)))

	fallback.Store(code.IsJSON, template.Must(template.New(code.IsJSON).Parse(message.IsJSON)))
	fallback.Store(code.IsJSONArray, template.Must(template.New(code.IsJSONArray).Parse(message.IsJSONArray)))
	fallback.Store(code.IsJSONObject, template.Must(template.New(code.IsJSONObject).Parse(message.IsJSONObject)))
	fallback.Store(code.IsJSONString, template.Must(template.New(code.IsJSONString).Parse(message.IsJSONString)))
	fallback.Store(code.IsUUID, template.Must(template.New(code.IsUUID).Parse(message.IsUUID)))
	fallback.Store(code.IsUUIDV1, template.Must(template.New(code.IsUUIDV1).Parse(message.IsUUIDV1)))
	fallback.Store(code.IsUUIDV2, template.Must(template.New(code.IsUUIDV2).Parse(message.IsUUIDV2)))
	fallback.Store(code.IsUUIDV3, template.Must(template.New(code.IsUUIDV3).Parse(message.IsUUIDV3)))
	fallback.Store(code.IsUUIDV4, template.Must(template.New(code.IsUUIDV4).Parse(message.IsUUIDV4)))
	fallback.Store(code.IsUUIDV5, template.Must(template.New(code.IsUUIDV5).Parse(message.IsUUIDV5)))
	fallback.Store(code.IsULID, template.Must(template.New(code.IsULID).Parse(message.IsULID)))
	fallback.Store(code.IsBase64, template.Must(template.New(code.IsBase64).Parse(message.IsBase64)))
	fallback.Store(code.IsBase32, template.Must(template.New(code.IsBase32).Parse(message.IsBase32)))

	fallback.Store(code.IsIP, template.Must(template.New(code.IsIP).Parse(message.IsIP)))
	fallback.Store(code.IsIPv4, template.Must(template.New(code.IsIPv4).Parse(message.IsIPv4)))
	fallback.Store(code.IsIPv6, template.Must(template.New(code.IsIPv6).Parse(message.IsIPv6)))
	fallback.Store(code.IsURL, template.Must(template.New(code.IsURL).Parse(message.IsURL)))
	fallback.Store(code.IsURLWithScheme, template.Must(template.New(code.IsURLWithScheme).Parse(message.IsURLWithScheme)))
	fallback.Store(code.IsRequestURI, template.Must(template.New(code.IsRequestURI).Parse(message.IsRequestURI)))
	fallback.Store(code.IsURLQuery, template.Must(template.New(code.IsURLQuery).Parse(message.IsURLQuery)))

	fallback.Store(code.IsEnum, template.Must(template.New(code.IsEnum).Parse(message.IsEnum)))
	fallback.Store(code.IsEnumString, template.Must(template.New(code.IsEnumString).Parse(message.IsEnumString)))
	fallback.Store(code.IsEnumValue, template.Must(template.New(code.IsEnumValue).Parse(message.IsEnumValue)))

	fallback.Store(code.IsPathExists, template.Must(template.New(code.IsPathExists).Parse(message.IsPathExists)))
	fallback.Store(code.IsPathNotExists, template.Must(template.New(code.IsPathNotExists).Parse(message.IsPathNotExists)))
	fallback.Store(code.IsPathFile, template.Must(template.New(code.IsPathFile).Parse(message.IsPathFile)))
	fallback.Store(code.IsPathDir, template.Must(template.New(code.IsPathDir).Parse(message.IsPathDir)))
	fallback.Store(code.IsPathAbsolute, template.Must(template.New(code.IsPathAbsolute).Parse(message.IsPathAbsolute)))
	fallback.Store(code.IsPathRelative, template.Must(template.New(code.IsPathRelative).Parse(message.IsPathRelative)))

	translations.Store(fallbackLanguage, fallback)
}
