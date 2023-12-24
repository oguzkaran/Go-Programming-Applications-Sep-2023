package main

import (
	"PaymentServiceApp/csd/console"
	"PaymentServiceApp/data/entity"
	"PaymentServiceApp/data/repository"
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

	userRepository := repository.NewUserRepository(db)

	for {
		user := getUser()

		if user == nil {
			break
		}

		_, err := userRepository.Save(user)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println("All users:")

		users, err := userRepository.FindAll()

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		for _, u := range users {
			fmt.Printf("%s, %s, %s, %s\n", u.Username, u.Password, u.Name, u.Phone)
		}
	}

	err = db.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
