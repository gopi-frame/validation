package validation

import (
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
)

func JSON(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsJSON().SetValue(s)).SetAttribute(attribute)
}

func JSONArray(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsJSONArray().SetValue(s)).SetAttribute(attribute)
}

func JSONObject(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsJSONObject().SetValue(s)).SetAttribute(attribute)
}

func JSONString(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsJSONString().SetValue(s)).SetAttribute(attribute)
}

func UUID(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsUUID().SetValue(s)).SetAttribute(attribute)
}

func UUIDV1(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsUUIDV1().SetValue(s)).SetAttribute(attribute)
}

func UUIDV2(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsUUIDV2().SetValue(s)).SetAttribute(attribute)
}

func UUIDV3(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsUUIDV3().SetValue(s)).SetAttribute(attribute)
}

func UUIDV4(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsUUIDV4().SetValue(s)).SetAttribute(attribute)
}

func UUIDV5(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsUUIDV5().SetValue(s)).SetAttribute(attribute)
}

func ULID(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsULID().SetValue(s)).SetAttribute(attribute)
}

func Base64(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsBase64().SetValue(s)).SetAttribute(attribute)
}

func Base32(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsBase32().SetValue(s)).SetAttribute(attribute)
}
