/*
------------------------------------------------------------------------------------------------------------------------

	Sınıf Çalışması: Aşağıdaki uygulamayı yazınız:
		+ Uygulama aşağıdaki menü ile başlayacaktır:
				1. User Operations
				2. Product Operations
				3. Payment Operations
				4. Exit
				Option:

		+ User Operations menusü seçildiğinde aşağıdaki menu açılacaktır
				User Operations
				1. Register
				2. Login
				3. Find
				5. Quit
				Option:
			+ Register seçildiğinde yeni kullanıcıya ilişkin bilgiler stdin'den istenecek ve kayıt işlemi yapılacaktır.
			Kullanıcının var olması durumu mesaj ile belirtilecek ve menüye geri dönülecektir.

			+ Find seçildiğinde stdin'den username istenecek ve ilgili user'ın bilgileri stdout'a gönderilip menüye geri
			dönülecektir.

			+ Quit seçildiğinde ana menüye geri dönülecektir

		- Product Operations menusü seçildiğinde aşağıdaki menu açılacaktır
				User Operations
				1. Add
				2. Delete
				3. Find By Product Code
				4. Find In Stock
				5. Quit
				Option:
			- Add seçildiğinde stdin'den alınan bilgilere göre ilgili ürün eklenecektir

			- Delete seçildiğinde stdin'den alınan ürün koduna ilişkin ürün varsa silinecektir. Silme işlemi onay istenerek
			yapılacaktır. Eğer belirlenen ürüne kodun ilişkin ürün yoksa ilgili mesaj kullanıcıya verilecektir.

			- Find By Product Code seçildiğinde ilgili ürüne ilişkin detaylar listelenecektir

			- Find In Stock seçildiğinde stokta stdin'den alınan ürün koduna ilişkik ürün varsa listelenecek, veritabanında
			varsa fakat stokta yoksa veya ürün hiç veritabanında yoksa uygun mesajlar verilecektir

			- Quit seçildiğinde ana menüye geri dönülecektir
		+ Exit seçildiğinde uygulama sonlandırılacaktır
	Açıklamalar:
		1. Geliştirme sırasında uyugun katmanlara uygun eklentiler yapılacaktır.
		2. Uygulamada PaymentAppDataService yapısının fonksiyonları ile (yani servis katmanına erişilerek) işlem
		yapılacaktır
		3. Tüm repository'ler CrudRepository arayüzünü destekleyecektir. İhtiyaç olmayan metotlar panic fonksiyonu
		çağrılarak bırakılacaktır (not implemented yet!...)

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"PaymentServiceApp/app"
	_ "github.com/lib/pq"
)

func main() {
	app.Run()
}
