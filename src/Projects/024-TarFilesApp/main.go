package main

import (
	"020-ConcatFilesApp/csd/err"
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func tarFiles(fd *os.File) {
	defer func() { _ = fd.Close() }()
	gw := gzip.NewWriter(fd)
	defer func() { _ = gw.Close() }()

	tw := tar.NewWriter(gw)
	defer func() { _ = tw.Close() }()
	n := len(os.Args) - 1

	for i := 2; i < n; i++ {
		fs, e := os.Open(os.Args[i])

		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Can not open file:%s\n", os.Args[i])
			continue
		}

		defer func() { _ = fs.Close() }()

		stat, e := fs.Stat()

		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Can not get stat file:%s\n", os.Args[i])
			continue
		}

		if stat.IsDir() {
			_, _ = fmt.Fprintf(os.Stderr, "%s is a directory\n", os.Args[i])
			continue
		}

		h := &tar.Header{Name: os.Args[i], Size: stat.Size(), Mode: int64(stat.Mode()), ModTime: stat.ModTime()}

		if e = tw.WriteHeader(h); e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Can not write header\n")
			continue
		}
		
		if _, e = io.Copy(tw, fs); e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Can not gzip file:%s\n", os.Args[i])
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		err.ExitFailure("usage: csd_zf <file1> <file2> ... <fileN> <dest_path>")
	}

	fd, e := os.Create(fmt.Sprintf("%s.tar.gz", os.Args[len(os.Args)-1]))

	if e != nil {
		err.ExitFailureError("OpenFile", e)
	}

	tarFiles(fd)
}
