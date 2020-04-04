package goish

import (
	"io"
	"os"
)

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)
const (
	debugCode = iota
	releaseCode
	testCode
)

var goishMode = debugCode
var modeName = DebugMode

// DefaultWriter ...
var DefaultWriter io.Writer = os.Stdout

// SetMode sets gin mode according to input string.
func SetMode(value string) {
	switch value {
	case DebugMode, "":
		goishMode = debugCode
	case ReleaseMode:
		goishMode = releaseCode
	case TestMode:
		goishMode = testCode
	default:
		panic("goish mode unknown: " + value)
	}
	if value == "" {
		value = DebugMode
	}
	modeName = value
}
