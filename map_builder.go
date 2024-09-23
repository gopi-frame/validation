package validation

import (
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
)

func ContainsKey[K comparable, V any](attribute string, m map[K]V, key K) validation.ValidatorBuilder {
	return NewBuilder(validator.IsContainsKey[K, V](key).SetValue(m)).SetAttribute(attribute)
}

func NotContainsKey[K comparable, V any](attribute string, m map[K]V, key K) validation.ValidatorBuilder {
	return NewBuilder(validator.IsNotContainsKey[K, V](key).SetValue(m)).SetAttribute(attribute)
}
