//go:build linux

// Package main is the entry point for qfm.
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"

	"github.com/omarluq/qfm/internal/versioninfo"
)

var rootCmd = &cobra.Command{
	Use:   "qfm",
	Short: "A blazing-fast file manager for Quickshell desktop environments",
	Long: `qfm is a lightweight Qt/QML file manager built in Go, designed for
visual consistency with Quickshell-based desktop shells. Linux-only. Wayland-only.`,
}

func main() {
	os.Exit(run())
}

func run() int {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	rootCmd.SetVersionTemplate("{{.Name}} {{.Version}}\n")

	fangOpts := []fang.Option{
		fang.WithVersion(versioninfo.String()),
	}

	if err := fang.Execute(ctx, rootCmd, fangOpts...); err != nil {
		return 1
	}
	return 0
}
