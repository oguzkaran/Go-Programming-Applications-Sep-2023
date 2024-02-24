/*
------------------------------------------------------------------------------------------------------------------------

	Aşağıdaki demo örnekteki calculateCounter fonksiyonunda return deyiminden önce UnLock yapılmazsa deadlock oluşur. Go
	runtime goroutine'ler için deadlock durumlarını yakalar ve panic fonksiyonu çağrılır

------------------------------------------------------------------------------------------------------------------------
*/
package main

import (
	"SampleGoLand/csd/console"
	"fmt"
	"math/rand"
	"sync"
)

type Counter struct {
	value int32
	total int64
}

func calculateCounter(counter *Counter, mutex *sync.Mutex) {
	mutex.Lock()
	var value = int32(rand.Intn(20) - 10)

	mutex.Lock()
	counter.value += value
	counter.total += int64(counter.value)
	mutex.Unlock()
	mutex.Unlock()
}

func goroutineCallback(n int, counter *Counter, wg *sync.WaitGroup, mutex *sync.Mutex) {
	for i := 0; i < n; i++ {
		calculateCounter(counter, mutex)
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	count := console.ReadInt("Input count:", "")
	n := console.ReadInt("Input number of goroutines:", "")

	counter := Counter{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go goroutineCallback(count, &counter, &wg, &mutex)
	}

	wg.Wait()

	fmt.Printf("Value:%d\n", counter.value)
	fmt.Printf("Total:%d\n", counter.total)
}
