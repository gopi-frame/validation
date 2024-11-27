package message

const (
	IsBlank                = "{{.attribute}} should be blank."
	IsNotBlank             = "{{.attribute}} should not be blank."
	IsIn                   = "{{.attribute}} should be one of {{.values}}."
	IsNotIn                = "{{.attribute}} should not be one of {{.values}}."
	IsEqualTo              = "{{.attribute}} should be equal to {{.value}}."
	IsNotEqualTo           = "{{.attribute}} should not be equal to {{.value}}."
	IsLessThan             = "{{.attribute}} should be less than {{.value}}."
	IsLessThanOrEqualTo    = "{{.attribute}} should be less than or equal to {{.value}}."
	IsGreaterThan          = "{{.attribute}} should be greater than {{.value}}."
	IsGreaterThanOrEqualTo = "{{.attribute}} should be greater than or equal to {{.value}}."
)

const (
	IsLength           = "{{.attribute}} should have length {{.length}}."
	IsMinLength        = "{{.attribute}} should have length greater than or equal to {{.min}}."
	IsMaxLength        = "{{.attribute}} should have length less than or equal to {{.max}}."
	IsStartsWith       = "{{.attribute}} should start with {{.prefix}}."
	IsStartsWithAny    = "{{.attribute}} should start with one of {{.prefixes}}."
	IsNotStartsWith    = "{{.attribute}} should not start with {{.prefix}}."
	IsNotStartsWithAny = "{{.attribute}} should not start with any of {{.prefixes}}."
	IsEndsWith         = "{{.attribute}} should end with {{.suffix}}."
	IsEndsWithAny      = "{{.attribute}} should end with one of {{.suffixes}}."
	IsNotEndsWith      = "{{.attribute}} should not end with {{.suffix}}."
	IsNotEndsWithAny   = "{{.attribute}} should not end with any of {{.suffixes}}."
	IsMatch            = "{{.attribute}} should match {{.pattern}}."
	IsNotMatch         = "{{.attribute}} should not match {{.pattern}}."
	IsContains         = "{{.attribute}} should contain {{.substring}}."
	IsNotContains      = "{{.attribute}} should not contain {{.substring}}."
	IsUpper            = "{{.attribute}} should be uppercase."
	IsLower            = "{{.attribute}} should be lowercase."
	IsAlpha            = "{{.attribute}} should only contain letter."
	IsAlphaNumeric     = "{{.attribute}} should only contain letter and number."
	IsAlphaDash        = "{{.attribute}} should only contain letter, number and dash (-, _)."
	IsAscii            = "{{.attribute}} should only contain ascii letter (a-z, A-Z)."
	IsAsciiNumeric     = "{{.attribute}} should only contain ascii letter (a-z, A-Z) and number."
	IsAsciiDash        = "{{.attribute}} should only contain ascii letter (a-z, A-Z), number and dash (-, _)."
	IsNumber           = "{{.attribute}} should be a number."
	IsPositiveNumber   = "{{.attribute}} should be a positive number."
	IsNegativeNumber   = "{{.attribute}} should be a negative number."
	IsInteger          = "{{.attribute}} should be an integer."
	IsPositiveInteger  = "{{.attribute}} should be a positive integer."
	IsNegativeInteger  = "{{.attribute}} should be a negative integer."
	IsBinary           = "{{.attribute}} should be a binary number."
	IsOctal            = "{{.attribute}} should be an octal number."
	IsHexadecimal      = "{{.attribute}} should be a hexadecimal number."
	IsDecimal          = "{{.attribute}} should be a decimal number."
)

const (
	IsIncludes = "{{.attribute}} should include {{.values}}."
	IsExcludes = "{{.attribute}} should exclude {{.values}}."
	IsUnique   = "{{.attribute}} should not contain duplicate elements."
	IsCount    = "{{.attribute}} should contain {{.count}} element(s)."
	IsMinCount = "{{.attribute}} should contain at least {{.count}} element(s)."
	IsMaxCount = "{{.attribute}} should contain at most {{.count}} element(s)."
)

const (
	IsContainsKey    = "{{.attribute}} should contain key {{.key}}."
	IsNotContainsKey = "{{.attribute}} should not contain key {{.key}}."
)

const (
	IsTime              = "{{.attribute}} should be a valid time in format {{.layout}}."
	IsDuration          = "{{.attribute}} should be a valid duration."
	IsTimezone          = "{{.attribute}} should be a valid timezone."
	IsBefore            = "{{.attribute}} should be before {{.time}}."
	IsBeforeOrEqualTo   = "{{.attribute}} should be before or equal to {{.time}}."
	IsAfter             = "{{.attribute}} should be after {{.time}}."
	IsAfterOrEqualTo    = "{{.attribute}} should be after or equal to {{.time}}."
	IsBeforeTZ          = "{{.attribute}} in timezone {{.timezone}} should be before {{.time}}."
	IsAfterTZ           = "{{.attribute}} in timezone {{.timezone}} should be after {{.time}}."
	IsBeforeOrEqualToTZ = "{{.attribute}} in timezone {{.timezone}} should be before or equal to {{.time}}."
	IsAfterOrEqualToTZ  = "{{.attribute}} in timezone {{.timezone}} should be after or equal to {{.time}}."
)

const (
	IsJSON       = "{{.attribute}} should be a valid JSON."
	IsJSONArray  = "{{.attribute}} should be a valid JSON array."
	IsJSONObject = "{{.attribute}} should be a valid JSON object."
	IsJSONString = "{{.attribute}} should be a valid JSON string."
	IsUUID       = "{{.attribute}} should be a valid UUID."
	IsUUIDV1     = "{{.attribute}} should be a valid version 1 UUID."
	IsUUIDV2     = "{{.attribute}} should be a valid version 2 UUID."
	IsUUIDV3     = "{{.attribute}} should be a valid version 3 UUID."
	IsUUIDV4     = "{{.attribute}} should be a valid version 4 UUID."
	IsUUIDV5     = "{{.attribute}} should be a valid version 5 UUID."
	IsULID       = "{{.attribute}} should be a valid ULID."
	IsBase64     = "{{.attribute}} should be a valid base64 string."
	IsBase32     = "{{.attribute}} should be a valid base32 string."
)

const (
	IsIP            = "{{.attribute}} should be a valid IP address."
	IsIPv4          = "{{.attribute}} should be a valid IPv4 address."
	IsIPv6          = "{{.attribute}} should be a valid IPv6 address."
	IsURL           = "{{.attribute}} should be a valid URL."
	IsURLWithScheme = "{{.attribute}} should be a valid URL with scheme {{.scheme}}."
	IsRequestURI    = "{{.attribute}} should be a valid request URI."
	IsURLQuery      = "{{.attribute}} should be a valid URL query string."
)

const (
	IsEnum       = "{{.attribute}} should be a valid enum value."
	IsEnumString = "{{.attribute}} should be a valid enum value."
	IsEnumValue  = "{{.attribute}} should be a valid enum value."
)

const (
	IsPathExists    = "{{.attribute}} should be an existing path."
	IsPathNotExists = "{{.attribute}} should not be an existing path."
	IsPathDir       = "{{.attribute}} should be a directory."
	IsPathFile      = "{{.attribute}} should be a file."
	IsPathAbsolute  = "{{.attribute}} should be an absolute path."
	IsPathRelative  = "{{.attribute}} should be a relative path."
)
