/*
------------------------------------------------------------------------------------------------------------------------

		Aralarında öncelik sonralık ilişkisi olan, bellekte elemanların peşpeşe olarak yaratılması gerekmediği veri
		yapılarına bağlı liste (linked list) denir. Bağlı listeler çift (doubly) ve tek (singly) olmak üzere iki genel
		gruba ayrılır. Bağlı listenin bir elemanı bir önceki elemandan sonra gelmek zorunda olmadığına göre önceki elemanın
		sonraki elemanın adresini tutması gerekir. Burada aslında eleman doğrudan tuttuğu değer düşünülmemelidir. Bağlı
		listede tutulan değer ile birlikte adreslerin de (şüphesiz başka bilgiler de olabilir) bulunduğu yapıya
		"node" denir. Go'da collection içerisinde bulunun list paketi içerisinde doubly linked veri yapısı gerçekleştirilmiştir.
		Bu veri yapısında node terimi yerine soyutlama amaçlı Element terimi kullanılmıştır.

		Bu durumda bir çift bağlı liste nasıl organize edilebilir? Bağlı litenin kendisi bir yapı ile temsil edilebilir.
		Bu yapı içerisinde ilk eleman (head) ve son elemanın (tail) adresleri tutulabilir. Ayrıca listenin eleman sayısı da
		tutulabilir. Örneğin CSDIntList gibi bir çift bağlı listenin yapıları şu şekilde olabilir:
			type IntNode struct {
				Value int
				next  *IntNode
				prev  *IntNode
			}

			type CSDIntList struct {
				head *IntNode
				tail *IntNode
				size int
		   }

	Anahtar Notlar: Burada bir çift bağlı liste (doubly linked list) generic olmayacak  şekilde geliştirilecektir. Generic
	konusuna gelindiğinde veri yapısı generic yapılacaktır

------------------------------------------------------------------------------------------------------------------------
*/
package main

import (
	"container/list"
	"fmt"
	"math/rand"
)

import (
	"SampleGoLand/csd/container"
)

func fillRandom(lst *container.CSDIntList, count int) {
	for i := 0; i < count; i++ {
		lst.PushBack(rand.Intn(100))
	}
}

func main() {
	list.New()
	lst := container.New()

	fillRandom(lst, 10)
	fmt.Println(lst.IsEmpty())
	fmt.Printf("Size:%d\n", lst.Size())

	lst.Clear()

	fmt.Printf("Size:%d\n", lst.Size())

	fmt.Println(lst.IsEmpty())
}
