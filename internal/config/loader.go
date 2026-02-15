//go:build linux

package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Load reads and parses configuration using Viper.
// It searches XDG_CONFIG_HOME/qfm/ for config files.
func Load(path string) (*Config, error) {
	v := viper.New()

	if path != "" {
		v.SetConfigFile(path)
	} else {
		configDir := xdgConfigDir()
		v.SetConfigName("config")
		v.SetConfigType("json")
		v.AddConfigPath(configDir)
		v.AddConfigPath(".")
	}

	v.SetEnvPrefix("QFM")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	setDefaults(v)

	if err := v.ReadInConfig(); err != nil {
		// Config file is optional — use defaults if not found.
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &cfg, nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("logging.level", "info")
	v.SetDefault("logging.format", "console")
	v.SetDefault("view.mode", "list")
	v.SetDefault("view.show_hidden", false)
	v.SetDefault("view.sort_by", "name")
}

func xdgConfigDir() string {
	if dir := os.Getenv("XDG_CONFIG_HOME"); dir != "" {
		return filepath.Join(dir, "qfm")
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(".config", "qfm")
	}
	return filepath.Join(home, ".config", "qfm")
}
