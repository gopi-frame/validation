package validation

import (
	"context"
	"testing"
	"time"

	"github.com/gopi-frame/validation/code"
	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Time("value", "2023-01-01 00:00:00", time.DateTime))
		assert.False(t, validated.Fails())
	})

	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Time("value", "2023-01-01 00:00:00", time.RFC822))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"02 Jan 06 15:04 MST\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestANSIC(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), ANSIC("value", "Sun Jan 01 00:00:00 2023"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), ANSIC("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"Mon Jan _2 15:04:05 2006\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestUnixDate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), UnixDate("value", "Mon Sep  2 09:30:30 CST 2024"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), UnixDate("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"Mon Jan _2 15:04:05 MST 2006\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestRubyDate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RubyDate("value", "Sun Jan 01 00:00:00 +0000 2023"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RubyDate("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"Mon Jan 02 15:04:05 -0700 2006\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestRFC822(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RFC822("value", "02 Sep 24 09:32 CST"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RFC822("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"02 Jan 06 15:04 MST\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestRFC822Z(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RFC822Z("value", "02 Sep 24 09:32 +0000"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RFC822Z("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"02 Jan 06 15:04 -0700\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestRFC850(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RFC850("value", "Sunday, 01-Jan-23 00:00:00 CST"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RFC850("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"Monday, 02-Jan-06 15:04:05 MST\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestRFC1123(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RFC1123("value", "Mon, 02 Sep 2024 09:33:37 CST"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RFC1123("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"Mon, 02 Jan 2006 15:04:05 MST\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestRFC1123Z(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RFC1123Z("value", "Mon, 02 Sep 2024 09:33:37 +0000"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RFC1123Z("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"Mon, 02 Jan 2006 15:04:05 -0700\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestRFC3339(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RFC3339("value", "2023-01-01T00:00:00Z"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RFC3339("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"2006-01-02T15:04:05Z07:00\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestRFC3339Nano(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RFC3339Nano("value", "2023-01-01T00:00:00.000000000Z"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), RFC3339Nano("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"2006-01-02T15:04:05.999999999Z07:00\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestKitchen(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Kitchen("value", "01:00AM"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Kitchen("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"3:04PM\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestStamp(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Stamp("value", "Sep  2 09:38:40"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Stamp("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"Jan _2 15:04:05\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestStampMilli(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), StampMilli("value", "Sep  2 09:38:40.000"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), StampMilli("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"Jan _2 15:04:05.000\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestStampMicro(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), StampMicro("value", "Sep  2 09:38:40.000000"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), StampMicro("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"Jan _2 15:04:05.000000\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestStampNano(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), StampNano("value", "Sep  2 09:38:40.000000000"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), StampNano("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"Jan _2 15:04:05.000000000\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestDateTime(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), DateTime("value", "2023-01-01 00:00:00"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), DateTime("value", "2023-01-01"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"2006-01-02 15:04:05\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestDateOnly(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), DateOnly("value", "2023-01-01"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), DateOnly("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"2006-01-02\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestTimeOnly(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), TimeOnly("value", "00:00:00"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), TimeOnly("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid time in format \"15:04:05\".", validated.GetError("value", code.IsTime).Error())
	})
}

func TestDuration(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Duration("value", "1h1m1s"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Duration("value", "2023-01-01 00:00:00"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid duration.", validated.GetError("value", code.IsDuration).Error())
	})
}

func TestTimezone(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Timezone("value", "Asia/Kolkata"))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Timezone("value", "Asia/unknown"))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be a valid timezone.", validated.GetError("value", code.IsTimezone).Error())
	})
}

func TestBefore(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Before("value", time.Now().Format(time.DateTime), time.DateTime, time.Now().Add(10*time.Second)))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		a := time.UnixMilli(1725244734000).Format(time.DateTime)
		b := time.UnixMilli(1725244733000)
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), Before("value", a, time.DateTime, b))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "value should be before \""+b.Format(time.DateTime)+"\".", validated.GetError("value", code.IsBefore).Error())
		}
	})
}

func TestBeforeTZ(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), BeforeTZ("value", time.Now().Format(time.DateTime), time.DateTime, time.UTC, time.Now().Add(10*time.Hour)))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		a := time.UnixMilli(1725244734000).In(time.UTC).Format(time.DateTime)
		b := time.UnixMilli(1725244734000).Add(-24 * time.Hour)
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), BeforeTZ("value", a, time.DateTime, time.UTC, b))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value in timezone \"UTC\" should be before \""+b.Format(time.DateTime)+"\".", validated.GetError("value", code.IsBeforeTZ).Error())
	})
}

