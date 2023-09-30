/*
----------------------------------------------------------------------------------------------------------------------

	Go'da tüm fonksiyonların dışında paket içerisinde bildirilen değişkenlere global değişkenler (global variables)	denir.
	Global değişkenler aynı dosyadaki tüm fonksiyonlar içerisinde görülebilirdir. Bu duruma ilişkin bazı ayrıntılar
	ileride ele alınacaktır

----------------------------------------------------------------------------------------------------------------------
*/
package main

import "fmt"

var g_a int32 = 34

func main() {
	fmt.Println(g_a)
	foo()
	fmt.Println(g_a)
	bar()
	fmt.Println(g_a)
}

func foo() {
	g_a = 10
}

func bar() {
	g_a = 20
}
