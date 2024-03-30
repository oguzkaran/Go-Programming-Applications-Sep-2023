package main

import (
	"020-ConcatFilesApp/csd/err"
	"fmt"
	"io"
	"os"
)

func concatFiles(fd *os.File) {
	defer func() { _ = fd.Close() }()
	n := len(os.Args) - 1

	for i := 2; i < n; i++ {
		fs, e := os.Open(os.Args[i])

		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Can not open file:%s\n", os.Args[i])
			continue
		}
		_, e = io.Copy(fd, fs)

		if e != nil {
			err.ExitFailureError("Copy Problem", e)
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		err.ExitFailure("usage: csd_cf <file1> <file2> ... <fileN> <dest_path>")
	}

	fd, e := os.OpenFile(os.Args[len(os.Args)-1], os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if e != nil {
		err.ExitFailureError("OpenFile", e)
	}

	concatFiles(fd)
}
