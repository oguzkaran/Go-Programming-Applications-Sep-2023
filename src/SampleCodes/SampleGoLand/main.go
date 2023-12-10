/*
------------------------------------------------------------------------------------------------------------------------

	Programcı kısıt belirlemek için yine bir interface yazar. Bu interface içerisinde kısıt için belirleyeceği türleri
	| atomu ile belirtir

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"SampleGoLand/csd/util/slice"
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{1, 2, 3, 4, 5, 6, 7}

	if slice.Equals(a, b) {
		fmt.Println("Same arrays")
	} else {
		fmt.Println("Different arrays")
	}
}
