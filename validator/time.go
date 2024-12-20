package validator

import (
	"context"
	"strconv"
	"time"

	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/code"
	error2 "github.com/gopi-frame/validation/errpack"
	"github.com/gopi-frame/validation/is"
	"github.com/gopi-frame/validation/message"
)

func IsTime(layout string) StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if !is.Time(s, layout) {
			return builder.BuildError(code.IsTime, message.IsTime, error2.NewParam("layout", strconv.Quote(layout)))
		}
		return nil
	}
}

func IsANSIC() StringRuleFunc {
	return IsTime(time.ANSIC)
}

func IsUnixDate() StringRuleFunc {
	return IsTime(time.UnixDate)
}

func IsRubyDate() StringRuleFunc {
	return IsTime(time.RubyDate)
}

func IsRFC822() StringRuleFunc {
	return IsTime(time.RFC822)
}

func IsRFC822Z() StringRuleFunc {
	return IsTime(time.RFC822Z)
}

func IsRFC850() StringRuleFunc {
	return IsTime(time.RFC850)
}

func IsRFC1123() StringRuleFunc {
	return IsTime(time.RFC1123)
}

func IsRFC1123Z() StringRuleFunc {
	return IsTime(time.RFC1123Z)
}

func IsRFC3339() StringRuleFunc {
	return IsTime(time.RFC3339)
}

func IsRFC3339Nano() StringRuleFunc {
	return IsTime(time.RFC3339Nano)
}

func IsKitchen() StringRuleFunc {
	return IsTime(time.Kitchen)
}

func IsStamp() StringRuleFunc {
	return IsTime(time.Stamp)
}

func IsStampMilli() StringRuleFunc {
	return IsTime(time.StampMilli)
}

func IsStampMicro() StringRuleFunc {
	return IsTime(time.StampMicro)
}

func IsStampNano() StringRuleFunc {
	return IsTime(time.StampNano)
}

func IsDateTime() StringRuleFunc {
	return IsTime(time.DateTime)
}

func IsDateOnly() StringRuleFunc {
	return IsTime(time.DateOnly)
}

func IsTimeOnly() StringRuleFunc {
	return IsTime(time.TimeOnly)
}

func IsDuration() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if !is.Duration(s) {
			return builder.BuildError(code.IsDuration, message.IsDuration)
		}
		return nil
	}
}

func IsTimezone() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if !is.Timezone(s) {
			return builder.BuildError(code.IsTimezone, message.IsTimezone)
		}
		return nil
	}
}

func IsBefore(layout string, other time.Time) StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if t, err := time.ParseInLocation(layout, s, time.Local); err != nil {
			return builder.BuildError(code.IsTime, message.IsTime, error2.NewParam("format", layout))
		} else if !t.Before(other) {
			return builder.BuildError(
				code.IsBefore,
				message.IsBefore,
				error2.NewParam("time", strconv.Quote(other.Format(layout))),
			)
		}
		return nil
	}
}

func IsBeforeTZ(layout string, tz *time.Location, other time.Time) StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if t, err := time.ParseInLocation(layout, s, tz); err != nil {
			return builder.BuildError(code.IsTime, message.IsTime, error2.NewParam("format", layout))
		} else if !t.Before(other) {
			return builder.BuildError(
				code.IsBeforeTZ,
				message.IsBeforeTZ,
				error2.NewParam("time", strconv.Quote(other.Format(layout))),
				error2.NewParam("timezone", strconv.Quote(tz.String())),
			)
		}
		return nil
	}
}

func IsBeforeOrEqualTo(layout string, other time.Time) StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if t, err := time.ParseInLocation(layout, s, time.Local); err != nil {
			return builder.BuildError(code.IsTime, message.IsTime, error2.NewParam("format", layout))
		} else if t.After(other) {
			return builder.BuildError(
				code.IsBeforeOrEqualTo,
				message.IsBeforeOrEqualTo,
				error2.NewParam("time", strconv.Quote(other.Format(layout))),
			)
		}
		return nil
	}
}

func IsBeforeOrEqualToTZ(layout string, tz *time.Location, other time.Time) StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if t, err := time.ParseInLocation(layout, s, tz); err != nil {
			return builder.BuildError(code.IsTime, message.IsTime, error2.NewParam("format", layout))
		} else if t.After(other) {
			return builder.BuildError(
				code.IsBeforeOrEqualToTZ,
				message.IsBeforeOrEqualToTZ,
				error2.NewParam("time", strconv.Quote(other.Format(layout))),
				error2.NewParam("timezone", strconv.Quote(tz.String())),
			)
		}
		return nil
	}
}

func IsAfter(layout string, other time.Time) StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if t, err := time.ParseInLocation(layout, s, time.Local); err != nil {
			return builder.BuildError(code.IsTime, message.IsTime, error2.NewParam("format", layout))
		} else if !t.After(other) {
			return builder.BuildError(
				code.IsAfter,
				message.IsAfter,
				error2.NewParam("time", strconv.Quote(other.Format(layout))),
			)
		}
		return nil
	}
}

func IsAfterTZ(layout string, tz *time.Location, other time.Time) StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if t, err := time.ParseInLocation(layout, s, tz); err != nil {
			return builder.BuildError(code.IsTime, message.IsTime, error2.NewParam("format", layout))
		} else if !t.After(other) {
			return builder.BuildError(
				code.IsAfterTZ,
				message.IsAfterTZ,
				error2.NewParam("time", strconv.Quote(other.Format(layout))),
				error2.NewParam("timezone", strconv.Quote(tz.String())),
			)
		}
		return nil
	}
}

func IsAfterOrEqualTo(layout string, other time.Time) StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if t, err := time.ParseInLocation(layout, s, time.Local); err != nil {
			return builder.BuildError(code.IsTime, message.IsTime, error2.NewParam("format", layout))
		} else if t.Before(other) {
			return builder.BuildError(
				code.IsAfterOrEqualTo,
				message.IsAfterOrEqualTo,
				error2.NewParam("time", strconv.Quote(other.Format(layout))),
			)
		}
		return nil
	}
}

func IsAfterOrEqualToTZ(layout string, tz *time.Location, other time.Time) StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, s string) validation.Error {
		if t, err := time.ParseInLocation(layout, s, tz); err != nil {
			return builder.BuildError(code.IsTime, message.IsTime, error2.NewParam("format", layout))
		} else if t.Before(other) {
			return builder.BuildError(
				code.IsAfterOrEqualToTZ,
				message.IsAfterOrEqualToTZ,
				error2.NewParam("time", strconv.Quote(other.Format(layout))),
				error2.NewParam("timezone", strconv.Quote(tz.String())),
			)
		}
		return nil
	}
}
