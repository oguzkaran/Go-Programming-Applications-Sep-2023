package main

import (
	"017-RandomNumberGeneratorPosixApp/csd/err"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		err.ExitFailure("usage:csd_rg <count>")
	}

	count, e := strconv.Atoi(os.Args[1])

	if e != nil || count <= 0 {
		err.ExitFailure("invalid count value")
	}

	f, e := os.Open("/dev/random")

	if e != nil {
		err.ExitFailureError("Open:", e)
	}

	defer func() { _ = f.Close() }()

	var value int64

	for i := 0; i < count; i++ {
		e = binary.Read(f, binary.LittleEndian, &value)

		if e != nil {
			err.ExitFailureError("binary.Read", e)
		}
		fmt.Println(value)
	}
}
