package mage

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// Makefile analog: clean the artifacts (and .makefiles).
func Clean(ctx context.Context) error {
	LogPrintF("[global] Cleaning ...")

	if IsVerbose() {
		LogPrintF("\tremoving directory ./.makefiles ...")
	}

	if err := os.RemoveAll("./.makefiles"); err != nil {
		return fmt.Errorf("unable to remove makefiles directory (./.makefiles): %w", err)
	}

	dirEntry, dirErr := os.ReadDir("./artifacts")
	if dirErr != nil {
		return fmt.Errorf("unable to read directory (./artifacts): %w", dirErr)
	}

	for _, entry := range dirEntry {
		if strings.EqualFold(entry.Name(), "data") || !entry.IsDir() {
			continue
		}

		dirName := fmt.Sprintf("./artifacts/%s", entry.Name())

		if IsVerbose() {
			LogPrintF("\tremoving directory %s ...", dirName)
		}

		if err := os.RemoveAll(dirName); err != nil {
			return fmt.Errorf("unable to remove artifacts directory (%s): %w", dirName, err)
		}
	}

	return nil
}
