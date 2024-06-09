package file

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadTestFile(path string) []string {
	f, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defer func() { _ = f.Close() }()

	reader := bufio.NewReader(f)

	var lines []string

	for {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		if err == io.EOF && len(line) == 0 {
			break
		}
		line = strings.TrimRight(line, "\r\n")
		lines = append(lines, line)
	}

	return lines
}

func GetOutputOfPrintNumber(f func(int), a int) string {
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f(a)

	_ = w.Close()
	os.Stdout = old
	_, _ = buf.ReadFrom(r)
	return buf.String()
}
