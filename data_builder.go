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

func UUIDv1(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsUUIDv1().SetValue(s)).SetAttribute(attribute)
}

func UUIDv2(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsUUIDv2().SetValue(s)).SetAttribute(attribute)
}

func UUIDv3(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsUUIDv3().SetValue(s)).SetAttribute(attribute)
}

func UUIDv4(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsUUIDv4().SetValue(s)).SetAttribute(attribute)
}

func UUIDv5(attribute string, s string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsUUIDv5().SetValue(s)).SetAttribute(attribute)
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
