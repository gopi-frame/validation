package validator

import (
	"context"
	"os"
	"path/filepath"

	"github.com/gopi-frame/validation/message"

	"github.com/gopi-frame/validation/code"

	"github.com/gopi-frame/contract/validation"
	"github.com/gopi-frame/validation/errpack"
)

func IsPathExists() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if _, err := os.Stat(value); err != nil {
			return builder.BuildError(
				code.IsPathExists,
				message.IsPathExists,
				errpack.NewParam("value", value),
			)
		}
		return nil
	}
}

func IsPathNotExists() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if _, err := os.Stat(value); err == nil {
			return builder.BuildError(
				code.IsPathNotExists,
				message.IsPathNotExists,
				errpack.NewParam("value", value),
			)
		}
		return nil
	}
}

func IsPathFile() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if info, err := os.Stat(value); err != nil {
			return builder.BuildError(
				code.IsPathFile,
				message.IsPathFile,
				errpack.NewParam("value", value),
			)
		} else if info.IsDir() {
			return builder.BuildError(
				code.IsPathFile,
				message.IsPathFile,
				errpack.NewParam("value", value),
			)
		}
		return nil
	}
}

func IsPathDir() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if info, err := os.Stat(value); err != nil {
			return builder.BuildError(
				code.IsPathDir,
				message.IsPathDir,
				errpack.NewParam("value", value),
			)
		} else if !info.IsDir() {
			return builder.BuildError(
				code.IsPathDir,
				code.IsPathDir,
				errpack.NewParam("value", value),
			)
		}
		return nil
	}
}

func IsPathAbsolute() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if !filepath.IsAbs(value) {
			return builder.BuildError(
				code.IsPathAbsolute,
				message.IsPathAbsolute,
				errpack.NewParam("value", value),
			)
		}
		return nil
	}
}

func IsPathRelative() StringRuleFunc {
	return func(ctx context.Context, builder validation.ErrorBuilder, value string) validation.Error {
		if filepath.IsAbs(value) {
			return builder.BuildError(
				code.IsPathRelative,
				message.IsPathRelative,
				errpack.NewParam("value", value),
			)
		}
		return nil
	}
}
