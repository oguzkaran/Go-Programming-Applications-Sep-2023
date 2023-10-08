/*
------------------------------------------------------------------------------------------------------------------------

	Sabitler (literals): Program içerisinde doğrudan yazılan değerlere denir. Go'da sabitler genel olarak şu şekilde
	kategorize edilmiştir
		- integer literals
		- floating point literals
		- complex literals
		- boolean literals
		- rune literals
		- string literals

	Burada integer, floating point complex ve rune türden sabitlere numeric contants da denilmektedir.

	Sabitlerden operatörlerden ve değişkenlerden oluşan ifadelere sabit ifadesi (constant expression) denir. Bu anlamda
	bir sabit de tek başına bir sabit ifadesidir. Go'da sabit olarak ele alınan özel değişkenler de bulunmaktadır. Aslında
	yukarıdaki kategoriler sabitlerin türlerine ilişkindir. Sabitlerin türlerine ilişkin kurallar şu şekildedir:

	- Sayı nokta içermiyorsa, herhangi bir son ek almamışsa ve tamsayı türüne ilişkindir.
		- Bu durumda sabit int türü sınırları içerisindeyse int türdendir. Sayı int sınırları dışındaysa error oluşur
		- Tamsayılar için int türünden size olarak küçük olan türlerden sabit yoktur
		- Sayı nokta içermiyorsa ve  sonunda L (küçük veya büyük) bitişik olarak eki varsa sabit long türdendir.
		- Sayı nokta içermiyorsa ve sonunda U (küçük veya büyük) eki varsa unsigned bir sabittir

	- Sayı nokta içeriyorsa gerçek sayı sabitidir (floating point literal)

	- Genel biçimi şu şekilde olan sabitler complex türdendir [(]<real value>+<imaginary value>i[)]
	Burada parantez konduğundan + işeareti ile boşluklar ayrılabilir

	- bool türden iki sabit vardır: true, false

	- rune sabitler aslında genel olarak karakter sabitlerine karşılık gelir ve tek tırnak içerisinde bir karakter ile
	yazılır

	- İki tırnak içerisinde yazılan, tırnaklarıyla birlikte yazılar string sabitleridir

------------------------------------------------------------------------------------------------------------------------
*/
package main

import "fmt"

func main() {
	fmt.Printf("%T\n", 'D')
}
