package mage

var GolangCILintVersion = "1.53.3"

type ReleaseType string

const (
	DebugRelease      ReleaseType = "debug"
	ProductionRelease ReleaseType = "release"
)
