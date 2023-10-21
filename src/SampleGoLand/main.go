/*
------------------------------------------------------------------------------------------------------------------------

	Sınıf Çalışması: Klavyeden sıfır girilene kadar alınan int türden sayılardan pozitif olanlarının ve negatif olanlarının
	ayrı ayrı kaç tane olduklarını ve ayrı ayrı toplamlarını bulan programı yazınız. Hiç pozitif sayı ya da negatif
	girilmemişse duruma uygun mesajı veriniz

------------------------------------------------------------------------------------------------------------------------
*/

package main

import "fmt"

func main() {
	runApp()
}

func runApp() {
	printResults(findPosNegResults())
}

func printResults(posSum int, posCount int, negSum int, negCount int) {
	printPosResult(posSum, posCount)
	printNegResult(negSum, negCount)
}

func printPosResult(posSum int, posCount int) {
	if posCount != 0 {
		fmt.Printf("You entered %d positive number with total: %d\n", posCount, posSum)
	} else {
		fmt.Println("You did not enter a positive number!...")
	}
}

func printNegResult(negSum int, negCount int) {
	if negCount != 0 {
		fmt.Printf("You entered %d negative number with total: %d\n", negCount, negSum)
	} else {
		fmt.Println("You did not enter a negative number!...")
	}
}

func findPosNegResults() (int, int, int, int) {
	posSum, posCount, negSum, negCount := 0, 0, 0, 0

	var val int

	for {
		fmt.Print("Input a number:")
		fmt.Scan(&val)

		if val == 0 {
			break
		}

		if val > 0 {
			posSum += val
			posCount++
		} else {
			negSum += val
			negCount++
		}
	}

	return posSum, posCount, negSum, negCount
}
