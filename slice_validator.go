package validation

import (
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
)

func Includes[T comparable](attribute string, s []T, elements ...T) validation.ValidatorBuilder {
	return NewBuilder(validator.IsIncludes(elements...).SetValue(s)).SetAttribute(attribute)
}

func Excludes[T comparable](attribute string, s []T, elements ...T) validation.ValidatorBuilder {
	return NewBuilder(validator.IsExcludes(elements...).SetValue(s)).SetAttribute(attribute)
}

func Unique[T comparable](attribute string, s []T) validation.ValidatorBuilder {
	return NewBuilder(validator.IsUnique[T]().SetValue(s)).SetAttribute(attribute)
}

func Count[T comparable](attribute string, s []T, count int) validation.ValidatorBuilder {
	return NewBuilder(validator.IsCount[T](count).SetValue(s)).SetAttribute(attribute)
}

func MinCount[T comparable](attribute string, s []T, count int) validation.ValidatorBuilder {
	return NewBuilder(validator.IsMinCount[T](count).SetValue(s)).SetAttribute(attribute)
}

func MaxCount[T comparable](attribute string, s []T, count int) validation.ValidatorBuilder {
	return NewBuilder(validator.IsMaxCount[T](count).SetValue(s)).SetAttribute(attribute)
}
