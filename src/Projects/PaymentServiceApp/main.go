/*
------------------------------------------------------------------------------------------------------------------------

	Sınıf Çalışması: Aşağıdaki uygulamayı yazınız:
		- Uygulama aşağıdaki menü ile başlayacaktır:
				1. Users
				2. Products
				3. Payments
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
	"PaymentServiceApp/csd/console"
	"PaymentServiceApp/data/dal"
	"PaymentServiceApp/data/entity"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func getUser() *entity.User {
	username := console.ReadString("Input username:")

	if username == "quit" {
		return nil
	}

	return entity.NewUser(username, console.ReadString("Input password:"), console.ReadString("Input name:"),
		console.ReadString("Input phone:"))
}

func main() {
	connStr := "postgresql://postgres:csd1993@localhost/gs23_paymentdb?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Connected...")

	helper := dal.NewPaymentAppHelper(db)

	for {
		username := console.ReadString("Input username:")

		if username == "quit" {
			break
		}

		user, err := helper.FindUserByUsername(username)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if user != nil {
			fmt.Printf("User %s found:%s, %s\n", username, user.Name, user.Password)
		} else {
			fmt.Printf("User %s not found\n", username)
		}

	}

	err = db.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
