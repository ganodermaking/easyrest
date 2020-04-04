package easyrest

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

var easyrestMode = debugCode
var modeName = DebugMode

// DefaultWriter ...
var DefaultWriter io.Writer = os.Stdout

// SetMode sets gin mode according to input string.
func SetMode(value string) {
	switch value {
	case DebugMode, "":
		easyrestMode = debugCode
	case ReleaseMode:
		easyrestMode = releaseCode
	case TestMode:
		easyrestMode = testCode
	default:
		panic("mode unknown: " + value)
	}
	if value == "" {
		value = DebugMode
	}
	modeName = value
}
