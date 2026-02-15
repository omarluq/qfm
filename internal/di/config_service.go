//go:build linux

package di

import (
	"fmt"
	"log/slog"
	"sync/atomic"

	"github.com/samber/do/v2"

	"github.com/omarluq/qfm/internal/config"
)

// ConfigService wraps the loaded configuration with hot-reload support.
type ConfigService struct {
	config atomic.Pointer[config.Config]
	path   string
}

// Get returns the current configuration via atomic load.
func (c *ConfigService) Get() *config.Config {
	return c.config.Load()
}

// Reload reloads configuration from disk.
func (c *ConfigService) Reload() error {
	cfg, err := config.Load(c.path)
	if err != nil {
		return fmt.Errorf("config reload failed: %w", err)
	}
	c.config.Store(cfg)
	slog.Info("config reloaded", "path", c.path)
	return nil
}

// Shutdown implements do.Shutdowner.
func (c *ConfigService) Shutdown() error {
	return nil
}

// NewConfig loads the configuration from the config path.
func NewConfig(i do.Injector) (*ConfigService, error) {
	path := do.MustInvokeNamed[string](i, ConfigPathKey)

	cfg, err := config.Load(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load config from %s: %w", path, err)
	}

	svc := &ConfigService{
		path: path,
	}
	svc.config.Store(cfg)

	return svc, nil
}
