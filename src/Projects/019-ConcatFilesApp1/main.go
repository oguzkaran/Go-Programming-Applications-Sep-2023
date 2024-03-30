/*
------------------------------------------------------------------------------------------------------------------------
   Sınıf Çalışması: Komut satırından aşağıdaki gibi çalışan programı yazınız
        ./csd_cf <block size> <file1> <file2> ... <fileN> <dest_path>

    - Program file1, file2, file3, ..., fileN ile belirtilen dosyaları birleştirerek dest ile belirtilen dosyaya
    ekleyecektir

    - Dosyalar komut satırından alınan block_size uzunluğu kadar okunacaktır

    - Program olmayan kaynak dosyalar için uygun mesajı verecek, akışına devam edecektir

    - Hedef dosya varsa üzerine yazılacaktır (overwrite)
------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"019-ConcatFilesApp/csd/err"
	"fmt"
	"os"
	"strconv"
)

func concatFiles(fd *os.File, blockSize int) {
	defer func() { _ = fd.Close() }()
	n := len(os.Args) - 1

	for i := 2; i < n; i++ {
		fs, e := os.Open(os.Args[i])
		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Can not find file:%s\n", os.Args[i])
			continue
		}

		defer func() { _ = fs.Close() }()

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
