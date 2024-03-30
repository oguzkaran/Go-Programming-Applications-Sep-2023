package main

import (
	"019-ConcatFilesApp/csd/err"
	"fmt"
	"io"
	"os"
	"strconv"
)

func copyFile(fd *os.File, fs *os.File, buf []byte) error {
	for {
		defer func() { _ = fs.Close() }()
		result, e := fs.Read(buf)
		if e != nil && e != io.EOF {
			return e
		}
		if e == io.EOF {
			break
		}

		writeBuf := make([]byte, result)
		copy(writeBuf, buf)
		_, e = fd.Write(writeBuf)
		if e != nil {
			return e
		}
	}

	return nil
}

func concatFiles(fd *os.File, blockSize int) {
	defer func() { _ = fd.Close() }()
	n := len(os.Args) - 1

	buf := make([]byte, blockSize)

	for i := 2; i < n; i++ {
		fs, e := os.Open(os.Args[i])

		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Can not open file:%s\n", os.Args[i])
			_ = os.Remove(fd.Name())
			break
		}
		e = copyFile(fd, fs, buf)

		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Read problem in file:%s\n", os.Args[i])
		}
	}
}

func main() {
	if len(os.Args) < 4 {
		err.ExitFailure("usage: csd_cf <block size> <file1> <file2> ... <fileN> <dest_path>")
	}

	blockSize, e := strconv.Atoi(os.Args[1])

	if e != nil {
		err.ExitFailure("invalid block size value!...")
	}

	fd, e := os.OpenFile(os.Args[len(os.Args)-1], os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)

	if e != nil {
		err.ExitFailureError("OpenFile", e)
	}

	concatFiles(fd, blockSize)
}
