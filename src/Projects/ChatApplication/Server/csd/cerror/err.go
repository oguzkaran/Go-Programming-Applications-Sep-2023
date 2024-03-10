package cerror

import (
	"fmt"
	"os"
)

func CheckError(err error, message string) {
	CheckErrorCode(err, message, 1)
}

func CheckErrorCode(err error, message string, code int) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s:%s", message, err.Error())
		os.Exit(code)
	}
}
