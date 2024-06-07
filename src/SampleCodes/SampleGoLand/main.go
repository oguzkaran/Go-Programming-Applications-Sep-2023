/*
------------------------------------------------------------------------------------------------------------------------
    Homework-001-1.sorunun bir çözümü
    (Not:Çözüm çalışma sorusunun verildiği tarihte işlenmiş konulara göre yazılmıştır)
------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"SampleGoLand/csd/console"
	"SampleGoLand/csd/recursion/recursionutil"
	"fmt"
)

func main() {
	for {
		n := console.ReadInt("Input n:", "")

		if n <= 0 {
			break
		}

		fmt.Printf("Fibonacci(%d) = %d\n", n, recursionutil.Fibonacci(n))
	}
}
