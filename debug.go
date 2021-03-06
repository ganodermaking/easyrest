package easyrest

import (
	"fmt"
	"strings"
)

func isDebugging() bool {
	return easyrestMode == debugCode
}

func debugPrint(format string, values ...interface{}) {
	if isDebugging() {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(DefaultWriter, "[EASYREST-debug] "+format, values...)
	}
}

func debugPrintRoute(method string, pattern string, handler HandlerFunc) {
	if isDebugging() {
		handlerName := nameOfFunction(handler)
		debugPrint("%-6s %-25s --> %s\n", method, pattern, handlerName)
	}
}
