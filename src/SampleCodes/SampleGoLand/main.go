/*
------------------------------------------------------------------------------------------------------------------------

	Go'da iota isimli bir tamsayı sabit vardır. Bu sabit ile peşpeşe olacak şekilde bazı değerler belirlenebilmektedir. Go'da
	bazı dillerde bulunan enum türü bulunmamaktadır. Ancak iota kullanılarak çeşitli sabitler belirlenebilmektedir. Bu da
	enum ihtiyacını karşılamaktadır. Go'da const anahtar sözcüğü ile sabit tanımlaması yapılabilmektedir. Aşağıdaki demo
	örneği inceleyiniz

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"SampleGoLand/csd/game/card"
	"fmt"
)

func main() {
	deck := card.NewDeck()

	for _, c := range deck {
		fmt.Println(c)
	}
}
