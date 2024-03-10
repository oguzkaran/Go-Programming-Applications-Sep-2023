package cerror

import (
	"fmt"
	"os"
)

func ExitFailure(message string) {
	_, _ = fmt.Fprintf(os.Stderr, "%s\n", message)
	os.Exit(1)
}
