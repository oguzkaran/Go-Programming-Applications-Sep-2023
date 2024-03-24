/*
------------------------------------------------------------------------------------------------------------------------

	Aşağıdaki örnekte blok blok dosya kopyalaması yapılmıştır. Burada blok uzunluğu kullanıcı tarafından belirlenmiştir

------------------------------------------------------------------------------------------------------------------------
*/
package main

import (
	"014-CopyFilesViaBlockApp/csd/err"
	"io"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		err.ExitFailure("usage:csd_copy <src path> <dest path> <block size>")
	}

	size, e := strconv.Atoi(os.Args[3])

	if e != nil {
		err.ExitFailure("invalid block size!...")
	}

	fs, e := os.OpenFile(os.Args[1], os.O_RDONLY, 0)

	if e != nil {
		err.ExitFailureError("OpenFile", e)
	}

	fd, e := os.OpenFile(os.Args[2], os.O_CREATE|os.O_WRONLY, 0777)

	if e != nil {
		err.ExitFailureError("OpenFile", e)
	}

	data := make([]byte, size)

	for {
		n, e := fs.Read(data)

		if e != nil && e != io.EOF {
			err.ExitFailureError("Read", e)
		}

		if n == 0 {
			break
		}

		writeData := make([]byte, n)
		copy(writeData, data)

		_, e = fd.Write(writeData)

		if e != nil {
			err.ExitFailureError("Write", e)
		}
	}
}
