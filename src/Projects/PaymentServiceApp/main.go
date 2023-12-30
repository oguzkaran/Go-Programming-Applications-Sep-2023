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
