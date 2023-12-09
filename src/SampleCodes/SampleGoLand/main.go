/*
------------------------------------------------------------------------------------------------------------------------

	Interface'ler Go'da özel bir tür olarak içerisinde metot ya da metotların prototiplerini barındırır. Tipik olarak
	bir yapı ilgili metodu yazdığında o arayüzü mantıksal olarak desteklemiş (implementation). interface aslında bir
	tür bildirimi olduğundano interface türünden bir nesne tanımlanabilir. Bu anlamda interface aslında soyutlamayı (abstraction)
	sağlar. Buradaki soyutlama bir takım türlerin ortak davranışlarını içermesi anlamındadır. Bu anlamda Go'da runtime
	polymorphism (RTP) denilen bir kavram belirli ölçüde interface'ler ile sağlanabilir.

	Anahtar Notlar: Programlamaya biyolojiden aktarılmıiş önemli bir kavram da polymorphism'dir. Polymorphism genel olarak
	iki gruba ayrılır: runtime polymorphism (RTP), compile time polymorphism (CTP). Aslında özellike RTP nesne yönelimli
	programlamada teknik (NYPT) olarak da desteklenen bir kavramdır. Ancak mantıksal kavramı NYPT'ni desteklemeyen diller
	için de söz konusudur. RTP'nin Biyoloji'de şu şekilde tanımlanabilir: Farklı doku ya da organların evrim süreci
	içerisinde temel işlevleri aynı kalmak koşuluyla o işlevi yerine getirme biçiminin değişebilmesidir. Programlamada
	RTP için pek çok tanım ve betimleme söz konusu olabilse de, bunların hepsi şu 3 tanım ile özetlenebilir.
	1. Biyolojik Tanım: Bir fonksiyonun bir türde yeniden gerçekleştirilmesidir
	2. Yazılım Mühendisliği Tanımı: Türden bağımsız kod yazmaktır
	3. Aşağı seviyeli tanım: Önceden yazılmış kodların sonradan yazılmış kodları çağırabilmesidir
	Bu tanımlar NYPT açısından bazı başka kavramları da içerir. Burada Go bakış açısıyla tanımlanmıştır

	Bir interface bildiriminin genel biçimi şu şekildedir:
		type <interface ismi> interface {
			<metot prototipleri/imzaları>
		}

	Burada metot bildirimlerinde func anahtar sözcüğü kullanılmaz.


	Aslında bir interface bir anlaşmayı (contract) temsil eder ve bu interface'i destekleyen türler de bu anlaşmaya
	uyan türlerdir. Bir interface'i destekleyen bir yapı türünden nesne desteklediği interface türünden bir nesneye
	doğrudan atanabilir (implicit conversion).
------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	capp "SampleGoLand/csd/parser/source/app"
	dapp "SampleGoLand/deniz/app"
)

func main() {
	capp.RunWhitespaceCountParserApp()
	dapp.RunDigitCountParserApp()
}
