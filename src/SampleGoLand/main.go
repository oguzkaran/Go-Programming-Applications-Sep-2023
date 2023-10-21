/*
------------------------------------------------------------------------------------------------------------------------

	Sınıf Çalışması: Parametresi ile aldığı int türden bir değerden büyük ilk Fibonacci sayısını döndüren nextFibonacci
	fonksiyonunu yazınız ve test ediniz

------------------------------------------------------------------------------------------------------------------------
*/

package main

import "fmt"

func main() {
	runFactorialTest()
}

func runFactorialTest() {
	for n := -1; n <= 10; n++ {
		fmt.Printf("%d! = %d\n", n, factorial(n))
	}
}

func factorial(n int) int {
	result := 1

	for ; n > 1; n-- {
		result *= n
	}

	return result
}
