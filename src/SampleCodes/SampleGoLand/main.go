/*
------------------------------------------------------------------------------------------------------------------------

	Aşağıdaki demo örnekte yaratılan thread'leri stack alanları farklı olduğundan i döngü değişkeni de her thread için
	ayrı olarak yaratılır

------------------------------------------------------------------------------------------------------------------------
*/
package main

import (
	"fmt"
	"time"
)

type Sample struct {
	val int
}

func goroutineCallback(count int, ps *Sample) {
	for i := 0; i < count; i++ {
		ps.val++
	}
	fmt.Println("goroutineCallback ends!...")
}

func main() {
	s := Sample{}
	go goroutineCallback(10_000_000, &s)
	go goroutineCallback(10_000_000, &s)
	time.Sleep(3 * time.Second)
	fmt.Printf("value = %d\n", s.val)
}
