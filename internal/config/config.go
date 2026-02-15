//go:build linux

// Package config handles qfm configuration loading and management.
package config

// Config represents the complete qfm configuration.
type Config struct {
	Logging LogConfig  `mapstructure:"logging" yaml:"logging"`
	View    ViewConfig `mapstructure:"view" yaml:"view"`
}

// LogConfig defines logging settings.
type LogConfig struct {
	Level  string `mapstructure:"level"  yaml:"level"`
	Format string `mapstructure:"format" yaml:"format"`
}

// ViewConfig defines default view settings.
type ViewConfig struct {
	Mode       string `mapstructure:"mode"        yaml:"mode"`
	SortBy     string `mapstructure:"sort_by"     yaml:"sort_by"`
	ShowHidden bool   `mapstructure:"show_hidden" yaml:"show_hidden"`
}
