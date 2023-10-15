/*
------------------------------------------------------------------------------------------------------------------------

	Sınıf Çalışması: Katsayıları klavyeden girilen ikinci dereceden bir denklemin köklerini bulan programı yazınız

------------------------------------------------------------------------------------------------------------------------
*/

package main

import "fmt"

func main() {
	var a int

	fmt.Print("Bir sayı giriniz:")
	fmt.Scan(&a)

	if isOdd(a) {
		fmt.Println("Tek sayı!...")
	} else {
		fmt.Println("Çift sayı!...")
	}
}

func isEven(val int) bool {
	return val%2 == 0
}

func isOdd(val int) bool {
	return !isEven(val)
}
