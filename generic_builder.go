package validation

import (
	"cmp"
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
)

// Blank returns a builder function that validates the given value is blank.
func Blank[T comparable](attribute string, value T) validation.ValidatorBuilder {
	return NewBuilder(validator.IsBlank[T]().SetValue(value)).SetAttribute(attribute)
}

// NotBlank returns a builder function that validates the given value is not blank.
func NotBlank[T comparable](attribute string, value T) validation.ValidatorBuilder {
	return NewBuilder(validator.IsNotBlank[T]().SetValue(value)).SetAttribute(attribute)
}

func In[T comparable](attribute string, value T, values ...T) validation.ValidatorBuilder {
	return NewBuilder(validator.IsIn(values...).SetValue(value)).SetAttribute(attribute)
}

func NotIn[T comparable](attribute string, value T, values ...T) validation.ValidatorBuilder {
	return NewBuilder(validator.IsNotIn(values...).SetValue(value)).SetAttribute(attribute)
}

func EqualTo[T comparable](attribute string, value T, other T) validation.ValidatorBuilder {
	return NewBuilder(validator.IsEqualTo(other).SetValue(value)).SetAttribute(attribute)
}

func NotEqualTo[T comparable](attribute string, value T, other T) validation.ValidatorBuilder {
	return NewBuilder(validator.IsNotEqualTo(other).SetValue(value)).SetAttribute(attribute)
}

func LessThan[T cmp.Ordered](attribute string, value T, other T) validation.ValidatorBuilder {
	return NewBuilder(validator.IsLessThan(other).SetValue(value)).SetAttribute(attribute)
}

func LessThanOrEqualTo[T cmp.Ordered](attribute string, value T, other T) validation.ValidatorBuilder {
	return NewBuilder(validator.IsLessThanOrEqualTo(other).SetValue(value)).SetAttribute(attribute)
}

func GreaterThan[T cmp.Ordered](attribute string, value T, other T) validation.ValidatorBuilder {
	return NewBuilder(validator.IsGreaterThan(other).SetValue(value)).SetAttribute(attribute)
}

func GreaterThanOrEqualTo[T cmp.Ordered](attribute string, value T, other T) validation.ValidatorBuilder {
	return NewBuilder(validator.IsGreaterThanOrEqualTo(other).SetValue(value)).SetAttribute(attribute)
}
