/*
------------------------------------------------------------------------------------------------------------------------

	Aşağıdaki demo örnekte text bir dosyadan karakter karakter okuma yapılmış ve stdout'a bastırılmıştır

------------------------------------------------------------------------------------------------------------------------
*/
package main

import (
	"016-PrintFileApp/csd/err"
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		err.ExitFailure("usage: csd_cat <file path>")
	}

	f, e := os.Open(os.Args[1])

	if e != nil {
		err.ExitFailureError("Open:", e)
	}

	defer func() { _ = f.Close() }()

	reader := bufio.NewReader(f)

	for {
		line, e := reader.ReadString('\n')

		if e == io.EOF {
			break
		}

		if e != nil {
			err.ExitFailureError("ReadString", e)
		}
		for _, c := range line {
			fmt.Print(string(c))
		}
	}
}
