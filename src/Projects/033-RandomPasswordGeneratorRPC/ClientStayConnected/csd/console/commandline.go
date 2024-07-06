package console

import (
	"fmt"
	"os"
)

func CheckLengthEquals(len, argsLen int, message string) {
	if argsLen != len {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", message)
		os.Exit(1)
	}
}
