package translator

import (
	"fmt"
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/code"
	"github.com/gopi-frame/validation/message"
	"strings"
	"sync"
	"text/template"
)

var cache = new(sync.Map)

func Register(key string, msg string) {
	cache.Store(key, template.Must(template.New(key).Parse(msg)))
}

func RegisterTpl(key string, t *template.Template) {
	cache.Store(key, t)
}

type Translator struct {
	cache map[string]*template.Template
}

func New() *Translator {
	return &Translator{}
}

func (t *Translator) T(key string, params map[string]any) string {
	if strings.HasPrefix(key, "attribute.") {
		key = key[len("attribute."):]
		return key
	}
	sb := new(strings.Builder)
	if v, ok := cache.Load(key); ok {
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

func (t *Translator) Locale(_ string) validation.Translator {
	return t
}

func init() {
	cache.Store(code.IsBlank, template.Must(template.New(code.IsBlank).Parse(message.IsBlank)))
	cache.Store(code.IsNotBlank, template.Must(template.New(code.IsNotBlank).Parse(message.IsNotBlank)))
	cache.Store(code.IsIn, template.Must(template.New(code.IsIn).Parse(message.IsIn)))
	cache.Store(code.IsNotIn, template.Must(template.New(code.IsNotIn).Parse(message.IsNotIn)))
	cache.Store(code.IsEqualTo, template.Must(template.New(code.IsEqualTo).Parse(message.IsEqualTo)))
	cache.Store(code.IsNotEqualTo, template.Must(template.New(code.IsNotEqualTo).Parse(message.IsNotEqualTo)))
	cache.Store(code.IsLessThan, template.Must(template.New(code.IsLessThan).Parse(message.IsLessThan)))
	cache.Store(code.IsLessThanOrEqualTo, template.Must(template.New(code.IsLessThanOrEqualTo).Parse(message.IsLessThanOrEqualTo)))
	cache.Store(code.IsGreaterThan, template.Must(template.New(code.IsGreaterThan).Parse(message.IsGreaterThan)))
	cache.Store(code.IsGreaterThanOrEqualTo, template.Must(template.New(code.IsGreaterThanOrEqualTo).Parse(message.IsGreaterThanOrEqualTo)))

	cache.Store(code.IsLength, template.Must(template.New(code.IsLength).Parse(message.IsLength)))
	cache.Store(code.IsMinLength, template.Must(template.New(code.IsMinLength).Parse(message.IsMinLength)))
	cache.Store(code.IsMaxLength, template.Must(template.New(code.IsMaxLength).Parse(message.IsMaxLength)))
	cache.Store(code.IsStartsWith, template.Must(template.New(code.IsStartsWith).Parse(message.IsStartsWith)))
	cache.Store(code.IsStartsWithAny, template.Must(template.New(code.IsStartsWithAny).Parse(message.IsStartsWithAny)))
	cache.Store(code.IsNotStartsWith, template.Must(template.New(code.IsNotStartsWith).Parse(message.IsNotStartsWith)))
	cache.Store(code.IsNotStartsWithAny, template.Must(template.New(code.IsNotStartsWithAny).Parse(message.IsNotStartsWithAny)))
	cache.Store(code.IsEndsWith, template.Must(template.New(code.IsEndsWith).Parse(message.IsEndsWith)))
	cache.Store(code.IsEndsWithAny, template.Must(template.New(code.IsEndsWithAny).Parse(message.IsEndsWithAny)))
	cache.Store(code.IsNotEndsWith, template.Must(template.New(code.IsNotEndsWith).Parse(message.IsNotEndsWith)))
	cache.Store(code.IsNotEndsWithAny, template.Must(template.New(code.IsNotEndsWithAny).Parse(message.IsNotEndsWithAny)))
	cache.Store(code.IsMatch, template.Must(template.New(code.IsMatch).Parse(message.IsMatch)))
	cache.Store(code.IsNotMatch, template.Must(template.New(code.IsNotMatch).Parse(message.IsNotMatch)))
	cache.Store(code.IsContains, template.Must(template.New(code.IsContains).Parse(message.IsContains)))
	cache.Store(code.IsNotContains, template.Must(template.New(code.IsNotContains).Parse(message.IsNotContains)))
	cache.Store(code.IsUpper, template.Must(template.New(code.IsUpper).Parse(message.IsUpper)))
	cache.Store(code.IsLower, template.Must(template.New(code.IsLower).Parse(message.IsLower)))
	cache.Store(code.IsAlpha, template.Must(template.New(code.IsAlpha).Parse(message.IsAlpha)))
	cache.Store(code.IsAlphaNumeric, template.Must(template.New(code.IsAlphaNumeric).Parse(message.IsAlphaNumeric)))
	cache.Store(code.IsAlphaDash, template.Must(template.New(code.IsAlphaDash).Parse(message.IsAlphaDash)))
	cache.Store(code.IsAscii, template.Must(template.New(code.IsAscii).Parse(message.IsAscii)))
	cache.Store(code.IsAsciiNumeric, template.Must(template.New(code.IsAsciiNumeric).Parse(message.IsAsciiNumeric)))
	cache.Store(code.IsAsciiDash, template.Must(template.New(code.IsAsciiDash).Parse(message.IsAsciiDash)))
	cache.Store(code.IsNumber, template.Must(template.New(code.IsNumber).Parse(message.IsNumber)))
	cache.Store(code.IsPositiveNumber, template.Must(template.New(code.IsPositiveNumber).Parse(message.IsPositiveNumber)))
	cache.Store(code.IsNegativeNumber, template.Must(template.New(code.IsNegativeNumber).Parse(message.IsNegativeNumber)))
	cache.Store(code.IsInteger, template.Must(template.New(code.IsInteger).Parse(message.IsInteger)))
	cache.Store(code.IsPositiveInteger, template.Must(template.New(code.IsPositiveInteger).Parse(message.IsPositiveInteger)))
	cache.Store(code.IsNegativeInteger, template.Must(template.New(code.IsNegativeInteger).Parse(message.IsNegativeInteger)))
	cache.Store(code.IsBinary, template.Must(template.New(code.IsBinary).Parse(message.IsBinary)))
	cache.Store(code.IsOctal, template.Must(template.New(code.IsOctal).Parse(message.IsOctal)))
	cache.Store(code.IsHexadecimal, template.Must(template.New(code.IsHexadecimal).Parse(message.IsHexadecimal)))
	cache.Store(code.IsDecimal, template.Must(template.New(code.IsDecimal).Parse(message.IsDecimal)))

	cache.Store(code.IsIncludes, template.Must(template.New(code.IsIncludes).Parse(message.IsIncludes)))
	cache.Store(code.IsExcludes, template.Must(template.New(code.IsExcludes).Parse(message.IsExcludes)))
	cache.Store(code.IsUnique, template.Must(template.New(code.IsUnique).Parse(message.IsUnique)))
	cache.Store(code.IsCount, template.Must(template.New(code.IsCount).Parse(message.IsCount)))
	cache.Store(code.IsMinCount, template.Must(template.New(code.IsMinCount).Parse(message.IsMinCount)))
	cache.Store(code.IsMaxCount, template.Must(template.New(code.IsMaxCount).Parse(message.IsMaxCount)))

	cache.Store(code.IsContainsKey, template.Must(template.New(code.IsContainsKey).Parse(message.IsContainsKey)))
	cache.Store(code.IsNotContainsKey, template.Must(template.New(code.IsNotContainsKey).Parse(message.IsNotContainsKey)))

	cache.Store(code.IsTime, template.Must(template.New(code.IsTime).Parse(message.IsTime)))
	cache.Store(code.IsDuration, template.Must(template.New(code.IsDuration).Parse(message.IsDuration)))
	cache.Store(code.IsTimezone, template.Must(template.New(code.IsTimezone).Parse(message.IsTimezone)))
	cache.Store(code.IsBefore, template.Must(template.New(code.IsBefore).Parse(message.IsBefore)))
	cache.Store(code.IsBeforeOrEqualTo, template.Must(template.New(code.IsBeforeOrEqualTo).Parse(message.IsBeforeOrEqualTo)))
	cache.Store(code.IsAfter, template.Must(template.New(code.IsAfter).Parse(message.IsAfter)))
	cache.Store(code.IsAfterOrEqualTo, template.Must(template.New(code.IsAfterOrEqualTo).Parse(message.IsAfterOrEqualTo)))
	cache.Store(code.IsBeforeTZ, template.Must(template.New(code.IsBeforeTZ).Parse(message.IsBeforeTZ)))
	cache.Store(code.IsBeforeOrEqualToTZ, template.Must(template.New(code.IsBeforeOrEqualToTZ).Parse(message.IsBeforeOrEqualToTZ)))
	cache.Store(code.IsAfterTZ, template.Must(template.New(code.IsAfterTZ).Parse(message.IsAfterTZ)))
	cache.Store(code.IsAfterOrEqualToTZ, template.Must(template.New(code.IsAfterOrEqualToTZ).Parse(message.IsAfterOrEqualToTZ)))

	cache.Store(code.IsJSON, template.Must(template.New(code.IsJSON).Parse(message.IsJSON)))
	cache.Store(code.IsJSONArray, template.Must(template.New(code.IsJSONArray).Parse(message.IsJSONArray)))
	cache.Store(code.IsJSONObject, template.Must(template.New(code.IsJSONObject).Parse(message.IsJSONObject)))
	cache.Store(code.IsJSONString, template.Must(template.New(code.IsJSONString).Parse(message.IsJSONString)))
	cache.Store(code.IsUUID, template.Must(template.New(code.IsUUID).Parse(message.IsUUID)))
	cache.Store(code.IsUUIDV1, template.Must(template.New(code.IsUUIDV1).Parse(message.IsUUIDV1)))
	cache.Store(code.IsUUIDV2, template.Must(template.New(code.IsUUIDV2).Parse(message.IsUUIDV2)))
	cache.Store(code.IsUUIDV3, template.Must(template.New(code.IsUUIDV3).Parse(message.IsUUIDV3)))
	cache.Store(code.IsUUIDV4, template.Must(template.New(code.IsUUIDV4).Parse(message.IsUUIDV4)))
	cache.Store(code.IsUUIDV5, template.Must(template.New(code.IsUUIDV5).Parse(message.IsUUIDV5)))
	cache.Store(code.IsULID, template.Must(template.New(code.IsULID).Parse(message.IsULID)))
	cache.Store(code.IsBase64, template.Must(template.New(code.IsBase64).Parse(message.IsBase64)))
	cache.Store(code.IsBase32, template.Must(template.New(code.IsBase32).Parse(message.IsBase32)))

	cache.Store(code.IsIP, template.Must(template.New(code.IsIP).Parse(message.IsIP)))
	cache.Store(code.IsIPv4, template.Must(template.New(code.IsIPv4).Parse(message.IsIPv4)))
	cache.Store(code.IsIPv6, template.Must(template.New(code.IsIPv6).Parse(message.IsIPv6)))
	cache.Store(code.IsURL, template.Must(template.New(code.IsURL).Parse(message.IsURL)))
	cache.Store(code.IsURLWithScheme, template.Must(template.New(code.IsURLWithScheme).Parse(message.IsURLWithScheme)))
	cache.Store(code.IsRequestURI, template.Must(template.New(code.IsRequestURI).Parse(message.IsRequestURI)))
	cache.Store(code.IsURLQuery, template.Must(template.New(code.IsURLQuery).Parse(message.IsURLQuery)))
}
