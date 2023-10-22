/*
------------------------------------------------------------------------------------------------------------------------

	Sınıf Çalışması: Parametresi ile aldığı int türden bir değerden büyük ilk Fibonacci sayısını döndüren nextFibonacci
	fonksiyonunu yazınız ve test ediniz

------------------------------------------------------------------------------------------------------------------------
*/

package main

import "fmt"

func main() {
	runFibonacciNumberTest()
}

func runFibonacciNumberTest() {
	for {
		var n int

		fmt.Print("n değerini giriniz:")
		fmt.Scan(&n)

		if n <= 0 {
			break
		}

		fmt.Printf("%d. Fibonacci sayısı:%d\n", n, fibonacciNumber(n))
	}

	fmt.Println("Tekrar yapıyor musunuz?")
}

func fibonacciNumber(n int) int {
	if n <= 2 {
		return n - 1
	}

	prev1, prev2 := 1, 0
	value := prev1 + prev2

	for i := 3; i < n; i++ {
		prev2 = prev1
		prev1 = value
		value = prev1 + prev2
	}

	return value
}
