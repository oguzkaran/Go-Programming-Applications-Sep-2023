/*
------------------------------------------------------------------------------------------------------------------------

	Sınıf Çalışması: Aşağıdaki uygulamayı yazınız:
		- Uygulama aşağıdaki menü ile başlayacaktır:
				1. User Operations
				2. Product Operations
				3. Payment Operations
				4. Exit
				Option:

		Users menusü seçildiğinde aşağıdaki menu açılacaktır
				User Operations
				1. Register
				2. Login
				3. Find
				4. Quit
				Option:
			- Register seçildiğinde yeni kullanıcıya ilişkin bilgiler stdin'den istenecek ve kayıt işlemi yapılacaktır.
			Kullanıcının var olması durumu mesaj ile belirtilecek ve menüye geri dönülecektir.

			- Find seçildiğinde stdin'den username istenecek ve ilgili user'ın bilgileri stdout'a gönderilip menüye geri
			dönülecektir.

			- Quit seçildiğinde ana menüye geri dönülecektir

		Exit seçildiğinde uygulama sonlandırılacaktır
	Açıklamalar:
		- Uygulamada PaymentAppDataService yapısının fonksiyonları ile (yani servis katmanına erişilerek) işlem
		yapılacaktır
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
