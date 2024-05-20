package validation

import (
	"reflect"

	"github.com/gopi-frame/contract/container"
	"github.com/gopi-frame/contract/support"
	"github.com/gopi-frame/contract/validation"
)

// ServerProvider validation server provider
type ServerProvider struct {
	support.ServerProvider
}

// Validator get validator instance
func (s *ServerProvider) Validator(c container.Container) *Validator {
	return c.Get("validator").(*Validator)
}

// Register register
func (s *ServerProvider) Register(c container.Container) {
	c.Bind("validator", func(c container.Container) any {
		return NewValidator()
	})
	c.Alias("validator", reflect.TypeFor[Validator]().String())
	c.Alias("validator", reflect.TypeFor[validation.Engine]().String())
}
