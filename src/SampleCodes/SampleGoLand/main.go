/*
------------------------------------------------------------------------------------------------------------------------

	github.com/golang-collections/collections/queue paketindeki Queue veri yapısı ile FIFO kuyruk temsil edilmektedir

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
)

func main() {
	q := queue.New()

	for i := 0; i < 10; i++ {
		q.Enqueue(i * 10)
	}

	for {
		value := q.Dequeue()
		if q.Len() == 0 {
			break
		}
		fmt.Printf("%v ", value)
	}
}
