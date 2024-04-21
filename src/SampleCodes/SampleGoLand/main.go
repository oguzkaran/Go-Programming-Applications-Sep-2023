/*
------------------------------------------------------------------------------------------------------------------------

	<< operatörü birinci operandına ilişkin tamsayının bitlerini, ikinci operandına ilişkin sayı kadar pozisyon sola
	kaydırılmasından oluşan değeri üretir. Sınır dışına çıkan bitler için sağdan sıfıf biti ile besleme yapılır

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"SampleGoLand/csd/util/bitwise"
	"fmt"
)

func main() {
	var a uint32 = 0x0000001

	for a != 0 {
		fmt.Println(bitwise.BitStrUInt32(a))
		a <<= 1
	}
}
