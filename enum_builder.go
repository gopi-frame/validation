package validation

import (
	"github.com/gopi-frame/contract/enum"
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
)

func IsEnum[T enum.Enum](attribute string, value T) validation.ValidatorBuilder {
	return NewBuilder(validator.IsEnum[T]().SetValue(value)).SetAttribute(attribute)
}

func IsEnumString[T enum.Enum](attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsEnumString[T]().SetValue(value)).SetAttribute(attribute)
}
