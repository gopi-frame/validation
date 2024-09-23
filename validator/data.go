package validator

import (
	"context"
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/code"
	"github.com/gopi-frame/validation/is"
	"github.com/gopi-frame/validation/message"
)

func IsJSON() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if !is.JSON(value) {
			return builder.BuildError(code.IsJSON, message.IsJSON)
		}
		return nil
	}
}

func IsJSONArray() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if !is.JSONArray(value) {
			return builder.BuildError(code.IsJSONArray, message.IsJSONArray)
		}
		return nil
	}
}

func IsJSONObject() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if !is.JSONObject(value) {
			return builder.BuildError(code.IsJSONObject, message.IsJSONObject)
		}
		return nil
	}
}

func IsJSONString() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if !is.JSONString(value) {
			return builder.BuildError(code.IsJSONString, message.IsJSONString)
		}
		return nil
	}
}

func IsUUID() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if !is.UUID(value) {
			return builder.BuildError(code.IsUUID, message.IsUUID)
		}
		return nil
	}
}

func IsUUIDv1() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if !is.UUIDV1(value) {
			return builder.BuildError(code.IsUUIDV1, message.IsUUIDV1)
		}
		return nil
	}
}

func IsUUIDv2() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if !is.UUIDV2(value) {
			return builder.BuildError(code.IsUUIDV2, message.IsUUIDV2)
		}
		return nil
	}
}

func IsUUIDv3() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if !is.UUIDV3(value) {
			return builder.BuildError(code.IsUUIDV3, message.IsUUIDV3)
		}
		return nil
	}
}

func IsUUIDv4() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if !is.UUIDV4(value) {
			return builder.BuildError(code.IsUUIDV4, message.IsUUIDV4)
		}
		return nil
	}
}

func IsUUIDv5() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if !is.UUIDV5(value) {
			return builder.BuildError(code.IsUUIDV5, message.IsUUIDV5)
		}
		return nil
	}
}

func IsULID() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if !is.ULID(value) {
			return builder.BuildError(code.IsULID, message.IsULID)
		}
		return nil
	}
}

func IsBase64() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if !is.Base64(value) {
			return builder.BuildError(code.IsBase64, message.IsBase64)
		}
		return nil
	}
}

func IsBase32() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if !is.Base32(value) {
			return builder.BuildError(code.IsBase32, message.IsBase32)
		}
		return nil
	}
}
