/*
------------------------------------------------------------------------------------------------------------------------

	10 -> 000A
			L.E.	B.E.
	1FC0 -> 0A		00
	1FC1 -> 00		0A

------------------------------------------------------------------------------------------------------------------------
*/
package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var x uint64 = 0x000000000000000A

	buf := make([]byte, 8)
	binary.NativeEndian.PutUint64(buf, x)
	for _, b := range buf {
		fmt.Printf("%X ", b)
	}

	fmt.Println()
	val := binary.NativeEndian.Uint64(buf)

	fmt.Printf("val = %X, val = %d\n", val, val)
	fmt.Printf("x = %X, x = %d\n", x, x)

	fmt.Println()
}
