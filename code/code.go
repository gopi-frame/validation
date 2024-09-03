package code

// generic type validator codes
const (
	IsBlank                = "is_blank"
	IsNotBlank             = "is_not_blank"
	IsIn                   = "is_in"
	IsNotIn                = "is_not_in"
	IsEqualTo              = "is_equal"
	IsNotEqualTo           = "is_not_equal"
	IsLessThan             = "is_less_than"
	IsLessThanOrEqualTo    = "is_less_than_or_equal_to"
	IsGreaterThan          = "is_greater_than"
	IsGreaterThanOrEqualTo = "is_greater_than_or_equal_to"
)

// string type validator codes
const (
	IsLength           = "is_length"
	IsMinLength        = "is_min_length"
	IsMaxLength        = "is_max_length"
	IsStartsWith       = "is_starts_with"
	IsStartsWithAny    = "is_starts_with_any"
	IsNotStartsWith    = "is_not_starts_with"
	IsNotStartsWithAny = "is_not_starts_with_any"
	IsEndsWith         = "is_ends_with"
	IsEndsWithAny      = "is_ends_with_any"
	IsNotEndsWith      = "is_not_ends_with"
	IsNotEndsWithAny   = "is_not_ends_with_any"
	IsMatch            = "is_match"
	IsNotMatch         = "is_not_match"
	IsContains         = "is_contains"
	IsNotContains      = "is_not_contains"
	IsUpper            = "is_upper"
	IsLower            = "is_lower"
	IsAlpha            = "is_alpha"
	IsAlphaNumeric     = "is_alpha_numeric"
	IsAlphaDash        = "is_alpha_dash"
	IsAscii            = "is_ascii"
	IsAsciiNumeric     = "is_ascii_numeric"
	IsAsciiDash        = "is_ascii_dash"
	IsNumber           = "is_number"
	IsPositiveNumber   = "is_positive_number"
	IsNegativeNumber   = "is_negative_number"
	IsInteger          = "is_integer"
	IsPositiveInteger  = "is_positive_integer"
	IsNegativeInteger  = "is_negative_integer"
	IsBinary           = "is_binary"
	IsOctal            = "is_octal"
	IsHexadecimal      = "is_hexadecimal"
	IsDecimal          = "is_decimal"
)

// slice type validator codes
const (
	IsIncludes = "is_includes"
	IsExcludes = "is_excludes"
	IsUnique   = "is_unique"
	IsCount    = "is_count"
	IsMinCount = "is_min_count"
	IsMaxCount = "is_max_count"
)

// map type validator codes
const (
	IsContainsKey    = "is_contains_key"
	IsNotContainsKey = "is_not_contains_key"
)

// time validator codes
const (
	IsTime              = "is_time"
	IsDuration          = "is_duration"
	IsTimezone          = "is_timezone"
	IsBefore            = "is_before"
	IsBeforeOrEqualTo   = "is_before_or_equal_to"
	IsAfter             = "is_after"
	IsAfterOrEqualTo    = "is_after_or_equal_to"
	IsBeforeTZ          = "is_before_tz"
	IsAfterTZ           = "is_after_tz"
	IsBeforeOrEqualToTZ = "is_before_or_equal_to_tz"
	IsAfterOrEqualToTZ  = "is_after_or_equal_to_tz"
)

// data structure validator codes
const (
	IsJSON       = "is_json"
	IsJSONArray  = "is_json_array"
	IsJSONObject = "is_json_object"
	IsJSONString = "is_json_string"
	IsUUID       = "is_uuid"
	IsUUIDV1     = "is_uuid_v1"
	IsUUIDV2     = "is_uuid_v2"
	IsUUIDV3     = "is_uuid_v3"
	IsUUIDV4     = "is_uuid_v4"
	IsUUIDV5     = "is_uuid_v5"
	IsULID       = "is_ulid"
	IsBase64     = "is_base64"
	IsBase32     = "is_base32"
)

// net validator codes
const (
	IsIP            = "is_ip"
	IsIPv4          = "is_ipv4"
	IsIPv6          = "is_ipv6"
	IsURL           = "is_url"
	IsURLWithScheme = "is_url_with_schema"
	IsRequestURI    = "is_request_uri"
	IsURLQuery      = "is_url_query"
)
