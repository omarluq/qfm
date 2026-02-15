//go:build linux

package di

import (
	"log/slog"

	"github.com/samber/do/v2"
)

// RegisterSingletons registers all service providers as singletons.
func RegisterSingletons(i do.Injector) {
	do.Provide(i, NewConfig)
	do.Provide(i, NewLogger)
}

// EagerInit forces eager initialization of critical services.
func EagerInit(i do.Injector) error {
	if _, err := do.Invoke[*ConfigService](i); err != nil {
		return err
	}
	if _, err := do.Invoke[*slog.Logger](i); err != nil {
		return err
	}
	return nil
}
