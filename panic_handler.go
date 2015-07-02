package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/wulonghui/cli/terminal"
)

func displayCrashDialog(err interface{}, commandArgs string, stackTrace string) {
	if err != nil && err != terminal.QuietPanic {
		switch err := err.(type) {
		case error:
			printCrashDialog(err.Error(), commandArgs, stackTrace)
		case string:
			printCrashDialog(err, commandArgs, stackTrace)
		default:
			printCrashDialog("An unexpected type of error", commandArgs, stackTrace)
		}
	}
}

func crashDialog(errorMessage string, commandArgs string, stackTrace string) string {
	formattedString := `
	Something completely unexpected happened. This is a bug in %s.
	Tell us that you ran this command:

		%s

	using this version of the CLI:

		%s

	and that this error occurred:

		%s

	and this stack trace:

	%s
`

	return fmt.Sprintf(formattedString, NAME, commandArgs, VERSION, errorMessage, stackTrace)
}

func printCrashDialog(errorMessage string, commandArgs string, stackTrace string) {
	ui.Say(crashDialog(errorMessage, commandArgs, stackTrace))
}

func generateBacktrace() string {
	stackByteCount := 0
	STACK_SIZE_LIMIT := 1024 * 1024
	var bytes []byte
	for stackSize := 1024; (stackByteCount == 0 || stackByteCount == stackSize) && stackSize < STACK_SIZE_LIMIT; stackSize = 2 * stackSize {
		bytes = make([]byte, stackSize)
		stackByteCount = runtime.Stack(bytes, true)
	}
	stackTrace := "\t" + strings.Replace(string(bytes), "\n", "\n\t", -1)
	return stackTrace
}

func handlePanics() {
	commandArgs := strings.Join(os.Args, " ")
	stackTrace := generateBacktrace()

	err := recover()
	displayCrashDialog(err, commandArgs, stackTrace)
	if err != nil {
		os.Exit(1)
	}
}
