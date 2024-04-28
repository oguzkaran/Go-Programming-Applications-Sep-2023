/*
------------------------------------------------------------------------------------------------------------------------

	Soru: Parametresi ile aldığı bir tamsayının en yüksek anlamlı sıfır olan bitinin indeks numarasına geri dönen
	HighestClearBitIndexUInt32 fonksiyonu ile en düşük anlamlı sıfır olan bitinin indeks numarasına geri dönen
	LowestClearBitIndexUInt32 fonksiyonunu bitwise paketine ekleyiniz

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

		fmt.Println(bitwise.BitStrUInt32(a))
		if (a & 1) != 1 {
			fmt.Printf("%d is even\n", a)
		} else {
			fmt.Printf("%d is odd\n", a)
		}
		if a == 0 {
			break
		}
	}
}
