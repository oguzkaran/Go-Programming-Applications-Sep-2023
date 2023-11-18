/*
------------------------------------------------------------------------------------------------------------------------

	Aşağıdaki init fonksiyonları bildirim sırasına göre main fonksiyonundan önce çağrılır

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"SampleGoLand/numeric"
	"fmt"
	"math/rand"
)

func main() {
	for {
		a := rand.Intn(20) - 10
		b := rand.Intn(20) - 10
		p, status := numeric.New(a, b)

		fmt.Println("------------------------------------")
		fmt.Printf("a = %d, b = %d\n", a, b)
		if status == 1 {
			fmt.Printf("%d / %d\n", p.GetNumerator(), p.GetDenominator())
		} else if status == 0 {
			fmt.Println("Undefined")
		} else {
			fmt.Println("Indeterminate")
		}
		fmt.Println("------------------------------------")
		if a == 0 && b == 0 {
			break
		}
	}
}
