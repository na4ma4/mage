package mage

var GolangCILintVersion = "1.56.2"

type ReleaseType string

const (
	DebugRelease      ReleaseType = "debug"
	ProductionRelease ReleaseType = "release"
)
