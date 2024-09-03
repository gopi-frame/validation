package validation

import (
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
	"time"
)

func Time(attribute string, value string, layout string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsTime(layout).SetValue(value)).SetAttribute(attribute)
}

func ANSIC(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsANSIC().SetValue(value)).SetAttribute(attribute)
}

func UnixDate(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsUnixDate().SetValue(value)).SetAttribute(attribute)
}

func RubyDate(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsRubyDate().SetValue(value)).SetAttribute(attribute)
}

func RFC822(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsRFC822().SetValue(value)).SetAttribute(attribute)
}

func RFC822Z(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsRFC822Z().SetValue(value)).SetAttribute(attribute)
}

func RFC850(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsRFC850().SetValue(value)).SetAttribute(attribute)
}

func RFC1123(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsRFC1123().SetValue(value)).SetAttribute(attribute)
}

func RFC1123Z(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsRFC1123Z().SetValue(value)).SetAttribute(attribute)
}

func RFC3339(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsRFC3339().SetValue(value)).SetAttribute(attribute)
}

func RFC3339Nano(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsRFC3339Nano().SetValue(value)).SetAttribute(attribute)
}

func Kitchen(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsKitchen().SetValue(value)).SetAttribute(attribute)
}

func Stamp(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsStamp().SetValue(value)).SetAttribute(attribute)
}

func StampMilli(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsStampMilli().SetValue(value)).SetAttribute(attribute)
}

func StampMicro(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsStampMicro().SetValue(value)).SetAttribute(attribute)
}

func StampNano(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsStampNano().SetValue(value)).SetAttribute(attribute)
}

func DateTime(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsDateTime().SetValue(value)).SetAttribute(attribute)
}

func DateOnly(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsDateOnly().SetValue(value)).SetAttribute(attribute)
}

func TimeOnly(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsTimeOnly().SetValue(value)).SetAttribute(attribute)
}

func Duration(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsDuration().SetValue(value)).SetAttribute(attribute)
}

func Timezone(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsTimezone().SetValue(value)).SetAttribute(attribute)
}

func Before(attribute string, value string, layout string, other time.Time) validation.ValidatorBuilder {
	return NewBuilder(validator.IsBefore(layout, other).SetValue(value)).SetAttribute(attribute)
}

func BeforeTZ(attribute string, value string, layout string, tz *time.Location, other time.Time) validation.ValidatorBuilder {
	return NewBuilder(validator.IsBeforeTZ(layout, tz, other).SetValue(value)).SetAttribute(attribute)
}

func BeforeOrEqualTo(attribute string, value string, layout string, other time.Time) validation.ValidatorBuilder {
	return NewBuilder(validator.IsBeforeOrEqualTo(layout, other).SetValue(value)).SetAttribute(attribute)
}

func BeforeOrEqualToTZ(attribute string, value string, layout string, tz *time.Location, other time.Time) validation.ValidatorBuilder {
	return NewBuilder(validator.IsBeforeOrEqualToTZ(layout, tz, other).SetValue(value)).SetAttribute(attribute)
}

func After(attribute string, value string, layout string, other time.Time) validation.ValidatorBuilder {
	return NewBuilder(validator.IsAfter(layout, other).SetValue(value)).SetAttribute(attribute)
}

func AfterTZ(attribute string, value string, layout string, tz *time.Location, other time.Time) validation.ValidatorBuilder {
	return NewBuilder(validator.IsAfterTZ(layout, tz, other).SetValue(value)).SetAttribute(attribute)
}

func AfterOrEqualTo(attribute string, value string, layout string, other time.Time) validation.ValidatorBuilder {
	return NewBuilder(validator.IsAfterOrEqualTo(layout, other).SetValue(value)).SetAttribute(attribute)
}

func AfterOrEqualToTZ(attribute string, value string, layout string, tz *time.Location, other time.Time) validation.ValidatorBuilder {
	return NewBuilder(validator.IsAfterOrEqualToTZ(layout, tz, other).SetValue(value)).SetAttribute(attribute)
}
