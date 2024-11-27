package validation

import (
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
)

func PathExists(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsPathExists().SetValue(value)).SetAttribute(attribute)
}

func PathNotExists(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsPathNotExists().SetValue(value)).SetAttribute(attribute)
}

func PathFile(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsPathFile().SetValue(value)).SetAttribute(attribute)
}

func PathDir(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsPathDir().SetValue(value)).SetAttribute(attribute)
}

func PathAbsolute(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsPathAbsolute().SetValue(value)).SetAttribute(attribute)
}

func PathRelative(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsPathRelative().SetValue(value)).SetAttribute(attribute)
}
