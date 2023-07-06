package mage

import "os"

func IsDebug() bool {
	if os.Getenv("MAGEFILE_DEBUG") == "1" {
		return true
	}

	return false
}

func IsVerbose() bool {
	if os.Getenv("MAGEFILE_VERBOSE") == "1" {
		return true
	}

	if os.Getenv("MAGEFILE_DEBUG") == "1" {
		return true
	}

	return false
}
