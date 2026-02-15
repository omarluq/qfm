//go:build linux

package di

import (
	"log/slog"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	slogzerolog "github.com/samber/slog-zerolog/v2"

	"github.com/omarluq/qfm/internal/config"
)

// NewLogger creates a zerolog-backed slog.Logger and sets it as the default.
func NewLogger(i do.Injector) (*slog.Logger, error) {
	cfgSvc := do.MustInvoke[*ConfigService](i)
	cfg := cfgSvc.Get()

	zl := newZerolog(cfg)

	logger := slog.New(
		slogzerolog.Option{
			Level:  parseLevel(cfg.Logging.Level),
			Logger: &zl,
		}.NewZerologHandler(),
	)

	slog.SetDefault(logger)

	return logger, nil
}

func newZerolog(cfg *config.Config) zerolog.Logger {
	switch strings.ToLower(cfg.Logging.Format) {
	case "json":
		return zerolog.New(os.Stdout).With().Timestamp().Logger()
	default:
		return zerolog.New(
			zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"},
		).With().Timestamp().Logger()
	}
}

func parseLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "trace", "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