func TestBeforeOrEqualTo(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), BeforeOrEqualTo("value", time.Now().Format(time.DateTime), time.DateTime, time.Now()))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		a := time.UnixMilli(1725244734000).Format(time.DateTime)
		b := time.UnixMilli(1725244733000)
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), BeforeOrEqualTo("value", a, time.DateTime, b))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be before or equal to \""+b.Format(time.DateTime)+"\".", validated.GetError("value", code.IsBeforeOrEqualTo).Error())
	})
}

func TestBeforeOrEqualToTZ(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(
			context.Background(),
			BeforeOrEqualToTZ(
				"value",
				time.Now().Format(time.DateTime),
				time.DateTime,
				time.UTC,
				time.Now().Add(time.Hour*24),
			),
		)
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		a := time.UnixMilli(1725244734000).In(time.UTC).Format(time.DateTime)
		b := time.UnixMilli(1725244734000).Add(-24 * time.Hour)
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(
			context.Background(),
			BeforeOrEqualToTZ(
				"value",
				a,
				time.DateTime,
				time.UTC,
				b,
			),
		)
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "value in timezone \"UTC\" should be before or equal to \""+b.Format(time.DateTime)+"\".", validated.GetError("value", code.IsBeforeOrEqualToTZ).Error())
		}
	})
}

func TestAfter(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), After("value", time.Now().Format(time.DateTime), time.DateTime, time.Now().Add(-10*time.Second)))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		a := time.UnixMilli(1725244734000).Format(time.DateTime)
		b := time.UnixMilli(1725244735000)
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), After("value", a, time.DateTime, b))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be after \""+b.Format(time.DateTime)+"\".", validated.GetError("value", code.IsAfter).Error())
	})
}

func TestAfterTZ(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), AfterTZ("value", time.Now().Format(time.DateTime), time.DateTime, time.UTC, time.Now().Add(-10*time.Hour)))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		a := time.UnixMilli(1725244734000).In(time.UTC).Format(time.DateTime)
		b := time.UnixMilli(1725244734000).Add(24 * time.Hour)
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), AfterTZ("value", a, time.DateTime, time.UTC, b))
		if assert.True(t, validated.Fails()) {
			assert.Equal(t, "value in timezone \"UTC\" should be after \""+b.Format(time.DateTime)+"\".", validated.GetError("value", code.IsAfterTZ).Error())
		}
	})
}

func TestAfterOrEqualTo(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), AfterOrEqualTo("value", time.Now().Format(time.DateTime), time.DateTime, time.Now().Add(-10*time.Second)))
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		a := time.UnixMilli(1725244734000).Format(time.DateTime)
		b := time.UnixMilli(1725244735000)
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(context.Background(), AfterOrEqualTo("value", a, time.DateTime, b))
		assert.True(t, validated.Fails())
		assert.Equal(t, "value should be after or equal to \""+b.Format(time.DateTime)+"\".", validated.GetError("value", code.IsAfterOrEqualTo).Error())
	})
}

func TestAfterOrEqualToTZ(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(
			context.Background(),
			AfterOrEqualToTZ(
				"value",
				time.Now().Format(time.DateTime),
				time.DateTime,
				time.UTC,
				time.Now().Add(-10*time.Hour),
			),
		)
		assert.False(t, validated.Fails())
	})
	t.Run("invalid", func(t *testing.T) {
		a := time.UnixMilli(1725244734000).In(time.UTC).Format(time.DateTime)
		b := time.UnixMilli(1725244734000).Add(24 * time.Hour)
		v, err := NewValidator()
		if err != nil {
			t.Fatal(err)
		}
		validated := v.Validate(
			context.Background(),
			AfterOrEqualToTZ(
				"value",
				a,
				time.DateTime,
				time.UTC,
				b,
			),
		)
		if assert.True(t, validated.Fails()) {
			assert.Equal(
				t,
				"value in timezone \"UTC\" should be after or equal to \""+b.Format(time.DateTime)+"\".",
				validated.GetError("value", code.IsAfterOrEqualToTZ).Error(),
			)
		}
	})
}
