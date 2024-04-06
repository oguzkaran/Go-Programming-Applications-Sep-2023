package main

import (
	"020-ConcatFilesApp/csd/err"
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func zipFiles(fd *os.File) {
	defer func() { _ = fd.Close() }()
	zw := zip.NewWriter(fd)

	defer func() { _ = zw.Close() }()
	n := len(os.Args) - 1

	for i := 2; i < n; i++ {
		fs, e := os.Open(os.Args[i])

		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Can not open file:%s\n", os.Args[i])
			continue
		}

		defer func() { _ = fs.Close() }()
		w, e := zw.Create(os.Args[i])

		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Can not create a zip file:%s\n", os.Args[i])
			continue
		}

		if _, e = io.Copy(w, fs); e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Can not zip file:%s\n", os.Args[i])
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		err.ExitFailure("usage: csd_zf <file1> <file2> ... <fileN> <dest_path>")
	}

	fd, e := os.Create(fmt.Sprintf("%s.zip", os.Args[len(os.Args)-1]))

	if e != nil {
		err.ExitFailureError("OpenFile", e)
	}

	zipFiles(fd)
}
