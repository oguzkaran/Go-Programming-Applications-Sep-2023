/*
------------------------------------------------------------------------------------------------------------------------

	Gerçek sayı türleri için yani f format karakteri ile değerin noktadan sonraki kısmının kaç basamak olarak formatlanacağı
	belirlenebilir. Geri kalan değerler için bilimsel yuvarlama yapılır. Bunun için f ile % arasında nokta ile birlikte
	pozitif bir sayı yazılır. Bu sayı kaç basamağın formatlanacağını belirtir

------------------------------------------------------------------------------------------------------------------------
*/
package main

import "fmt"

func main() {
	var val float64

	fmt.Print("Input a value:")
	fmt.Scan(&val)

	fmt.Printf("val = %.10f\n", val)
}
