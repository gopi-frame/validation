package validation

import (
	"reflect"

	"github.com/gopi-frame/contract/container"
)

// ServerProvider validation server provider
type ServerProvider struct{}

// Register register
func (s *ServerProvider) Register(c container.Container) {
	c.Bind(reflect.TypeFor[Validator]().Name(), func(c container.Container) any {
		return NewValidator()
	})
}
