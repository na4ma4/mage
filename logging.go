package mage

import "log"

type LoggingFunc func(format string, v ...any)

var LogPrintF LoggingFunc = log.Printf

// Change the logger function while returning a function to revert to the previous logger function.
func ChangeLogger(f LoggingFunc) func() {
	oldLogger := LogPrintF
	LogPrintF = f
	return func() {
		LogPrintF = oldLogger
	}
}
