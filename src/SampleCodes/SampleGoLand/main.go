/*
------------------------------------------------------------------------------------------------------------------------

	Standart container/list paketi içerisindeki List yapısı

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"SampleGoLand/csd/container/clist"
	"fmt"
	"math/rand"
)

func fillRandom(lst *clist.List, size int) {
	for i := 0; i < size; i++ {
		lst.CList.PushBack(rand.Intn(100))
	}
}

func printList(lst *clist.List) {
	fmt.Printf("Size:%d\n", lst.CList.Len())
	for e := lst.CList.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}

	fmt.Println()
}

func main() {
	lst := clist.New()

	fillRandom(lst, 10)
	printList(lst)
	lst.CList.Init()
	printList(lst)
}
