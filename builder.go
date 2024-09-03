package validation

import (
	"context"
	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/validator"
	"strings"
)

type Path struct {
}

type Builder struct {
	validator validation.Validatable
	attribute string
	paths     []string
}

func NewBuilder(validator validation.Validatable) *Builder {
	return &Builder{
		validator: validator,
	}
}

func (b *Builder) SetAttribute(attribute string) validation.ValidatorBuilder {
	b.attribute = attribute
	return b
}

func (b *Builder) GetAttribute() string {
	return b.attribute
}

func (b *Builder) SetPath(paths ...string) validation.ValidatorBuilder {
	b.paths = paths
	return b
}

func (b *Builder) GetPath() []string {
	return b.paths
}

func (b *Builder) Build(ctx validation.ValidatorContext) {
	paths := b.paths
	if len(paths) == 0 {
		paths = []string{b.attribute}
	}
	ctx.AddValidate(strings.Join(paths, "."), validator.ValidatableFunc(func(ctx context.Context, builder validation.ErrorBuilder) validation.Error {
		if err := b.validator.Validate(ctx, builder); err != nil {
			err = err.AddParam(validator.NewParam("attribute", b.attribute))
			return err
		}
		return nil
	}))
}