/*
------------------------------------------------------------------------------------------------------------------------
	Aşağıdaki örnekte fonksiyon sürekli kendisini çağrıdığından. Diğer bir deyişle kendisini çağırmayı belli bir noktada
	sonlandırmadığından stack taşması oluşur. Fonksiyonun herhangi bir parametre değişkeni ve yerel değişkeni olmadığına
	dikkar ediniz
------------------------------------------------------------------------------------------------------------------------
*/

package main

func foo() {
	foo()
}

func main() {
	foo()
}
