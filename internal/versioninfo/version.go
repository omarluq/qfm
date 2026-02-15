//go:build linux

// Package versioninfo provides build version information for qfm.
package versioninfo

import (
	"fmt"
	"runtime/debug"
)

var (
	// Version is the semantic version (injected at build time via ldflags).
	Version = "dev"
	// Commit is the git commit hash (injected at build time via ldflags).
	Commit = "none"
	// BuildDate is the build timestamp (injected at build time via ldflags).
	BuildDate = "unknown"
)

func init() {
	applyBuildInfoFallback()
}

// String returns formatted version information.
func String() string {
	if Version != "dev" {
		return fmt.Sprintf("%s (%s)", Version, shortCommit(Commit))
	}
	return fmt.Sprintf("dev (%s)", shortCommit(Commit))
}

func applyBuildInfoFallback() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	if Version == "dev" && info.Main.Version != "" && info.Main.Version != "(devel)" {
		Version = info.Main.Version
	}
	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			if Commit == "none" {
				Commit = setting.Value
			}
		case "vcs.time":
			if BuildDate == "unknown" {
				BuildDate = setting.Value
			}
		}
	}
}

func shortCommit(commit string) string {
	if len(commit) <= 7 {
		return commit
	}
	return commit[:7]
}
