package mage

import (
	"context"

	"github.com/magefile/mage/sh"
)

// Run golang unit tests on project.
func GoTest(ctx context.Context) error {
	LogPrintF("[golang] Testing ...")
	args := []string{
		"test",
	}

	if IsVerbose() {
		args = append(args, "-test.v=1")
	}

	if err := sh.RunV("go", append(args, "./...")...); err != nil {
		return err
	}

	return nil
}

// Manage your deps, or running package managers.
func GoModDownload(ctx context.Context) error {
	LogPrintF("[golang] Downloading Modules ...")
	return sh.RunV("go", "mod", "download")
}

// Build the golang project with a specified release type.
func GoBuild(releaseType ReleaseType) func(context.Context) error {
	return func(ctx context.Context) error {
		LogPrintF("[golang] Building Release (%s) ...", releaseType)
		return nil
	}
}
