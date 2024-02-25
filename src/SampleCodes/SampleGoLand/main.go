/*
------------------------------------------------------------------------------------------------------------------------

	Yukarıdaki problem aşağıdaki gibi channel kullanarak çözülebilir

------------------------------------------------------------------------------------------------------------------------
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var value = make(chan int)

func producerGoroutineCallback(wg *sync.WaitGroup) {
	defer wg.Done()

	val := 0

	for {
		value <- val
		val++
		if val >= 99 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	close(value)
}

func consumerGoroutineCallback(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		val := <-value
		fmt.Printf("%d ", val)

		if val >= 98 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go producerGoroutineCallback(&wg)
	go consumerGoroutineCallback(&wg)

	wg.Wait()
}
