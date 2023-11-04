/*
------------------------------------------------------------------------------------------------------------------------

	Yapılar (Structures): Go'da programcı da built-in türler dışında türler tanımlayabilir (user defined type). Bu işlem
	tipik olarak yapılar (structures) kullanılarak gerçekleştirilir. Bir grup bilginin bir araya getirilmesiyle oluşan
	türe denir. Bir yapı aslında bir bileşik türdür (compund type). Yapılar özellikle tek başına anlam ifade etmeyen
	ancak biraraya geldiklerinde bir tür belirten nesneler için kullanılır. Yapılar aynı zamanda birer data structure'dır.
	Yapı bildiriminin genel biçimi şu şekildedir:
		type <isim> struct {
			<isim> <tür>
			<isim> <tür>
			...
		}
	Yapı içerisinde bildirilen değişkenlere yapı elemanı (structure member/field/member variable) denir. Yapı elemanlarına
	nokta operatörü ile erişilir. Nokta operatörü iki operandlı araek durumunda bir operatördür. Yapılar için birinci
	operandı yapı nesnesi veya yapı türünden gösterici olabilir. İkinci operandı ise yapının elemanı veya yapıya
	eklenmiş (extension) bir fonksiyon olabilir. Yapılara eklenen fonksiyonlar ayrıca ele alınacaktır. Bir yapı nesnesi
	yaratıldığında elemanları default değerlerini alırlar. Bir yapı Print ve Prinln gibi fonksiyonlara argüman olarak
	geçilirse bu durumda yapı eleman değerleri bildirim sırasıyla { ile } aralarında space ile ayrılarak listelenir
------------------------------------------------------------------------------------------------------------------------
*/

package main

import "fmt"

type Point struct {
	x, y int
}

type PointF struct {
	x, y float32
}

func main() {
	var p Point
	var pf PointF

	fmt.Println(p)
	fmt.Println(pf)

	p.x = 100
	p.y = 2000
	pf.x = 3.456
	pf.y = -5678.89

	fmt.Println(p)
	fmt.Println(pf)
}
