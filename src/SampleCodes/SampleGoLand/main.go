/*
------------------------------------------------------------------------------------------------------------------------

	Soru: Parametresi ile aldığı bir tamsayının en yüksek anlamlı sıfır olan bitinin indeks numarasına geri dönen
	HighestClearBitIndexUInt32 fonksiyonu ile en düşük anlamlı sıfır olan bitinin indeks numarasına geri dönen
	LowestClearBitIndexUInt32 fonksiyonunu bitwise paketine ekleyiniz. Fonksiyonlar sayının ilgili değere ilişkin bir
	biti yoksa -1 değerine geri dönecektir

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
		fmt.Printf("Lowest set bit Index:%d\n", bitwise.LowestSetBitIndexUInt32(a))
		fmt.Printf("Lowest clear bit Index:%d\n", bitwise.LowestClearBitIndexUInt32(a))
		if a == 0 {
			break
		}
	}
}
