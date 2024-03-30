package file

import (
	"io"
	"os"
)

func findByteTotal(f *os.File) (int, error) {
	total := 0
	buf := make([]byte, 1)

	for {
		n, e := f.Read(buf)

		if e != nil && e != io.EOF || n != 1 {
			return -1, e
		}

		if e == io.EOF {
			return total, nil
		}

		total += int(buf[0])
	}
}

func Checksum(path string) (bool, error) {
	f, e := os.Open(path)
	if e != nil {
		return false, e
	}
	total, e := findByteTotal(f)

	if e != nil {
		return false, nil
	}

	return total == 0, nil
}
