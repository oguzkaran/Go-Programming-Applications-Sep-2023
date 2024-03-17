package err

import (
	"fmt"
	"os"
)

func ExitFailure(message string) {
	_, _ = fmt.Fprintf(os.Stderr, "%s\n", message)
	os.Exit(1)
}

func ExitFailureError(messagePrefix string, err error) {
	_, _ = fmt.Fprintf(os.Stderr, "%s:%s\n", messagePrefix, err.Error())
	os.Exit(1)
}
