package validation

import (
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
)

func IP(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsIP().SetValue(value)).SetAttribute(attribute)
}

func IPv4(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsIP4().SetValue(value)).SetAttribute(attribute)
}

func IPv6(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsIP6().SetValue(value)).SetAttribute(attribute)
}

func URL(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsURL().SetValue(value)).SetAttribute(attribute)
}

func URLWithScheme(attribute string, value string, scheme string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsURLWithScheme(scheme).SetValue(value)).SetAttribute(attribute)
}

func RequestURI(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsRequestURI().SetValue(value)).SetAttribute(attribute)
}

func URLQuery(attribute string, value string) validation.ValidatorBuilder {
	return NewBuilder(validator.IsURLQuery().SetValue(value)).SetAttribute(attribute)
}
