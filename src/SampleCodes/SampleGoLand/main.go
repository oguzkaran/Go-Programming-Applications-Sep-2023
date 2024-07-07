/*
------------------------------------------------------------------------------------------------------------------------
    Aşağıdaki demo örneği inceleyiniz
------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
)

func main() {
	count := flag.Int("c", 10, "use as value")
	origin := flag.Int("o", 0, "use as a value")
	bound := flag.Int("b", 100, "use as value")

	flag.Parse()

	for i := 0; i < *count; i++ {
		fmt.Printf("%d ", rand.IntN(*bound-*origin)+*origin)
	}

	fmt.Println()
}
