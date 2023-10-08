/*
------------------------------------------------------------------------------------------------------------------------

	Tamsayılar çeşitli sayı sistemlerinde (base/radix) formatlanabilmektedir. Örneğin hexadecimal için x veya X, octal
	için o veya O kullanılabilir. Büyük O ile sayının başında sıfır ve o karakteri de eklenir. b format karakteri ile
	değer binary olarak formatlanabilir

------------------------------------------------------------------------------------------------------------------------
*/
package main

import "fmt"

func main() {
	var val int

	fmt.Print("Input a value:")
	fmt.Scan(&val)

	fmt.Printf("val = %d\n", val)
	fmt.Printf("val = %x\n", val)
	fmt.Printf("val = %X\n", val)
	fmt.Printf("val = %b\n", val)
	fmt.Printf("val = %o\n", val)
	fmt.Printf("val = %O\n", val)
}
