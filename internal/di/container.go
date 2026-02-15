//go:build linux

// Package di provides dependency injection using samber/do v2.
package di

import (
	"context"
	"fmt"

	"github.com/samber/do/v2"
)

// ConfigPathKey is the named key for the config path string.
const ConfigPathKey = "config.path"

// Container wraps the do.Injector with qfm-specific configuration.
type Container struct {
	injector *do.RootScope
}

// NewContainer creates and configures the DI container.
func NewContainer(configPath string) (*Container, error) {
	injector := do.New()

	do.ProvideNamedValue(injector, ConfigPathKey, configPath)

	RegisterSingletons(injector)

	if err := EagerInit(injector); err != nil {
		return nil, err
	}

	return &Container{
		injector: injector,
	}, nil
}

// Injector returns the underlying do.Injector for direct access.
func (c *Container) Injector() *do.RootScope {
	return c.injector
}

// MustInvoke resolves a service from the container or panics.
func MustInvoke[T any](c *Container) T {
	return do.MustInvoke[T](c.injector)
}

// ShutdownWithContext gracefully shuts down with context for timeout control.
func (c *Container) ShutdownWithContext(ctx context.Context) error {
	done := make(chan *do.ShutdownReport, 1)
	go func() {
		done <- c.injector.ShutdownWithContext(ctx)
	}()

	select {
	case report := <-done:
		if report != nil && !report.Succeed {
			return fmt.Errorf("shutdown failed: %s", report.Error())
		}
		return nil
	case <-ctx.Done():
		return fmt.Errorf("shutdown timed out: %w", ctx.Err())
	}
}
