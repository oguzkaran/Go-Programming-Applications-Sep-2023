/*
------------------------------------------------------------------------------------------------------------------------
    Homework-002-1.sorunun bir çözümü
    (Not: Çözüm çalışma sorusunun verildiği tarihte işlenmiş konulara göre yazılmıştır)
------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func printAbove(n int) {
	for i := 0; i < n; i++ {
		for k := 0; k < n-i; k++ {
			fmt.Print(" ")
		}

		for k := 0; k < 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

func printBelow(n int) {
	for i := 0; i < n; i++ {
		for k := 0; k < i; k++ {
			fmt.Print(" ")
		}

		for k := 0; k < 2*(n-i)-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

func printDiamond(n int) {
	printAbove(n)
	printBelow(n)
}

func run() {
	var n int
	fmt.Print("Input n:")
	_, _ = fmt.Scanf("%d", &n)
	printDiamond(n)
}

func main() {
	run()
}
