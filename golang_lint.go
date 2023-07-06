package mage

import (
	"context"
	"fmt"
	"strings"

	"github.com/princjef/mageutil/bintool"
)

var golangLinter = bintool.Must(bintool.New(
	"golangci-lint{{.BinExt}}",
	GolangCILintVersion,
	"https://github.com/golangci/golangci-lint/releases/download/v{{.Version}}/golangci-lint-{{.Version}}-{{.GOOS}}-{{.GOARCH}}{{.ArchiveExt}}",
	bintool.WithFolder(fmt.Sprintf("./artifacts/bin/golangci-lint-%s", GolangCILintVersion)),
))

// Run golangci-lint on project.
func GoLint(ctx context.Context) error {
	LogPrintF("[golang] Linting ...")
	if err := golangLinter.Ensure(); err != nil {
		return err
	}

	args := []string{
		"run",
		"--sort-results",
		"--max-same-issues=0",
		"--max-issues-per-linter=0",
		"./...",
	}

	if err := golangLinter.Command(strings.Join(args, " ")).Run(); err != nil {
		return err
	}

	return nil
}
