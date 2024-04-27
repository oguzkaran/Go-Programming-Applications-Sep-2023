/*
------------------------------------------------------------------------------------------------------------------------

	Bir tamsayının belirli bir bitinin değerinin elde edilmesi:


	if (a & 1 << k) != 0

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"SampleGoLand/csd/console"
	"SampleGoLand/csd/util/bitwise"
	"fmt"
)

func main() {
	for {
		a := uint32(console.ReadInt("Input number:", ""))
		k := uint32(console.ReadInt("Input bir index:", ""))

		fmt.Println(bitwise.BitStrUInt32(a))
		fmt.Println(bitwise.BitStrUInt32(bitwise.ToggleBitUInt32(a, k)))
		if a == 0 {
			break
		}
	}
}
