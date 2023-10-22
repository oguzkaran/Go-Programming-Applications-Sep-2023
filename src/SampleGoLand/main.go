/*
------------------------------------------------------------------------------------------------------------------------

	Sınıf Çalışması: Parametresi ile aldığı int türden bir değerin asal olup olmadığını test eden isPrime isimli
	fonksiyonu yazınız ve aşağıdaki kod ile test ediniz

	Kural: Pozitif bir sayı karekökünden küçük olan asal sayıların hiç birisine bölünemiyorsa asaldır (Eratosten)

	Not: Aşağıdaki isPrime fonksiyonu bir önceki algoritmaya göre oldukça hızlıdır. Ancak en hızlısı değildir

------------------------------------------------------------------------------------------------------------------------
*/

package main

import "fmt"

func main() {
	runIsPrimeTest()
}

func runIsPrimeTest() {
	fmt.Println(isPrime(6750161072220585911))

	fmt.Println("Tekrar yapıyor musunuz?")
}

func isPrime(val int) bool {
	if val < 1 {
		return false
	}

	if val%2 == 0 {
		return val == 2
	}

	if val%3 == 0 {
		return val == 3
	}

	if val%5 == 0 {
		return val == 5
	}

	if val%7 == 0 {
		return val == 7
	}

	for i := 11; i*i <= val; i += 2 {
		if val%i == 0 {
			return false
		}
	}

	return true
}
